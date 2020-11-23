package models

import (
	"context"

	"github.com/urchincolley/swiss-pair/pkg/application"
)

type Model interface {
	WithId(int)
	GetById(context.Context, *application.Application) error
	Delete(context.Context, *application.Application) error
	Create(context.Context, *application.Application) error
	Update(context.Context, *application.Application) error
}

type Models interface {
	List(ctx context.Context, app *application.Application) error
}
