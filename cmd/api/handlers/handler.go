package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/urchincolley/swiss-pair/cmd/api/models"
	"github.com/urchincolley/swiss-pair/pkg/application"
	"github.com/urchincolley/swiss-pair/pkg/middleware"
)

type Handler struct {
	Handle     func(app *application.Application) httprouter.Handle
	Middleware []middleware.Middleware
}

func (h Handler) Do(app *application.Application) httprouter.Handle {
	return middleware.Chain(h.Handle(app), h.Middleware...)
}

// Generalized validator for id route params
func IdRequestValidator(ck models.CtxKey) func(httprouter.Handle) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			uid := p.ByName("id")

			id, err := strconv.Atoi(uid)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "malformed id")
				return
			}

			ctx := context.WithValue(r.Context(), ck, id)
			r = r.WithContext(ctx)
			next(w, r, p)
		}
	}
}
