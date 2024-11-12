package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/arjunmalhotra1/go-template/internal/database"
	"github.com/arjunmalhotra1/go-template/internal/handler"
	"github.com/arjunmalhotra1/go-template/internal/repository"
	"github.com/arjunmalhotra1/go-template/internal/service"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	repo := repository.NewTaskRepository(db)
	svc := service.NewTaskService(repo)
	handler := handler.NewTaskHandler(svc)

	r := mux.NewRouter()
	r.HandleFunc("/tasks", handler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", handler.GetTasks).Methods("GET")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
