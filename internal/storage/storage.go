package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"task_tracker/internal/task"
	"time"
)

var tasks map[int]task.Task = make(map[int]task.Task)

// Сохраняет задачу в файл
func SaveTasksToFile(filename string, task *task.Task) error {
	jTask := *task
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

// достает данные из файла и записывает в map(может вернуть io.EOF)
func LoadTasksUpToFile(filename string) (map[int]task.Task, error) {
	jsonFile, err := os.OpenFile(filename, os.O_APPEND|os.O_EXCL|os.O_RDONLY, 0644)
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

// возвращает форматированую строку структуры 
func ListTask(task task.Task) string {
	return fmt.Sprintf("[ID:%d] Title: %s\t Status:%v\t CreatedAt:%s\t UpdatedAt: %s\n", task.ID, task.Title, task.Status, task.CreatedAt, task.UpdatedAt)
}
// изменяет мапу по айди и записывает в файл
func UpdateTask(filename string) error {
	tasks, err := LoadTasksUpToFile(filename)
	if err != nil && err != io.EOF {
		return fmt.Errorf("ошибка при загрузке данных с файла:%v", err)
	}

	idToUpdate, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return fmt.Errorf("ошибка парсинга аргумента id %v", err)
	}
	updatedTitle := os.Args[3]
	updatedAt := time.Now()
	task := task.AddTask(tasks[idToUpdate].ID, updatedTitle, true, tasks[idToUpdate].CreatedAt, &updatedAt)
	tasks[idToUpdate] = *task
	
	err = SaveInFileWithTrunc(filename, tasks)
	if err != nil{
		return fmt.Errorf("ошибка записи в файл deleted: %v", err)
	}

	return err
}

// Удаляет таску по айди.
func DeleteTask(filename string) error {
	tasks, err := LoadTasksUpToFile(filename)
	if err != nil && err != io.EOF {
		return fmt.Errorf("ошибка при загрузке данных с файла:%v", err)
	}

	idKey, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return fmt.Errorf("ошибка парсинга аргумента: %v", err)
	}
	delete(tasks, idKey)

	err = SaveInFileWithTrunc(filename, tasks)
	if err != nil{
		return fmt.Errorf("ошибка записи в файл deleted: %v", err)
	}
	return err
}

// Загружает файл в мапу, ищет последнее айди.
func FindLastId(filename string) (int, error) {
	tasks, err := LoadTasksUpToFile(filename)
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

// декодирует мапу и записывает в json файл 
func SaveInFileWithTrunc(filename string, tasks map[int]task.Task) error{
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
