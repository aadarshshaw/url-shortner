package main

import (
	"log"
	"net/http"
)

func Routes() {
	r := http.NewServeMux()
	r.HandleFunc("/", IndexFileHandler)
	r.HandleFunc("/ping", HandleSanity)
	r.HandleFunc("/all", GetAllURLs)
	r.HandleFunc("/create", CreateShortURL)
	r.HandleFunc("/r/", RedirectShortURL)
	log.Fatal(http.ListenAndServe(":8080", r), nil)
}
