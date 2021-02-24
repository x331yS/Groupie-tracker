package main

import (
	"./pkg"
	"fmt"
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/", pkg.Handler)
	http.HandleFunc("/about/", pkg.AboutHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	fmt.Println("Listening at localhost:1111\nHttp Status :", http.StatusOK)
	pkg.Openbrowser("http://localhost:1111")
	err := http.ListenAndServe(":1111", nil)
	if err != nil {
		log.Fatal(err)
	}
}
