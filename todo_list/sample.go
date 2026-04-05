package todolist
package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	
	"github.com/google/uuid" // Execute: go get github.com/google/uuid
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var (
	books   = make(map[string]Book)
	booksMu sync.Mutex
)

func main() {
	http.HandleFunc("/books/", booksHandler)
	println("Servidor CRUD de livros rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	id := ""
	if len(pathParts) > 2 {
		id = pathParts[2]
	}

	// Controle de métodos HTTP
	switch {
	case id == "":
		// Rota /books/
		switch r.Method {
		case http.MethodPost:
			createBook(w, r)
		case http.MethodGet:
			listBooks(w, r)
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	default:
		// Rota /books/{id}
		switch r.Method {
		case http.MethodGet:
			getBook(w, r, id)
		case http.MethodPut:
			updateBook(w, r, id)
		case http.MethodDelete:
			deleteBook(w, r, id)
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if newBook.Title == "" || newBook.Author == "" {
		http.Error(w, "Título e autor são obrigatórios", http.StatusBadRequest)
		return
	}

	booksMu.Lock()
	defer booksMu.Unlock()

	newBook.ID = uuid.New().String()
	books[newBook.ID] = newBook

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func listBooks(w http.ResponseWriter, r *http.Request) {
	booksMu.Lock()
	defer booksMu.Unlock()

	bookList := make([]Book, 0, len(books))
	for _, b := range books {
		bookList = append(bookList, b)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookList)
}

func getBook(w http.ResponseWriter, r *http.Request, id string) {
	booksMu.Lock()
	defer booksMu.Unlock()

	book, exists := books[id]
	if !exists {
		http.Error(w, "Livro não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request, id string) {
	booksMu.Lock()
	defer booksMu.Unlock()

	_, exists := books[id]
	if !exists {
		http.Error(w, "Livro não encontrado", http.StatusNotFound)
		return
	}

	var updatedBook Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedBook.ID = id // Mantém o ID original
	books[id] = updatedBook

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

func deleteBook(w http.ResponseWriter, r *http.Request, id string) {
	booksMu.Lock()
	defer booksMu.Unlock()

	_, exists := books[id]
	if !exists {
		http.Error(w, "Livro não encontrado", http.StatusNotFound)
		return
	}

	delete(books, id)
	w.WriteHeader(http.StatusNoContent)
}