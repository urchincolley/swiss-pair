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

func List(gen func() models.Models) HandleFunc {
	return func(app *application.Application) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			defer r.Body.Close()

			ms := gen()

			if err := ms.List(r.Context(), app); err != nil {
				logger.Error.Printf("list error: %s", err.Error())
				w.WriteHeader(HttpStatusFromError(err))
				fmt.Fprintf(w, "%e", err)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			resp, _ := json.Marshal(ms)
			w.Write(resp)
		}
	}
}
