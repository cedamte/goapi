package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe("localhost:11111", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		// The code above does the same thing
		//log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Asset not found ❌\n"))
		return
	}
	_, err := fmt.Fprintf(w, "Hello World 🤣 \n")
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("We are all winner 🚀 \n" +
		"Rest API v1\n"))
}
