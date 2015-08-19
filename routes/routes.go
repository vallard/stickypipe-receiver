package routes

import (
	"net/http"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/vallard/stickypipe-receiver/app"
	"github.com/vallard/stickypipe-receiver/handlers"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("views", t.filename)))
	})
	err := t.templ.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// API returns a handler for a set of routes.
func API() http.Handler {

	a := app.New()

	// Setup the file server to serve up static content such as
	// the index.html page.
	//a.TreeMux.NotFoundHandler = http.FileServer(http.Dir("views")).ServeHTTP
	a.TreeMux.NotFoundHandler = homeHandler

	// Initialize the routes for the API binding the route to the
	// handler code for each specified verb.
	a.Handle("GET", "/v1/pipes", handlers.Pipes.List)
	a.Handle("POST", "/v1/pipes", handlers.Pipes.Create)

	return a
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	t := templateHandler{filename: "index.html"}
	t.ServeHTTP(w, r)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	t := templateHandler{filename: "404.html"}
	t.ServeHTTP(w, r)
}
