package main

import (
	"./pkg/gt-filter"
	"./pkg/gt-search"
	"./pkg/gt-web"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	//Set timer
	fmt.Println("We are launching the Groupie-tracker project\n...Please wait the loading...")
	start := time.Now()
	//Parse JSON filter
	gt_filter.ParseJSON()
	//Data for search handler
	gt_search.Data()
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	//Main page handler
	http.HandleFunc("/", gt_web.IndexHandler)
	//Response page handler
	http.HandleFunc("/about/", gt_web.AboutHandler)
	//Search page handler
	http.HandleFunc("/search/", gt_search.SearchHandler)
	//Filter page handler
	http.HandleFunc("/filter/", gt_filter.FilterHandle)
	//Print timer
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Time elapsed for launching project :",elapsed)
	//LocalHost
	fmt.Println("Listening at localhost:1111\nHttp Status :", http.StatusOK)
	//Open favorite browser
	gt_web.Openbrowser("http://localhost:1111")
	err := http.ListenAndServe(":1111", nil)
	if err != nil {
		log.Fatal(err)
	}
}

