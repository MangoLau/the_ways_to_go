package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfUrl := r.URL.Path[len("/hello/"):]
	fmt.Fprintf(w, "Hello %s!", remPartOfUrl)
}

func shouthelloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfUrl := r.URL.Path[len("/shouthello/"):]
	fmt.Fprintf(w, "Hello %s!", strings.ToUpper(remPartOfUrl))
}

func main() {
	http.HandleFunc("/hello/", helloHandler)
	http.HandleFunc("/shouthello/", shouthelloHandler)
	err := http.ListenAndServe("localhost:9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}