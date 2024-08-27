package controllers

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hola como andas</h1>"))
}
  