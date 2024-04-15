package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/charly3pins/nukekubi/internal/model"
	"github.com/charly3pins/nukekubi/internal/view"
)

func Blog(rw http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("internal/content")
	if err != nil {
		log.Println("error reading directory", err)
		rw.Write([]byte("Internal Error"))
	}
	posts := make([]model.Post, len(files))
	for i, f := range files {
		posts[i].Title = f.Name()
	}
	view.Base(view.Blog(posts)).Render(r.Context(), rw)
}
