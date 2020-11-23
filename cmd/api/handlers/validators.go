package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/urchincolley/swiss-pair/cmd/api/models"
)

// Generalized validator for id route params
func IdRequestValidator(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		uid := p.ByName("id")

		id, err := strconv.Atoi(uid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "malformed id")
			return
		}

		ctx := context.WithValue(r.Context(), models.CtxKey("id"), id)
		r = r.WithContext(ctx)
		next(w, r, p)
	}
}
