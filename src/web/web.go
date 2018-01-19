package web

import (
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
)

const indexPath = "./bot-web/dist/index.html"

func getIndex() {

}

func StartWeb() error {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if file, err := ioutil.ReadFile(indexPath); err == nil {
			w.Write(file)
		}
	})
	return http.ListenAndServe(":8888", r)
}
