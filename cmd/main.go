package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"task_tracker/internal/storage"
	"task_tracker/internal/task"
)


var filename string = "tasks.json" 

func main() {
	// Точка входа программы
	// 
	// add "Заголовок" "Описание"
	// list
	// 
	switch  os.Args[1]{
	case "add":
		countID, err := storage.CounterID(filename)
		if err != nil{
			log.Fatal("Ошибка счетчика:",err)
		}
		task := task.AddTask(countID, os.Args[2], os.Args[3], true)
	
		log.Println("Task added:", task)
	
		err = storage.SaveTasksToFile(filename, task)
		if err != nil {
			log.Fatal("Error saving tasks to file:", err)
		}
	case "list":
		tasks, err := storage.LoadTasksUpToFile(filename)
		if err != nil && err != io.EOF{
			log.Fatal("Ошибка при загрузке данных с файла:", err)
		} 
		for i, v := range tasks {
			fmt.Print(tasks,storage.ListTask(i, v))	
		}
		
	case "update":
		// Call UpdateTask function
		
		// err = storage.UpdateTask(tasks, filename)
		// if err != nil{
		// 	log.Fatal("Ошибка записи в файл Update:",err)
		// }
		
	case "delete":
		// Call DeleteTask function
		//Запись данных файла в мапу
		tasks, err := storage.LoadTasksUpToFile(filename)
			if err != nil && err != io.EOF{
			log.Fatal("Ошибка при загрузке данных с файла:", err)
		} 
		// Получение ID для удаления и удаление 
		idKey, err := strconv.Atoi(os.Args[2])
		if err != nil{
			log.Fatal("Ошибка парсинга аргумента:", err)
		}
		delete(tasks, idKey)

		// Запись измененной мапы в файл
		err = storage.DeleteTask(tasks, filename)
		if err != nil{
			log.Fatal("Error saving tasks to file:", err)
		}

	default:
		log.Println("Unknown command")
		
	}
}