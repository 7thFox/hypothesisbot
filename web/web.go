package web

import (
	"net/http"

	"github.com/7thFox/hypothesisbot/database"

	"github.com/go-chi/chi"
)

const indexPath = "./bot-web/dist/index.html"

func getIndex() {

}

var db *database.Database

func StartWeb(d database.Database) error {
	db = &d
	r := chi.NewRouter()
	fs := http.StripPrefix("/ctl/", http.FileServer(http.Dir("./static/public")))
	r.Get("/ctl/*", func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hypothesis Bot Control"))
	})
	r.Route("/api", func(r chi.Router) {
		r.Route("/msgcount", MsgCount)
	})

	return http.ListenAndServe(":8080", r)
}
