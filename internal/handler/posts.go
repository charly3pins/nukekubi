package handler

import (
	"net/http"

	"github.com/charly3pins/nukekubi/internal/view"
)

func Posts(rw http.ResponseWriter, r *http.Request) {
	view.Base(view.Posts(posts)).Render(r.Context(), rw)
}
