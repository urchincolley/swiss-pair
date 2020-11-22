package events

import (
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

var DeleteHandler = handlers.Handler{
	Handle: deleteEvent,
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
		handlers.IdRequestValidator(models.CtxKey("eventid")),
	},
}

func deleteEvent(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		id := r.Context().Value(models.CtxKey("eventid"))
		event := &models.Event{ID: id.(int)}

		if err := event.Delete(r.Context(), app); err != nil {
			if errors.Is(err, errs.NotFound) {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "no event exists with id %d", id)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%e", err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
