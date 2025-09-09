package main

import (
	
	"log"
	"os"
	"task_tracker/internal/storage"
	"task_tracker/internal/task"
)
var counter int = 0
var tasks map[int]task.Task
var filename string = "tasks.json" 

func main() {
	tasks = make(map[int]task.Task)

	
	// Further implementation goes here
	switch {
	case os.Args[1] == "add":
		counter++
		task := task.AddTask(counter, os.Args[2], os.Args[3], true)
		tasks[counter] = *task
		log.Println("Task added:", task)
		// Call AddTask function
		// Save tasks to file
	
		err := storage.SaveTasksToFile(filename, task)
		if err != nil {
			log.Fatal("Error saving tasks to file:", err)
		}
	case os.Args[1] == "list":
		storage.LoadTasksUpToFile(filename)
	case os.Args[1] == "update":
		// Call UpdateTask function
	case os.Args[1] == "delete":
		// Call DeleteTask function
	default:
		log.Println("Unknown command")
		
	}
}