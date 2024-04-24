package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/charly3pins/nukekubi/internal/model"
	"github.com/charly3pins/nukekubi/internal/view"
)

func Search(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("error parsing form", err)
	}

	text := r.FormValue("search")
	var postsFound []model.Post
	for _, post := range posts {
		if !strings.Contains(strings.ToUpper(post.Title), strings.ToUpper(text)) {
			continue
		}
		postsFound = append(postsFound, post)
	}
	if len(postsFound) > 0 {
		view.PostSearch(postsFound).Render(r.Context(), rw)
	}
}
