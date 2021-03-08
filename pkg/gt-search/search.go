package gt_search

import (
	"../gt-error"
	"io/ioutil"
	"net/http"
	"net/url"
	"text/template"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	//Handle 404 error
	if r.URL.Path != "/search/" {
		gt_error.StatusNotFound(w, r)
		return
	}
	//if len(r.URL.Path) >= 8 {
	//	if r.URL.Path[0:6] != "/filter/" {
	//		gt_error.StatusNotFound(w, r)
	//		return
	//	}
	//}
	//Check for request method
	switch r.Method {
	case "GET":
		temp, err := template.ParseFiles("./static/templates/search.html")
		if err != nil {
			gt_error.InternalServerError(w, r)
			return
		}
		temp.Execute(w, All)
	case "POST":
		var toSearch string   //users input
		var searchType string //what to search

		//Check for invalid request data
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			gt_error.BadRequest(w, r)
			return
		}
		query, err := url.ParseQuery(string(body))
		if err != nil {
			gt_error.BadRequest(w, r)
			return
		}
		for i, v := range query {
			switch i {
			case "toSearch":
				toSearch = v[0]
			case "searchType":
				searchType = v[0]
			default:
				gt_error.StatusNotFound(w, r)
				return
			}
		}
		if searchType != "artist" && searchType != "member" && searchType != "location" && searchType != "firstAlbum" && searchType != "creationDate" {
			gt_error.StatusNotFound(w, r)
			return
		}

		//What to send to user
		switch searchType {
		case "artist":
			SendArtist(w, r, toSearch)
		case "member":
			SendMember(w, r, toSearch)
		case "location":
			SendLocation(w, r, toSearch)
		case "firstAlbum":
			SendFirstAlbum(w, r, toSearch)
		case "creationDate":
			SendCreationDate(w, r, toSearch)
		}
	}
}
