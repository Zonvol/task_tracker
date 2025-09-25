package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strconv"
	"task_tracker/internal/task"
	"time"
)

var tasks map[int]task.Task = make(map[int]task.Task)

// получение пути к файлу, current directory + filename
func GetFilepath()(string,error){
	currentDir, err := os.Getwd()
	if err != nil{
		return "", fmt.Errorf("cannot get filepath:%w",err)
	}
	filename := "tasks.json" 
	return path.Join(currentDir,filename), nil
}

// Сохраняет задачу в файл
func SaveTasksToFile(task *task.Task) error {
	filename, err := GetFilepath()
	if err != nil{
		return fmt.Errorf("cannot get filepath:%w",err)
	}
	jTask := task
	data, err := json.Marshal(jTask)
	if err != nil {
		return fmt.Errorf("cannot marshalling json file:%w", err)
	}
	data = append(data, '\n') // Add newline for better readability
	jsonFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("cannot create/open jsonfile: %w", err)
	}
	defer jsonFile.Close()
	log.Println("File created/opened successfully", jsonFile.Name())
	_, err = jsonFile.Write(data)
	if err != nil {
		return err
	}
	return err
}

// декодирует мапу и записывает в json файл
func SaveInFileWithTrunc(tasks map[int]task.Task) error {
	filename, err := GetFilepath()
	if err != nil{
		return fmt.Errorf("cannot get filepath:%w",err)
	}
	jsonFile, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("cannot create/open jsonfile: %w", err)
	}

	encoder := json.NewEncoder(jsonFile)
	for _, v := range tasks {
		if err := encoder.Encode(v); err != nil {
			return fmt.Errorf("ошибка encode: %v", err)
		}
	}
	return err
}

// достает данные из файла и записывает в map(может вернуть io.EOF)
func LoadTasksUpToFile() (map[int]task.Task, error) {
	filename, err := GetFilepath()
	if err != nil{
		return nil, fmt.Errorf("cannot get filepath:%w",err)
	}
	jsonFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("cannot open jsonfile: %v", err)
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	bufTask := new(task.Task)
	for {
		err = decoder.Decode(bufTask)

		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("ошибка при декодировании: %v", err)
		} else if err == io.EOF {
			break
		}
		tasks[bufTask.ID] = *bufTask
	}
	return tasks, err
}

// Создает файл если его не было, добавляет задачу
func AddTaskToFile() error {
	
	lastID, err := FindLastId()
	if err != nil {
		return fmt.Errorf("ошибка нахождения последнего айди:%w", err)
	}

	createdTime := time.Now()
	task := task.AddTask(lastID+1, os.Args[2], true, createdTime, nil)

	log.Print("Task added:\n", ListTask(*task))

	err = SaveTasksToFile(task)
	if err != nil {
		return fmt.Errorf("error saving tasks to file:%w", err)
	}
	return nil
}

// изменяет мапу по айди и записывает в файл
func UpdateTask() error {
	tasks, err := LoadTasksUpToFile()
	if err != nil && err != io.EOF {
		return fmt.Errorf("ошибка при загрузке данных с файла:%v", err)
	}

	idToUpdate, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return fmt.Errorf("ошибка парсинга аргумента id: введите число! %v", err)
	}
	updatedTitle := os.Args[3]
	updatedAt := time.Now()
	task := task.AddTask(tasks[idToUpdate].ID, updatedTitle, true, tasks[idToUpdate].CreatedAt, &updatedAt)
	tasks[idToUpdate] = *task

	err = SaveInFileWithTrunc(tasks)
	if err != nil {
		return fmt.Errorf("ошибка записи в файл deleted: %v", err)
	}

	return err
}

// Удаляет таску по айди.
func DeleteTask() error {
	tasks, err := LoadTasksUpToFile()
	if err != nil && err != io.EOF {
		return fmt.Errorf("ошибка при загрузке данных с файла:%v", err)
	}

	idKey, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return fmt.Errorf("ошибка парсинга аргумента id: введите число! %v", err)
	}
	delete(tasks, idKey)

	err = SaveInFileWithTrunc(tasks)
	if err != nil {
		return fmt.Errorf("ошибка записи в файл deleted: %v", err)
	}
	return err
}

// Загружает файл в мапу, ищет последнее айди.
func FindLastId() (int, error) {
	tasks, err := LoadTasksUpToFile()
	if err != nil && err != io.EOF {
		log.Fatal("Ошибка при загрузке данных с файла(findLastId):", err)
	}
	lastID := 0
	for _, v := range tasks {
		if lastID < v.ID {
			lastID = v.ID
		}
	}
	return lastID, nil
}

// возвращает форматированую строку структуры задачи
func ListTask(task task.Task) string {
	return fmt.Sprintf("[ID:%d] Title: %s\t Status:%v\t CreatedAt:%s\t UpdatedAt: %s\n", task.ID, task.Title, task.Status, task.CreatedAt, task.UpdatedAt)
}


// возвращает форматированую строку со списком команд
func СommandList() string {
	return ("\tСписок команд:\n1. add [Название]\n2. list\n3. update [id] [Новое название]\n4. delete [id]")
}

// Подсчитывает количество строк в файле (каждая задача по команде 'add' записывается с новой строки)
// func CounterID(filename string) (int, error) {
// 	jsonFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0644)
// 	if err != nil {
// 		return 0, fmt.Errorf("cannot open jsonfile: %v", err)
// 	}
// 	defer jsonFile.Close()

// 	scanner := bufio.NewScanner(jsonFile)
// 	countID := 0

// 	for scanner.Scan() {
// 		countID++
// 	}
// 	if scanner.Err() != nil {
// 		return 0, fmt.Errorf("ошибка при сканировании:%v", err)
// 	}
// 	return countID, nil
// }
