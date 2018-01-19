package web

import (
	"net/http"

	"github.com/go-chi/chi"
)

func StartWeb() error {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	return http.ListenAndServe(":8888", r)
}
