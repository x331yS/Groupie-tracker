package gt_favicon

import "net/http"

func FaviconHandler(w http.ResponseWriter, r * http.Request) {
	//Handler for favicon
	http.ServeFile(w, r, "static/assets/img/favicon.ico")
}
