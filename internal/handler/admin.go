package handler

import (
	"net/http"

	"github.com/charly3pins/nukekubi/internal/view"
)

func Admin(rw http.ResponseWriter, r *http.Request) {
	view.Base(view.Admin(posts)).Render(r.Context(), rw)
}
