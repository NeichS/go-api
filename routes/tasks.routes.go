package routes

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/NeichS/api-rest-crud/db"
	"github.com/NeichS/api-rest-crud/models"
)


func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func CreateTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task;
	json.NewDecoder(r.Body).Decode(&task) //asigna el contenido del request a la variable

	createdUser := db.DB.Create(&task)

	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(&task)
	}
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("task not found"))
		
	} else {
		json.NewEncoder(w).Encode(&task)
	}

}
func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)  
	var task models.Task
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		
	} else {
		db.DB.Delete(&task)
		json.NewEncoder(w).Encode(&task)
	}
}