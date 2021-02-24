package pkg

import (
	"html/template"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if artists == nil {
		ConnnectParseAPI(w, r, "artists")
	}
	if r.URL.Path != "/" {
		StatusNotFound(w, r)
		return
	}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, artists)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	if artists == nil {
		ConnnectParseAPI(w, r, "artists")
	}
	id, _ := strconv.Atoi(r.URL.Path[len("/about/"):])
	if id >= 1 && id <= len(artists) {
		ConnnectParseAPI(w, r, "relation")
		id = id - 1
		artist := artists[id]
		about := &About{ID: id, Image: artist.Image, Name: artist.Name, Members: artist.Members, FirstAlbum: artist.FirstAlbum, CreationDate: artist.CreationDate, RelationData: relations}
		t, _ := template.ParseFiles("templates/about.html")
		t.Execute(w, about)
	} else {
		BadRequest(w, r)
	}
}

