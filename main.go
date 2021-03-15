package main

import (
	"./pkg/gt-favicon"
	"./pkg/gt-filter"
	"./pkg/gt-index"
	"./pkg/gt-search"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	//Set timer
	start := time.Now()
	fmt.Println("We are launching the Groupie-tracker project\n...Please wait the loading...")
	//Parse JSON filter
	gt_filter.ParseJSON()
	//Data for search handler
	gt_search.Data()
	//For Directory settings
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	//Main page handler
	http.HandleFunc("/", gt_index.IndexHandler)
	//Response page handler
	http.HandleFunc("/about/", gt_index.AboutHandler)
	//Search page handler
	http.HandleFunc("/search/", gt_search.SearchHandler)
	//Filter page handler
	http.HandleFunc("/filter/", gt_filter.FilterHandle)
	//Favicon page handler
	http.HandleFunc("/favicon.ico/", gt_favicon.FaviconHandler)
	//Print timer
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Time elapsed for launching project :", elapsed)
	//LocalHost
	fmt.Println("Listening at localhost:1111\nHttp Status :", http.StatusOK)
	//Open favorite browser
	gt_index.Openbrowser("http://localhost:1111")
	err := http.ListenAndServe(":1111", nil)
	if err != nil {
		log.Fatal(err)
	}
}
