package main

import (
	"fmt"
	"net/http"

	"github.com/NeichS/api-rest-crud/db"
	"github.com/NeichS/api-rest-crud/models"
	"github.com/NeichS/api-rest-crud/routes"
	"github.com/gorilla/mux"
)

func main(){

  db.DBConnection()

  db.DB.AutoMigrate(models.Task{})
  db.DB.AutoMigrate(models.User{})

  r := mux.NewRouter()

  r.HandleFunc("/", routes.HomeHandler)

  r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
  r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
  r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
  r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

  http.ListenAndServe(":3004", r)

  fmt.Print("por que estoy leyendo esto???")
} 
