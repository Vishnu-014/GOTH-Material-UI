package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/Vishnu-014/goth-app/pages"
)

func main() {
    r := chi.NewRouter()

    // Static files
    r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Routes
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        pages.HomePage().Render(r.Context(), w)
    })

    r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("<p>Hello from HTMX!</p>"))
    })

    http.ListenAndServe(":8080", r)
}