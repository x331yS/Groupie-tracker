package gt_index

import (
	"../gt-error"
	"html/template"
	"net/http"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//Get API
	if Artists == nil {
		ConnnectParseAPI(w, r, "artists")
	}
	if r.URL.Path != "/" {
		//Error 404 IndexHandler
		gt_error.StatusNotFound(w, r)
		return
	}
	//Index Template
	t, _ := template.ParseFiles("static/templates/index.html")
	_ = t.Execute(w,Artists)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	//Get API
	if Artists == nil {
		ConnnectParseAPI(w, r, "artists")
	}
	//Add ID of the artist after the page about
	id, _ := strconv.Atoi(r.URL.Path[len("/about/"):])
	suiv := id +1
	prec := id -1
	if suiv == 53 {
		suiv = 1
	}
	if prec == 0 {
		prec = 52
	}
	if id >= 1 && id <= len(Artists) {
		ConnnectParseAPI(w, r, "relation")
		id = id - 1
		artist := Artists[id]
		//Get Data from API
		about := &About{ID: id, Suiv: suiv, Prec: prec,Image: artist.Image, Name: artist.Name, Members: artist.Members, FirstAlbum: artist.FirstAlbum, CreationDate: artist.CreationDate, RelationData: Relations}
		t, _ := template.ParseFiles("./static/templates/about.html")
		_ = t.Execute(w, about)

	} else {
		//Error 400 Handler
		gt_error.StatusNotFound(w, r)
	}
}
