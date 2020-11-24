// cmd/api/router/router.go

package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/urchincolley/swiss-pair/cmd/api/handlers/eventplayers"
	"github.com/urchincolley/swiss-pair/cmd/api/handlers/events"
	"github.com/urchincolley/swiss-pair/cmd/api/handlers/players"
	"github.com/urchincolley/swiss-pair/pkg/application"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()
	mux.POST("/players", players.CreateHandler.Do(app))
	mux.GET("/players/:id", players.GetHandler.Do(app))
	mux.GET("/players", players.ListHandler.Do(app))
	mux.PUT("/players/:id", players.UpdateHandler.Do(app))
	mux.DELETE("/players/:id", players.DeleteHandler.Do(app))

	mux.POST("/events", events.CreateHandler.Do(app))
	mux.GET("/events/:id", events.GetHandler.Do(app))
	mux.GET("/events", events.ListHandler.Do(app))
	mux.PUT("/events/:id", events.UpdateHandler.Do(app))
	mux.DELETE("/events/:id", events.DeleteHandler.Do(app))

	mux.PATCH("/events/:id/players", eventplayers.UpdateHandler.Do(app))
	mux.GET("/events/:id/players", eventplayers.ListHandler.Do(app))
	mux.DELETE("/events/:id/players", eventplayers.UpdateHandler.Do(app))

	// IS THIS A DB TABLE AT ALL?
	// eid = event id, pid = player id
	//mux.GET("/events/:id/records", records.GetHandler.Do(app)) // standings (one player) // body {"player_id": ""}
	//mux.GET("/events/:id/records", records.ListHandler.Do(app)) // standings

	// could results and pairings can be same table, multiple routes
	//mux.GET("/events/:id/results", records.GetHandler.Do(app)) // results (one table) // body {"round": "", "table": ""}
	//mux.PUT("/events/:id/results", records.UpdateHandler.Do(app)) // results submission (one table)

	// tid = table id, eid = event id, pid = player id
	//mux.GET("/events/:id/pairings", pairings.GetHandler.Do(app)) // body {"round": "", "table": ""}

	//mux.GET("/events/:id/pairings", pairings.ListHandler.Do(app))
	return mux
}
