package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"time"
	"os"
	"strings"
)

type Task struct {
	Id int `json:id`
	Title string `json:title`
	Done bool `json:done`
}

type TaskEdit struct {
	Done bool `json:done`
}

// func main() {
// 	var option int
// 	var tasks []Task

// 	for true {
// 		fmt.Printf("Select 1 to Add and 2 to List and 3 to exit and 4 to check items, 5 to save tasks to json: ")
// 		fmt.Scan(&option)

// 		if option == 3 {
// 			break;
// 		} else if option == 1 {
// 			fmt.Printf("Type the task title: ")
// 			var title string
// 			fmt.Scan(&title)
// 			tasks = append(tasks, Task{
// 				Title: title,
// 				Done: false,
// 			})
// 		} else if option == 2 {
// 			for i, task := range tasks {
// 				fmt.Printf("Id: %d, Title: %s, Done: %t\n", i, task.Title, task.Done)
// 			}
// 		} else if option == 4 {
// 			fmt.Printf("Type the task id: ")
// 			var taskId int
// 			fmt.Scan(&taskId)
// 			fmt.Printf("Type 1 for check and 2 to uncheck: ")
// 			var check int
// 			fmt.Scan(&check)
// 			if check == 1 { tasks[taskId].Done = true } else { tasks[taskId].Done = false }
// 		} else if option == 5 {
// 			jsonData, err := json.MarshalIndent(tasks, "", " ")
// 			if err != nil {
// 				fmt.Println("Error converting to JSON", err)
// 			}

// 			var err2 = os.WriteFile("tasks.json", jsonData, 0644)
// 			if err2 != nil {
// 				fmt.Println("Error saving the file: ", err2)
// 			}

// 			fmt.Println("File saved successfully!")
// 		}
// 	}
// }

var tasks []Task

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/task", tasksHandler)
	http.HandleFunc("/save_tasks", saveTasksHandler)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func saveTasksHandler(w http.ResponseWriter, r *http.Request) {
	//
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	id := ""
	if len(pathParts) > 2 {
		id = pathParts[2]
	}

	if id == "" {
		switch r.Method {
			case http.MethodPost:
				createTask(w, r)
			case http.MethodGet:
				// find all tasks
		}
	} else {
		if (r.Method == http.MethodPatch) {
			// check one task
		}
	}
}

func findTaskById(id int) (Task, bool){
	for _, task := range tasks {
		if task.Id == id {
			return task, true
		}
	}

	return Task{}, false
}

func updateTask(w http.ResponseWriter, r *http.Request, id int) {
	var taskEdit TaskEdit
	err := json.NewDecoder(r.Body).Decode(&taskEdit)

	if err != nil {
		http.Error(w, "Invalid body!", http.StatusBadRequest)
		return
	}

	var task := findTaskById(id)
}