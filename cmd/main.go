package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"task_tracker/internal/storage"
)

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
		// Call AddTask function
		if len(os.Args) < 3 {
			fmt.Print("Неверный ввод команды add!\n", storage.СommandList())
			return
		}
		if err := storage.AddTaskToFile(); err != nil{
			log.Fatal("Ошибка при добавлении задачи add:", err)
		}
		
	case "list":
		// Call Formating list of tasks
		tasks, err := storage.LoadTasksUpToFile()
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
		
		if err := storage.UpdateTask(); err != nil{
			log.Fatal("Ошибка записи в файл Update:",err)
		}
		
	case "delete":
		// Call DeleteTask function
		if len(os.Args) < 3 {
			fmt.Print("Неверный ввод команды delete!\n", storage.СommandList())
			return
		}
		if err := storage.DeleteTask(); err != nil{
			log.Fatal("Error saving tasks to file:", err)
		}

	default:
		log.Println("Unknown command")
	}
}

