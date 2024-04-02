package api

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()
	taskAPI := NewTaskAPI(db)

	router.HandleFunc("/api/tasks", taskAPI.GetAllTasks).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", taskAPI.GetTaskByID).Methods("GET")
	router.HandleFunc("/api/tasks", taskAPI.CreateTask).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", taskAPI.UpdateTask).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", taskAPI.DeleteTask).Methods("DELETE")

	return router
}
