package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"task_tracker/internal/storage"
	"task_tracker/internal/task"
	"time"
)


var filename string = "tasks.json" 

func main() {

	// Точка входа программы:
	// go run .\cmd\main.go [команда]...
	// add ["Заголовок"]  		*в кавычках ("")
	// list
	// update [id] ["Новый заголовок"]
	// delete [id]
	if len(os.Args)<2{
		fmt.Print(storage.СommandList())
		return
	}
	switch  os.Args[1]{
	case "add":
		if len(os.Args) < 3 {
			fmt.Print("Неверный ввод команды add!\n", storage.СommandList())
			return
		}
		lastID, err := storage.FindLastId(filename)
		if err != nil{
			log.Fatal("Ошибка нахождения последнего айди:",err)
		}

		createdTime := time.Now()
		task := task.AddTask(lastID + 1, os.Args[2], true, createdTime, nil)
	
		log.Print("Task added:\n", storage.ListTask(*task))
	
		err = storage.SaveTasksToFile(filename, task)
		if err != nil {
			log.Fatal("Error saving tasks to file:", err)
		}
	case "list":
		tasks, err := storage.LoadTasksUpToFile(filename)
		if err != nil && err != io.EOF{
			log.Fatal("Ошибка при загрузке данных с файла:", err)
		} 
		for _, v := range tasks {
			fmt.Print(storage.ListTask(v))	
		}
		
	case "update":
		// Call UpdateTask function
		if len(os.Args) < 4 {
			fmt.Print("Неверный ввод команды update!\n", storage.СommandList())
			return
		}
		err := storage.UpdateTask(filename)
		if err != nil{
			log.Fatal("Ошибка записи в файл Update:",err)
		}
		
	case "delete":
		// Call DeleteTask function
		if len(os.Args) < 3 {
			fmt.Print("Неверный ввод команды delete!\n", storage.СommandList())
			return
		}
		err := storage.DeleteTask(filename)
		if err != nil{
			log.Fatal("Error saving tasks to file:", err)
		}

	default:
		log.Println("Unknown command")
		
	}
}

