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

	log.Fatal(http.ListenAndServe(":8080", nil))
}
