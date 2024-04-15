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

const contentDir = "internal/content"

func Blog(rw http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(contentDir)
	if err != nil {
		log.Println("error reading directory", err)
		rw.Write([]byte("Internal Error"))
	}

	posts := make([]model.Post, len(files))
	for i, f := range files {
		openFile, err := os.Open(contentDir + "/" + f.Name())
		if err != nil {
			log.Printf("error opening file [%s]: %s", f.Name(), err)
			rw.Write([]byte("Internal Error"))
		}
		content, err := io.ReadAll(openFile)
		if err != nil {
			log.Printf("error reading all file [%s]: %s", f.Name(), err)
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
		post.Content = view.Unsafe(buf.String())
		posts[i] = post
	}
	view.Base(view.Blog(posts)).Render(r.Context(), rw)
}
