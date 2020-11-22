package players

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
	Handle: createPlayer,
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
		validateCreateRequest,
	},
}

func validateCreateRequest(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		player := &models.Player{}
		json.NewDecoder(r.Body).Decode(player)

		var err error
		if player.Email == "" {
			err = fmt.Errorf("email is required; ")
		}
		if player.FirstName == "" {
			err = fmt.Errorf("%wfirst_name is required; ", err)
		}
		if player.LastName == "" {
			err = fmt.Errorf("%wlast_name is required; ", err)
		}
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%e", err)
		}

		ctx := context.WithValue(r.Context(), models.CtxKey("player"), player)
		r = r.WithContext(ctx)
		next(w, r, p)
	}
}

func createPlayer(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		player := r.Context().Value(models.CtxKey("player")).(*models.Player)

		if err := player.Create(r.Context(), app); err != nil {
			if errors.Is(err, errs.AlreadyExists) {
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, "a player with email %s already exists", player.Email)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "player creaton failed")
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		resp, _ := json.Marshal(player)
		w.Write(resp)
	}
}
