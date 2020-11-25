package models

import (
	"context"

	"github.com/urchincolley/swiss-pair/pkg/application"
)

// Handlers for endpoints that have no url params
type Models interface {
	List(ctx context.Context, app *application.Application) error
}

// Handlers for endpoints that have one url param
type Model interface {
	PopulateFromContext(context.Context)
	Create(context.Context, *application.Application) error
	Get(context.Context, *application.Application) error
	Update(context.Context, *application.Application) error
	Delete(context.Context, *application.Application) error
}

type DoubleIndexModel interface {
	WithIds(int, int)
	Get(context.Context, *application.Application) error
}
