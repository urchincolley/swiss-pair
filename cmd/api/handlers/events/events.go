package events

import (
	"github.com/urchincolley/swiss-pair/cmd/api/handlers"
	"github.com/urchincolley/swiss-pair/cmd/api/models"
	"github.com/urchincolley/swiss-pair/pkg/middleware"
)

var ListHandler = handlers.Handler{
	Handle: handlers.List(models.GenEvents),
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
	},
}

var CreateHandler = handlers.Handler{
	Handle: handlers.Create(models.AsEvent),
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
		validateUpsertRequest,
	},
}

var GetHandler = handlers.Handler{
	Handle: handlers.Get(models.GenEvent),
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
		handlers.IdRequestValidator,
	},
}

var UpdateHandler = handlers.Handler{
	Handle: handlers.Update(models.AsEvent),
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
		handlers.IdRequestValidator,
		validateUpsertRequest,
	},
}

var DeleteHandler = handlers.Handler{
	Handle: handlers.Delete(models.GenEvent),
	Middleware: []middleware.Middleware{
		middleware.LogRequest,
		handlers.IdRequestValidator,
	},
}
