package players

import (
	"github.com/urchincolley/swiss-pair/cmd/api/handlers"
	"github.com/urchincolley/swiss-pair/cmd/api/models"
	"github.com/urchincolley/swiss-pair/pkg/middleware"
)

var CreateHandler = handlers.Handler{
	Handle: handlers.CreateItem(models.AsPlayer),
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
		validateUpsertRequest,
	},
}

var GetHandler = handlers.Handler{
	Handle: handlers.GetItem(models.GenPlayer),
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
		handlers.IdRequestValidator,
	},
}

var ListHandler = handlers.Handler{
	Handle: handlers.ListItems(models.GenPlayers),
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
	},
}

var UpdateHandler = handlers.Handler{
	Handle: handlers.UpdateItem(models.AsPlayer),
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
		handlers.IdRequestValidator,
		validateUpsertRequest,
	},
}

var DeleteHandler = handlers.Handler{
	Handle: handlers.DeleteItem(models.GenPlayer),
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
		handlers.IdRequestValidator,
	},
}
