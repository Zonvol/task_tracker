package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"task_tracker/internal/task"
)
var counter int
var tasks map[int]task.Task = make(map[int]task.Task)

func SaveTasksToFile(filename string, task *task.Task) error {
    jTask := *task
    data, err := json.Marshal(jTask)
    if err != nil {
        return err
    }
    data = append(data, '\n') // Add newline for better readability
    jsonFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY , 0644)
	if err != nil {
		log.Fatal("Cannot create/open jsonfile", err)
	}
	defer jsonFile.Close()
	log.Println("File created successfully", jsonFile.Name())
    _, err = jsonFile.Write(data)
    if err != nil {
        return err
    }
    
    return nil
}
func LoadTasksUpToFile(filename string) error{
     jsonFile, err := os.OpenFile(filename, os.O_APPEND|os.O_EXCL|os.O_RDONLY , 0644)
	if err != nil {
		log.Fatal("Cannot open jsonfile", err)
	}
	defer jsonFile.Close()
    decoder := json.NewDecoder(jsonFile)
    bufTask := new(task.Task)
    for {
        counter++
        err = decoder.Decode(bufTask);
        
        if err != nil && err != io.EOF{
            log.Fatal(err)
        } else if err == io.EOF {
            return io.EOF
        }
        tasks[counter] = *bufTask
        fmt.Print(counter,". ",listTask(counter, tasks))
    }
}
func listTask(id int, tasks map[int]task.Task)string{
   return fmt.Sprintf("Title: %s\t Description: %s\t Status:%v\n",tasks[id].Title,tasks[id].Description, tasks[id].Status )
}


