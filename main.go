package main

import (
	"fmt"
	"goapi/handlers"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("localhost:11111", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		// The code above does the same thing
		//log.Fatal(err)
	}
}
