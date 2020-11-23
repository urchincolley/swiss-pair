package players

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
			return
		}

		ctx := context.WithValue(r.Context(), models.CtxKey("item"), player)
		r = r.WithContext(ctx)
		next(w, r, p)
	}
}
