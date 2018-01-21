package web

import (
	"net/http"

	"github.com/go-chi/chi"
)

const indexPath = "./bot-web/dist/index.html"

func getIndex() {

}

func StartWeb() error {
	r := chi.NewRouter()
	fs := http.StripPrefix("/", http.FileServer(http.Dir("./bot-web/dist")))
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Foobar"))
	})

	return http.ListenAndServe(":8080", r)
}
