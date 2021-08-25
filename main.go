package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

var (
	//go:embed public
	public embed.FS

	tmpl = template.Must(template.ParseFS(public, "public/views/*"))
)

func main() {
	r := mux.NewRouter()
	registerHTML(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9899"
	}
	log.Println("listening in port", port)
	http.ListenAndServe(":"+port, r)
}

func registerHTML(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "home.html", nil)
		handleErr(err)
	})
	r.HandleFunc("/research", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "research.html", nil)
		handleErr(err)
	})
	r.HandleFunc("/talks", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "talks.html", &articles)
		handleErr(err)
	})
	r.PathPrefix("/public/").Handler(http.FileServer(http.FS(public)))
}

func handleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
