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

var posts []model.Post

func init() {
	files, err := os.ReadDir(contentDir)
	if err != nil {
		log.Println("error reading directory", err)
		panic(err)
	}

	for _, f := range files {
		openFile, err := os.Open(contentDir + "/" + f.Name())
		if err != nil {
			log.Printf("error opening file [%s]: %s", f.Name(), err)
			panic(err)
		}
		content, err := io.ReadAll(openFile)
		if err != nil {
			log.Printf("error reading all file [%s]: %s", f.Name(), err)
			panic(err)
		}
		var post model.Post
		rest, err := frontmatter.Parse(strings.NewReader(string(content)), &post)
		if err != nil {
			log.Printf("error parsing content [%s]: %s", string(content), err)
			panic(err)
		}

		var buf bytes.Buffer
		err = goldmark.Convert(rest, &buf)
		if err != nil {
			log.Printf("error converting content to goldmark buffer: %s", err)
			panic(err)
		}
		post.Content = view.Unsafe(buf.String())
		if post.Slug == "" {
			post.Slug = strings.Split(f.Name(), ".")[0]
		}
		posts = append(posts, post)
	}
}

func Blog(rw http.ResponseWriter, r *http.Request) {
	view.Base(view.Blog(posts)).Render(r.Context(), rw)
}
