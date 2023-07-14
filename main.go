package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Film struct {
	Title string
	Director string
}

func main() {
	fmt.Println("hello")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		
		film := map[string][]Film{
			"Films": {
				{Title: "Casablanca", Director: "Michael Curtiz"},
				{Title: "Cool Hand Luke", Director: "Stuart Rosenberg"},
				{Title: "Bullitt", Director: "Peter Yates"},
			},
		}

		tmpl.Execute(w, film)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		log.Print("HTMLX request received")
		log.Print(r.Header.Get("HX-Request"))
		time.Sleep(1 * time.Second)

		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		htmlStr := fmt.Sprintf("<li class='list-group-item light-border-subtle'> %s - %s</li>", title, director)
		tmpl, _ := template.New("t").Parse(htmlStr) 
		tmpl.Execute(w, nil)
	}



	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))

}