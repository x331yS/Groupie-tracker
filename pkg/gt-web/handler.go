package gt_web

import (
	"../gt-error"
	"../gt-json"
	"html/template"
	"net/http"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if gt_json.Artists == nil {
		gt_json.ConnnectParseAPI(w, r, "artists")
	}
	if r.URL.Path != "/" {
		gt_error.StatusNotFound(w, r)
		return
	}
	t, _ := template.ParseFiles("static/templates/index.html")
	t.Execute(w, gt_json.Artists)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	if gt_json.Artists == nil {
		gt_json.ConnnectParseAPI(w, r, "artists")
	}
	id, _ := strconv.Atoi(r.URL.Path[len("/about/"):])
	if id >= 1 && id <= len(gt_json.Artists) {
		gt_json.ConnnectParseAPI(w, r, "relation")
		id = id - 1
		artist := gt_json.Artists[id]
		about := &gt_json.About{ID: id, Image: artist.Image, Name: artist.Name, Members: artist.Members, FirstAlbum: artist.FirstAlbum, CreationDate: artist.CreationDate, RelationData: gt_json.Relations}
		t, _ := template.ParseFiles("./static/templates/about.html")
		t.Execute(w, about)
	} else {
		gt_error.BadRequest(w, r)
	}
}
