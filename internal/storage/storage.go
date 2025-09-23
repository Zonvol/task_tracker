package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"task_tracker/internal/task"
)

var counter int
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

	return nil
}

// достает данные из файла и записывает в map
func LoadTasksUpToFile(filename string) (map[int]task.Task, error) {
	jsonFile, err := os.OpenFile(filename, os.O_APPEND|os.O_EXCL|os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("cannot open jsonfile: %v", err)
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	bufTask := new(task.Task)
	for {
		counter++
		err = decoder.Decode(bufTask)

		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("ошибка при декодировании: %v", err)
		} else if err == io.EOF {
			break
		}
		tasks[counter] = *bufTask
	}
	return tasks, nil
}

// возвращает форматированую строку структуры
func ListTask(id int, task task.Task) string {
	return fmt.Sprintf("[ID:%d] Title: %s\t Description: %s\t Status:%v\n", task.ID + 1, task.Title, task.Description, task.Status)
}

// func update record by ID ДОДЕЛАТЬ!!!!
func DeleteTask(tasks map[int]task.Task, filename string) error {
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
	return nil
}

// Подсчитывает количество строк в файле (каждая задача по команде 'add' записывается с новой строки)
func CounterID(filename string) (int, error) {
	jsonFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return 0, fmt.Errorf("cannot open jsonfile: %v", err)
	}
	defer jsonFile.Close()

	scanner := bufio.NewScanner(jsonFile)
	countID := 0

	for scanner.Scan() {
		countID++
	}
	if scanner.Err() != nil {
		return 0, fmt.Errorf("ошибка при сканировании:%v", err)
	}
	return countID, nil
}
