// Player GET handler

package players

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/urchincolley/swiss-pair/cmd/api/handlers"
	"github.com/urchincolley/swiss-pair/cmd/api/models"
	"github.com/urchincolley/swiss-pair/pkg/application"
	errs "github.com/urchincolley/swiss-pair/pkg/errors"
	"github.com/urchincolley/swiss-pair/pkg/middleware"
)

var GetHandler = handlers.Handler{
	Handle: getPlayer,
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
		validatePlayerIdRequest,
	},
}

func validatePlayerIdRequest(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		uid := p.ByName("id")

		id, err := strconv.Atoi(uid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "malformed id")
			return
		}

		ctx := context.WithValue(r.Context(), models.CtxKey("playerid"), id)
		r = r.WithContext(ctx)
		next(w, r, p)
	}
}

func getPlayer(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		id := r.Context().Value(models.CtxKey("playerid"))
		player := &models.Player{ID: id.(int)}

		if err := player.GetByID(r.Context(), app); err != nil {
			if errors.Is(err, errs.NotFound) {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "no player exists with id %d", id)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%e", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp, _ := json.Marshal(player)
		w.Write(resp)
	}
}
