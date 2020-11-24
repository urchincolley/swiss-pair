package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/urchincolley/swiss-pair/cmd/api/models"
	"github.com/urchincolley/swiss-pair/pkg/application"
	"github.com/urchincolley/swiss-pair/pkg/logger"
)

func Upsert(as func(interface{}) models.DoubleIndexModel) HandleFunc {
	return func(app *application.Application) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			defer r.Body.Close()

			m := as(r.Context().Value(models.CtxKey("item")))
			eid := r.Context().Value(models.CtxKey("eid")).(int)
			pid := r.Context().Value(models.CtxKey("pid")).(int)
			m.WithIds(eid, pid)

			if err := m.Upsert(r.Context(), app); err != nil {
				logger.Error.Printf("upsert error: %s", err.Error())
				w.WriteHeader(HttpStatusFromError(err))
				fmt.Fprintf(w, "%e", err)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			resp, _ := json.Marshal(m)
			w.Write(resp)
		}
	}
}

func Drop(gen func() models.DoubleIndexModel) HandleFunc {
	return func(app *application.Application) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			defer r.Body.Close()

			m := gen()
			eid := r.Context().Value(models.CtxKey("eid")).(int)
			pid := r.Context().Value(models.CtxKey("pid")).(int)
			m.WithIds(eid, pid)

			if err := m.Drop(r.Context(), app); err != nil {
				logger.Error.Printf("drop error: %s", err.Error())
				w.WriteHeader(HttpStatusFromError(err))
				fmt.Fprintf(w, "%e", err)
				return
			}

			w.WriteHeader(http.StatusNoContent)
		}
	}
}
