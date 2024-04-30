package handler

import (
	"net/http"

	"github.com/charly3pins/nukekubi/internal/view"
)

func Post(rw http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		rw.Write([]byte("Bad Request: slug empty"))
	}
	for _, post := range posts {
		if post.Slug == slug {
			view.Base(view.Post(post)).Render(r.Context(), rw)
		}
	}
}
