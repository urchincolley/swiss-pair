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

func Create(as func(interface{}) models.Model) HandleFunc {
	return func(app *application.Application) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			defer r.Body.Close()

			m := as(r.Context().Value(models.CtxKey("item")))

			if err := m.Create(r.Context(), app); err != nil {
				logger.Error.Printf("get create: %s", err.Error())
				w.WriteHeader(HttpStatusFromError(err))
				fmt.Fprintf(w, "%e", err)
				return
			}

			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			resp, _ := json.Marshal(m)
			w.Write(resp)
		}
	}
}

func Get(gen func() models.Model) HandleFunc {
	return func(app *application.Application) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			defer r.Body.Close()

			m := gen()
			m.PopulateFromContext(r.Context())

			if err := m.Get(r.Context(), app); err != nil {
				logger.Error.Printf("get error: %s", err.Error())
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

func Update(as func(interface{}) models.Model) HandleFunc {
	return func(app *application.Application) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			defer r.Body.Close()

			m := as(r.Context().Value(models.CtxKey("item")))
			m.PopulateFromContext(r.Context())

			if err := m.Update(r.Context(), app); err != nil {
				logger.Error.Printf("update error: %s", err.Error())
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

func Delete(gen func() models.Model) HandleFunc {
	return func(app *application.Application) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			defer r.Body.Close()

			m := gen()
			m.PopulateFromContext(r.Context())

			if err := m.Delete(r.Context(), app); err != nil {
				logger.Error.Printf("delete error: %s", err.Error())
				w.WriteHeader(HttpStatusFromError(err))
				fmt.Fprintf(w, "%e", err)
				return
			}

			w.WriteHeader(http.StatusNoContent)
		}
	}
}
