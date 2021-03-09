package main

import (
	"./pkg/gt-filter"
	"./pkg/gt-search"
	"./pkg/gt-web"
	"fmt"
	"log"
	"net/http"
)

func main() {
	gt_filter.ParseJSON()
	gt_search.Data()
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", gt_web.IndexHandler)
	http.HandleFunc("/about/", gt_web.AboutHandler)
	http.HandleFunc("/search/", gt_search.SearchHandler)
	http.HandleFunc("/filter/", gt_filter.FilterHandle)
	fmt.Println("Listening at localhost:1111\nHttp Status :", http.StatusOK)
//	gt_web.Openbrowser("http://localhost:1111")
	err := http.ListenAndServe(":1111", nil)
	if err != nil {
		log.Fatal(err)
	}
}

