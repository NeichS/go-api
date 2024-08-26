package main

import (
	"net/http"

	"github.com/NeichS/api-rest-crud/database"
	"github.com/NeichS/api-rest-crud/models"
	"github.com/NeichS/api-rest-crud/routes"
	"github.com/gorilla/mux"
)


func main(){

  database.DBConnection()

  database.DB.AutoMigrate(models.Task{})
  database.DB.AutoMigrate(models.User{})
  
  r := mux.NewRouter()

  r.HandleFunc("/", routes.HomeHandler)

  http.ListenAndServe(":3003", r)
} 
