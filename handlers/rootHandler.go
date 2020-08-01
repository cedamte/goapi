package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
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
