package events

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/urchincolley/swiss-pair/cmd/api/handlers"
	"github.com/urchincolley/swiss-pair/cmd/api/models"
	"github.com/urchincolley/swiss-pair/pkg/application"
	errs "github.com/urchincolley/swiss-pair/pkg/errors"
	"github.com/urchincolley/swiss-pair/pkg/middleware"
)

var CreateHandler = handlers.Handler{
	Handle: createEvent,
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
		validateCreateRequest,
	},
}

func validateCreateRequest(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		event := &models.Event{}
		json.NewDecoder(r.Body).Decode(event)

		if event.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "name is required")
		}

		ctx := context.WithValue(r.Context(), models.CtxKey("event"), event)
		r = r.WithContext(ctx)
		next(w, r, p)
	}
}

func createEvent(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		event := r.Context().Value(models.CtxKey("event")).(*models.Event)

		if err := event.Create(r.Context(), app); err != nil {
			if errors.Is(err, errs.AlreadyExists) {
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, "a event with name %s already exists", event.Name)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "event creation failed")
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		resp, _ := json.Marshal(event)
		w.Write(resp)
	}
}
