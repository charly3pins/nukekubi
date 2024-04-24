package handler

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/charly3pins/nukekubi/internal/model"
	"github.com/charly3pins/nukekubi/internal/view"

	"github.com/adrg/frontmatter"
	"github.com/yuin/goldmark"
)

func PostEdit(rw http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		rw.Write([]byte("Bad Request: slug empty"))
	}
	openFile, err := os.Open(contentDir + "/" + slug + ".md")
	if err != nil {
		log.Printf("error opening file [%s].md: %s", slug, err)
		rw.Write([]byte("Internal Error"))
	}
	content, err := io.ReadAll(openFile)
	if err != nil {
		log.Printf("error reading all file [%s].md: %s", slug, err)
		rw.Write([]byte("Internal Error"))
	}
	var post model.Post
	rest, err := frontmatter.Parse(strings.NewReader(string(content)), &post)
	if err != nil {
		log.Printf("error parsing content [%s]: %s", string(content), err)
		rw.Write([]byte("Internal Error"))
	}

	var buf bytes.Buffer
	err = goldmark.Convert(rest, &buf)
	if err != nil {
		log.Printf("error converting content to goldmark buffer: %s", err)
		rw.Write([]byte("Internal Error"))
	}
	post.ContentRaw = buf.String()
	view.Base(view.PostEdit(post)).Render(r.Context(), rw)
}
