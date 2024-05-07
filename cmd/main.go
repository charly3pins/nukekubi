package main

import (
	"log"
	"net/http"

	"github.com/charly3pins/nukekubi/internal/handler"
	"github.com/charly3pins/nukekubi/internal/view"

	"github.com/a-h/templ"
)

func main() {
	http.Handle("/", templ.Handler(view.Base(view.Home())))
	http.HandleFunc("/blog", handler.Blog)
	http.HandleFunc("/blog/{slug}", handler.Post)

	// Manage posts
	http.HandleFunc("GET /posts", handler.Posts)
	http.HandleFunc("GET /posts/new", handler.PostNewForm)
	http.HandleFunc("POST /posts", handler.PostNew)
	http.HandleFunc("/posts/{slug}", handler.PostEdit)

	http.HandleFunc("/search", handler.Search)

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
