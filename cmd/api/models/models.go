package models

import (
	"context"

	"github.com/urchincolley/swiss-pair/pkg/application"
)

// Handlers for endpoints that have no url params
type NoIndexModel interface {
	List(ctx context.Context, app *application.Application) error
}

// Handlers for endpoints that have one url param
type SingleIndexModel interface {
	WithId(int)
	Create(context.Context, *application.Application) error
	GetById(context.Context, *application.Application) error
	Update(context.Context, *application.Application) error
	Delete(context.Context, *application.Application) error
}

type DoubleIndexModel interface {
	WithIds(int, int)
	Upsert(context.Context, *application.Application) error
	Drop(context.Context, *application.Application) error
}
