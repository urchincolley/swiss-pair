package events

import (
	"github.com/urchincolley/swiss-pair/cmd/api/handlers"
	"github.com/urchincolley/swiss-pair/cmd/api/models"
	"github.com/urchincolley/swiss-pair/pkg/middleware"
)

var CreateHandler = handlers.Handler{
	Handle: handlers.CreateItem(models.AsEvent),
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
		validateUpsertRequest,
	},
}

var GetHandler = handlers.Handler{
	Handle: handlers.GetItem(models.GenEvent),
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
		handlers.IdRequestValidator,
	},
}

var ListHandler = handlers.Handler{
	Handle: handlers.ListItems(models.GenEvents),
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
	},
}

var UpdateHandler = handlers.Handler{
	Handle: handlers.UpdateItem(models.AsEvent),
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
		handlers.IdRequestValidator,
		validateUpsertRequest,
	},
}

var DeleteHandler = handlers.Handler{
	Handle: handlers.DeleteItem(models.GenEvent),
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
		handlers.IdRequestValidator,
	},
}
