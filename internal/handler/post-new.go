package handler

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/adrg/frontmatter"
	"github.com/charly3pins/nukekubi/internal/model"
	"github.com/charly3pins/nukekubi/internal/view"
	"github.com/yuin/goldmark"
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
		Tags:        []string{"go", "htmx", "templ"}, // TODO: parse Tags
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
	postMD, err := readMarkdown(file)
	if err != nil {
		log.Println("error generating template post", err)
		rw.Write([]byte("Internal server error"))
		rw.WriteHeader(http.StatusInternalServerError)
	}
	posts = append(posts, *postMD)
	view.Base(view.Admin(posts)).Render(r.Context(), rw)
}

func readMarkdown(fi *os.File) (*model.Post, error) {
	f, err := os.Open(fi.Name())
	if err != nil {
		log.Printf("error opening file [%s]: %s", fi.Name(), err)
		return nil, err
	}
	content, err := io.ReadAll(f)
	if err != nil {
		log.Printf("error reading all file [%s]: %s", f.Name(), err)
		return nil, err
	}
	var post model.Post
	rest, err := frontmatter.Parse(strings.NewReader(string(content)), &post)
	if err != nil {
		log.Printf("error parsing content [%s]: %s", string(content), err)
		return nil, err
	}

	var buf bytes.Buffer
	err = goldmark.Convert(rest, &buf)
	if err != nil {
		log.Printf("error converting content to goldmark buffer: %s", err)
		return nil, err
	}
	post.Content = view.Unsafe(buf.String())
	if post.Slug == "" {
		post.Slug = strings.Split(f.Name(), ".")[0]
		return nil, err
	}
	return &post, nil
}
