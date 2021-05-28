package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

var (
	listenAddr 	= flag.String("http", ":8080", "http listen address")
	dataFile	= flag.String("file", "store.json", "data store file name")
	hostname	= flag.String("host", "localhost:8080", "http host name")
	masterAddr	= flag.String("master", "", "RPC master address")
	rpcEnabled	= flag.Bool("rpc", true, "enable RPC server")
)

var store Store

func main() {
	flag.Parse()
	if *masterAddr != "" {
		store = NewProxyStore(*masterAddr)
	} else {
		store = NewURLStore(*dataFile)
	}
	if *rpcEnabled { // the master is the rpc server:
		if err := rpc.RegisterName("Store", store); err != nil {
			log.Fatal("rpc.registerName error:", err)
		}
		rpc.HandleHTTP()
	}
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
	var key string
	if err := store.Put(&key, &url); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "http://%s/%s", *hostname, key)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	if key == "" {
		http.NotFound(w, r)
		return
	}
	url := ""
	if err := store.Get(&key, &url); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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