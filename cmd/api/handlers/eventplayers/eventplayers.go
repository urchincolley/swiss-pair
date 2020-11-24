package eventplayers

import (
	"github.com/urchincolley/swiss-pair/cmd/api/handlers"
	"github.com/urchincolley/swiss-pair/cmd/api/models"
	"github.com/urchincolley/swiss-pair/pkg/middleware"
)

var UpdateHandler = handlers.Handler{
	Handle: handlers.Update(models.AsEventPlayers),
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
		handlers.IdRequestValidator,
		validateUpdateRequest,
	},
}

var ListHandler = handlers.Handler{
	Handle: handlers.Get(models.GenEventPlayers),
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
		handlers.IdRequestValidator,
	},
}
