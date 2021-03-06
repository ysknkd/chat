package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

// ---- BEGIN templateHandler

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		templpath := filepath.Join("templates", t.filename)
		t.templ = template.Must(template.ParseFiles(templpath))
	})
	t.templ.Execute(w, nil)
}

// ---- END templateHandler

func main() {
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	// start the chat room
	go r.run()
	// start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
