package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/Vishnu-014/goth-app/pages"
    "strconv"
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

     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "/page?page=1", http.StatusFound)
    })

    http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
        page := 1
        if p := r.URL.Query().Get("page"); p != "" {
            if num, err := strconv.Atoi(p); err == nil {
                page = num
            }
        }

        pages.MainPage(page).Render(r.Context(), w)
    })

    http.ListenAndServe(":8080", r)
}