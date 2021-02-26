package main

import (
	"./pkg/gt-web"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", gt_web.IndexHandler)
	http.HandleFunc("/about/", gt_web.AboutHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	fmt.Println("Listening at localhost:1111\nHttp Status :", http.StatusOK)
	gt_web.Openbrowser("http://localhost:1111")
	err := http.ListenAndServe(":1111", nil)
	if err != nil {
		log.Fatal(err)
	}
}
