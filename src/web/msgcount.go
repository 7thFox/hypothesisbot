package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func MsgCount(r chi.Router) {
	r.Get("/hour-agg", msgCountResp("hour", true))
	r.Get("/day", msgCountResp("dayOfYear", false))
	r.Get("/day-agg", msgCountResp("dayOfWeek", true))
	r.Get("/week", msgCountResp("week", false))
	r.Get("/week-agg", msgCountResp("week", true))
	r.Get("/month", msgCountResp("month", false))
	r.Get("/month-agg", msgCountResp("month", true))
	r.Get("/year", msgCountResp("year", false))
}

func msgCountResp(by string, aggregate bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m, _ := (*db).MsgCount(by, aggregate)
		marsh, _ := json.Marshal(m)
		w.Header().Set("Content-Type", "application/json")
		w.Write(marsh)
	}
}
