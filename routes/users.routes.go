package routes

import "net/http"

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Algo agarrooo"))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user"))
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Post"))
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete"))
}
