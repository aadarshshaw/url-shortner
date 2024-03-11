package main

import (
	"net/http"
	"strings"
)

func HandleSanity(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	w.Write([]byte("Sanity Check"))
}

func IndexFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func GetAllURLs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	urls, err := ReadAllURLsFromDB()
	if err != nil {
		panic(err)
	}
	w.Write([]byte(urls))
}

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	url := r.FormValue("url")
	shorturl, err := WriteIfNotExists(url)
	if err != nil {
		panic(err)
	}
	htmlstr := `<p>Short URL: <a href="http://localhost:8080/r/` + shorturl + `">http://localhost:8080/r/` + shorturl + `</a></p>`
	w.Write([]byte(htmlstr))

}

func RedirectShortURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	shorturl := strings.TrimPrefix(r.URL.Path, "/r/")
	url, err := GetURLFromShortURL(shorturl)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
