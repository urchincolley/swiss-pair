// cmd/api/router/router.go

package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/urchincolley/swiss-pair/cmd/api/handlers/players"
	"github.com/urchincolley/swiss-pair/pkg/application"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/player/:id", players.GetHandler.Do(app))
	mux.GET("/players", players.ListHandler.Do(app))
	mux.POST("/players", players.CreateHandler.Do(app))
	mux.DELETE("/player/:id", players.DeleteHandler.Do(app))
	return mux
}
