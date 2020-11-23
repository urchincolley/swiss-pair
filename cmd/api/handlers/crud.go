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

func CreateItem(as func(interface{}) models.Model) HandleFunc {
	return func(app *application.Application) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			defer r.Body.Close()

			m := as(r.Context().Value(models.CtxKey("item")))

			if err := m.Create(r.Context(), app); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
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

func GetItem(gen func() models.Model) HandleFunc {
	return func(app *application.Application) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			defer r.Body.Close()

			id := r.Context().Value(models.CtxKey("id")).(int)
			m := gen()
			m.WithId(id)

			if err := m.GetById(r.Context(), app); err != nil {
				logger.Error.Printf("get error: %s", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
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

func ListItems(gen func() models.Models) HandleFunc {
	return func(app *application.Application) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			defer r.Body.Close()

			ms := gen()

			if err := ms.List(r.Context(), app); err != nil {
				logger.Error.Printf("list error: %s", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "%e", err)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			resp, _ := json.Marshal(ms)
			w.Write(resp)
		}
	}
}

func UpdateItem(as func(interface{}) models.Model) HandleFunc {
	return func(app *application.Application) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			defer r.Body.Close()

			m := as(r.Context().Value(models.CtxKey("item")))
			id := r.Context().Value(models.CtxKey("id")).(int)
			m.WithId(id)

			if err := m.Update(r.Context(), app); err != nil {
				logger.Error.Printf("updateEvent error: %s", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
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

func DeleteItem(gen func() models.Model) HandleFunc {
	return func(app *application.Application) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			defer r.Body.Close()

			id := r.Context().Value(models.CtxKey("id")).(int)
			m := gen()
			m.WithId(id)

			if err := m.Delete(r.Context(), app); err != nil {
				logger.Error.Printf("delete error: %s", err.Error())
				//w.WriteHeader(err.StatusCode())
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "%e", err)
				return
			}

			w.WriteHeader(http.StatusNoContent)
		}
	}
}
