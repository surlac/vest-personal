package httpserver

import (
	"fmt"
	"html/template"
	"net/http"
)

type HTTPServer struct {
}

type Page struct {
	Title string
	Body  string
}

func (s *HTTPServer) Start() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	body := fmt.Sprintf("Hi there, %s!", r.URL.Path[1:])
	page := &Page{
		Title: "Title here",
		Body:  body,
	}
	t, err := template.ParseFiles("surlac.org/vest/template/main.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, page)
}
