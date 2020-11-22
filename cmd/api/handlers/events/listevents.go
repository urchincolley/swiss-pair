package events

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/urchincolley/swiss-pair/cmd/api/handlers"
	"github.com/urchincolley/swiss-pair/cmd/api/models"
	"github.com/urchincolley/swiss-pair/pkg/application"
	"github.com/urchincolley/swiss-pair/pkg/middleware"
)

var ListHandler = handlers.Handler{
	Handle: listEvents,
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
	},
}

func listEvents(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		events := &models.Events{}

		if err := events.List(r.Context(), app); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%e", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		resp, _ := json.Marshal(events)
		w.Write(resp)
	}
}
