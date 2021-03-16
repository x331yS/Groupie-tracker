package gt_error

import (
	"html/template"
	"log"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	//Error 500 Template Handler
	w.WriteHeader(http.StatusInternalServerError)
	t, _ := template.ParseFiles("static/templates/500.html")
	err := t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	//Error 400 Template Handler
	w.WriteHeader(http.StatusBadRequest)
	t, err := template.ParseFiles("static/templates/400.html")
	if err != nil {
		InternalServerError(w, r)
	}
	_ = t.Execute(w, nil)
}

func StatusNotFound(w http.ResponseWriter, r *http.Request) {
	//Error 404 Template Handler
	w.WriteHeader(http.StatusNotFound)
	t, err := template.ParseFiles("static/templates/404.html")
	if err != nil {
		InternalServerError(w, r)
		return
	}
	_ = t.Execute(w, nil)
}