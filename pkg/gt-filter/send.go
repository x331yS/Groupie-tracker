package gt_filter

import (
	"../gt-error"
	"net/http"
	"strings"
	"text/template"
)

//Send info about artist
func SendArtist(w http.ResponseWriter, r *http.Request, toSearch string) {
	All.ID = -1
	//Searching artist id by name
	for i := 0; i < 52; i++ {
		if strings.ToLower(All.Artists[i].Name) == strings.ToLower(toSearch) {
			All.ID = i
			break
		}
	}
	//Sending not found page if artist not found
	if All.ID == -1 {
		temp, err := template.ParseFiles("./static/templates/noresult.html")
		if err != nil {
			gt_error.InternalServerError(w, r)
			return
		}
		_ = temp.Execute(w, toSearch)
	} else {
		//Send info to user
		temp, err := template.ParseFiles("./static/templates/artist.html")
		if err != nil {
			gt_error.InternalServerError(w, r)
			return
		}
		_ = temp.Execute(w, All)
	}
}
