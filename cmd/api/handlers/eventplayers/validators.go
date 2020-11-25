package eventplayers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/urchincolley/swiss-pair/cmd/api/models"
)

func validateUpdateRequest(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		epr := &models.EventPlayersRequest{}
		json.NewDecoder(r.Body).Decode(epr)

		if len(epr.PlayerIds) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "players are required")
			return
		}

		ps := make([]models.Player, len(epr.PlayerIds))
		for i, pid := range epr.PlayerIds {
			ps[i] = models.Player{ID: pid}
		}

		ctx := context.WithValue(r.Context(), models.CtxKey("method"), r.Method)
		ctx = context.WithValue(
			ctx, models.CtxKey("item"),
			models.EventPlayers{EventId: epr.EventId, Players: ps},
		)
		r = r.WithContext(ctx)
		next(w, r, p)
	}
}
