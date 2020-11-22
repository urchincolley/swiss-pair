package handlers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/urchincolley/swiss-pair/pkg/application"
	"github.com/urchincolley/swiss-pair/pkg/middleware"
)

type Handler struct {
	Handle     func(app *application.Application) httprouter.Handle
	Middleware []middleware.Middleware
}

func (h Handler) Do(app *application.Application) httprouter.Handle {
	return middleware.Chain(h.Handle(app), h.Middleware...)
}
