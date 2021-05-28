package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	listenAddr 	= flag.String("http", ":8080", "http listen address")
	dataFile	= flag.String("file", "store.gob", "data store file name")
	hostname	= flag.String("host", "localhost:8080", "http host name")
)

var store *URLStore

func main() {
	flag.Parse()
	store = NewURLStore(*dataFile)
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	if err := http.ListenAndServe(*listenAddr, nil); err != nil {
		log.Fatal("start error:", err)
		return
	}
}

func Add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, AddForm)
		return
	}
	key := store.Put(url)
	fmt.Fprintf(w, "http://%s/%s", *hostname, key)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	url := store.Get(key)
	if url == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`