package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

type Text interface {
	fetch(id int) (err error)
	create() (err error)
	update() (err error)
	delete() (err error)
}

type Post struct {
	Db      *sql.DB
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func handleRequest(t Text) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			err = handleGet(w, r, t)
		case "POST":
			err = handlePost(w, r, t)
		case "PUT":
			err = handlePut(w, r, t)
		case "DELETE":
			err = handleDelete(w, r, t)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func handleGet(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = post.fetch(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(post, "", "\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	return
}

func handlePut(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	return
}

func main() {
	var err error
	db, err := sql.Open("postgres", "user=gwp, dbname=gwp password=741213 sslmode=disable")
	if err != nil {
		panic(err)
	}
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/post/", handleRequest(&Post{Db: db}))
	server.ListenAndServe()
}
