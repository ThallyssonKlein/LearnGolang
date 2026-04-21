package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"time"
	"os"
	"strings"
	"sync"
	"strconv"
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

type TaskStore struct {
	mu		sync.Mutex
	tasks   []Task
	nextID  int
}

var store = &TaskStore{
	tasks:	[]Task{},
	nextID: 1,
}

func main() {
	mux := http.NewServeMux()

	http.HandleFunc("GET /ping", pingHandler)
	http.HandleFunc("POST /task", saveTasksHandler)
	http.HandleFunc("GET /task", listTasksHandler)
	http.HandleFunc("PUT /task/{id}", updateTaskHandler)
	http.HandleFunc("POST /save_tasks", saveFileHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", mux)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func saveTasksHandler(w http.ResponseWriter, r *http.Request) {
	var t Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "JSON Invalido", http.StatusBadRequest)
		return
	}
	
	store.mu.Lock()
	defer store.mu.Unlock()
	t.Id = store.nextID
	store.nextID++
	store.tasks = append(store.tasks, t)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

func listTasksHandler(w http.ResponseWriter, r *http.Request) {
	store.mu.Lock()
	defer store.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store.tasks)
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue(("id")))
	var tEdit TaskEdit

	if err := json.NewDecoder(r.Body).Decode(&tEdit); err != nil {
		http.Error(w, "JSOn Invalido", http.StatusBadRequest)
		return
	}

	store.mu.Lock()
	defer store.mu.Unlock()

	for i, task := range store.tasks {
		if task.Id == id {
			store.tasks[i].Done = tEdit.Done

			w.WriteHeader(http.StatusOK)
			return
		}
	}
}


func saveFileHandler(w http.ResponseWriter, r *http.Request) {
	store.mu.Lock()
	defer store.mu.Unlock()
	jsonData, err := json.MarshalIndent(store.tasks, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var err2 = os.WriteFile("tasks.json", jsonData, 0644)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}