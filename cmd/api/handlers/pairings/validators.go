package pairings

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/urchincolley/swiss-pair/cmd/api/models"
)

func validateRoundRequest(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		rs, ok := r.URL.Query()["round"]
		if !ok || len(rs) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "round is required")
			return
		}

		rd, err := strconv.Atoi(rs[0])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "malformed round number")
			return
		}

		ctx := context.WithValue(r.Context(), models.CtxKey("round"), rd)
		r = r.WithContext(ctx)
		next(w, r, p)
	}
}
