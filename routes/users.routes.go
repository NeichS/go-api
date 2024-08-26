package routes

import (
	"encoding/json"
	"net/http"
	"github.com/NeichS/api-rest-crud/db"
	"github.com/NeichS/api-rest-crud/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r) 
	//db.DB.First(&user, params["id"])
	db.DB.Preload("Tasks").First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		
	} else {
		//db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
		json.NewEncoder(w).Encode(&user) 
	}
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User;
	json.NewDecoder(r.Body).Decode(&user) //asigna el contenido del request a la variable

	createdUser := db.DB.Create(&user)

	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(&user)
	}
	
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)  
	var user models.User
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		
	} else {
		db.DB.Delete(&user) //esto hace borrado logico
		//si de verdad queremos removerlo de la base de datos usamos
		//db.DB.Unscoped().Delete(%user)
		json.NewEncoder(w).Encode(&user)
	}
}
