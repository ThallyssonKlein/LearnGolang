package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type Task struct {
	Title string `json:title`
	Done bool `json:done`
}

func main() {
	var option int
	var tasks []Task

	for true {
		fmt.Printf("Select 1 to Add and 2 to List and 3 to exit and 4 to check items, 5 to save tasks to json: ")
		fmt.Scan(&option)

		if option == 3 {
			break;
		} else if option == 1 {
			fmt.Printf("Type the task title: ")
			var title string
			fmt.Scan(&title)
			tasks = append(tasks, Task{
				Title: title,
				Done: false,
			})
		} else if option == 2 {
			for i, task := range tasks {
				fmt.Printf("Id: %d, Title: %s, Done: %t\n", i, task.Title, task.Done)
			}
		} else if option == 4 {
			fmt.Printf("Type the task id: ")
			var taskId int
			fmt.Scan(&taskId)
			fmt.Printf("Type 1 for check and 2 to uncheck: ")
			var check int
			fmt.Scan(&check)
			if check == 1 { tasks[taskId].Done = true } else { tasks[taskId].Done = false }
		} else if option == 5 {
			jsonData, err := json.MarshalIndent(tasks, "", " ")
			if err != nil {
				fmt.Println("Error converting to JSON", err)
			}

			var err2 = os.WriteFile("tasks.json", jsonData, 0644)
			if err2 != nil {
				fmt.Println("Error saving the file: ", err2)
			}

			fmt.Println("File saved successfully!")
		}
	}
}