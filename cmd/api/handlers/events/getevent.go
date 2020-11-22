package events

import (
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

var GetHandler = handlers.Handler{
	Handle: getEvent,
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
		handlers.IdRequestValidator(models.CtxKey("eventid")),
	},
}

func getEvent(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		id := r.Context().Value(models.CtxKey("eventid"))
		event := &models.Event{ID: id.(int)}

		if err := event.GetByID(r.Context(), app); err != nil {
			if errors.Is(err, errs.NotFound) {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "no event exists with id %d", id)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%e", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp, _ := json.Marshal(event)
		w.Write(resp)
	}
}