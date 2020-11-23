package events

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/urchincolley/swiss-pair/cmd/api/models"
)

func validateUpsertRequest(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		event := &models.Event{}
		json.NewDecoder(r.Body).Decode(event)

		if event.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "name is required")
			return
		}

		ctx := context.WithValue(r.Context(), models.CtxKey("item"), event)
		r = r.WithContext(ctx)
		next(w, r, p)
	}
}
