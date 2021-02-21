package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	r.Host("{subdomain}:8000").
		Path("/articles/{category}/{id:[0-9]+}").
		HandlerFunc(ArticleHandler).
		Name("article")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	url, _ := r.Get("article").URL(
		"subdomain", "localhost",
		"category", "technology",
		"id", "42")
	log.Println(url)
	log.Fatal(srv.ListenAndServe())
}
