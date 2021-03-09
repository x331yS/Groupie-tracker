package gt_search

import (
	"../gt-error"
	"net/http"
	"text/template"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	//Handle 404 error
	if r.URL.Path != "/search/" {
		gt_error.StatusNotFound(w, r)
		return
	}
	t, err := template.ParseFiles("./static/templates/search.html")
	if err != nil {
		gt_error.InternalServerError(w, r)
		return
	}
	_ = t.Execute(w, allData)
}
