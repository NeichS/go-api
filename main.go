package main

import (
	"net/http"
	"github.com/NeichS/api-rest-crud/db"
	"github.com/NeichS/api-rest-crud/models"
	"github.com/NeichS/api-rest-crud/controllers"
	"github.com/gorilla/mux"
)

func main(){

  db.DBConnection()

  db.DB.AutoMigrate(models.Task{})
  db.DB.AutoMigrate(models.User{})

  r := mux.NewRouter()

  r.HandleFunc("/", controllers.HomeHandler)

  r.HandleFunc("/users", controllers.GetUsersHandler).Methods("GET")
  r.HandleFunc("/users/{id}", controllers.GetUserHandler).Methods("GET")
  r.HandleFunc("/users", controllers.PostUserHandler).Methods("POST")
  r.HandleFunc("/users/{id}", controllers.DeleteUsersHandler).Methods("DELETE")

  //Tasks

  r.HandleFunc("/tasks", controllers.GetTasksHandler).Methods("GET")
  r.HandleFunc("/tasks/{id}", controllers.GetTaskHandler).Methods("GET")
  r.HandleFunc("/tasks", controllers.CreateTasksHandler).Methods("POST")
  r.HandleFunc("/tasks/{id}", controllers.DeleteTasksHandler).Methods("DELETE")

  http.ListenAndServe(":3000", r)
} 
