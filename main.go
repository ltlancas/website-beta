package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/russross/blackfriday/v2"
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
	log.Printf("listening at http://localhost:%s", port)
	http.ListenAndServe(":"+port, r)
}

type researchData struct {
	HeaderTitle string
	Link        string
	LinkAction  string
	Posts       []research_post
	ShowDesc    bool
}

func registerHTML(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := researchData{
			HeaderTitle: "Latest Work",
			Link:        "/research",
			LinkAction:  "Go to research",
			Posts:       research_posts[0:2],
		}
		err := tmpl.ExecuteTemplate(w, "home.html", data)
		handleErr(err)
	})
	r.HandleFunc("/research", func(w http.ResponseWriter, r *http.Request) {
		data := researchData{
			HeaderTitle: "Latest Work",
			Link:        "/listed-posts",
			LinkAction:  "See all research",
			Posts:       research_posts[0:2],
			ShowDesc:    true,
		}
		err := tmpl.ExecuteTemplate(w, "research.html", data)
		handleErr(err)
	})
	r.HandleFunc("/listed-posts", func(w http.ResponseWriter, r *http.Request) {
		data := researchData{
			HeaderTitle: "All research",
			Link:        "/research",
			LinkAction:  "Back to summary",
			Posts:       research_posts,
			ShowDesc:    true,
		}
		err := tmpl.ExecuteTemplate(w, "postpage.html", data)
		handleErr(err)
	})
	r.HandleFunc("/research/{id}", func(w http.ResponseWriter, r *http.Request) {
		fname := "public/" + mux.Vars(r)["id"] + ".md"
		blog, err := public.ReadFile(fname)
		if err != nil {
			err := tmpl.ExecuteTemplate(w, "notfound.html", nil)
			handleErr(err)
			return
		}
		output := blackfriday.Run(blog)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		err = tmpl.ExecuteTemplate(w, "post.html", string(output))
		handleErr(err)
	})
	r.HandleFunc("/talks", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "talks.html", &articles)
		handleErr(err)
	})
	r.PathPrefix("/public/").Handler(http.FileServer(http.FS(public)))
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "notfound.html", nil)
		handleErr(err)
	})
}

func handleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
