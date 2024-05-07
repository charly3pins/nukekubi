package handler

import (
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/charly3pins/nukekubi/internal/model"
	"github.com/charly3pins/nukekubi/internal/view"
)

func PostNewForm(rw http.ResponseWriter, r *http.Request) {
	view.Base(view.PostNewForm()).Render(r.Context(), rw)
}

func PostNew(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	post := model.Post{
		Title:       r.FormValue("title"),
		Author:      r.FormValue("author"),
		Description: r.FormValue("description"),
		Date:        r.FormValue("date"),
		Slug:        r.FormValue("slug"),
		ContentRaw:  r.FormValue("content"),
		Tags:        []string{"go", "htmx", "templ"},
	}
	mdTmpl := `
+++
title = "{{.Title}}"
date = "{{.Date}}"
author = "{{.Author}}"
description = "{{.Description}}"
tags = [{{range $i, $v := .Tags}}{{if $i}}, {{end}}"{{$v}}"{{end}}]
slug = "{{.Slug}}"

image = "/images/gopher-hiking.svg"
+++
{{.ContentRaw}}
`
	file, err := os.Create(contentDir + "/" + post.Slug + ".md")
	if err != nil {
		log.Println("error creating file for a post", err)
		rw.Write([]byte("Internal server error"))
		rw.WriteHeader(http.StatusInternalServerError)

	}

	t := template.Must(template.New("post").Parse(mdTmpl))
	err = t.Execute(file, post)
	if err != nil {
		log.Println("error generating template post", err)
		rw.Write([]byte("Internal server error"))
		rw.WriteHeader(http.StatusInternalServerError)

	}
	posts = append(posts, post)
	view.Base(view.Posts(posts)).Render(r.Context(), rw)
}

func join(sep string, s ...string) string {
	return strings.Join(s, sep)
}
