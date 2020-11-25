package models

import (
	"context"

	"github.com/urchincolley/swiss-pair/pkg/application"
)

type Event struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (e *Event) PopulateFromContext(ctx context.Context) {
	e.ID = ctx.Value(CtxKey("id")).(int)
}

func (e *Event) Create(ctx context.Context, app *application.Application) error {
	stmt := `
    INSERT INTO events (
      name
    ) VALUES ($1)
    RETURNING id
  `
	return app.DB.Client.QueryRowContext(ctx, stmt, e.Name).Scan(&e.ID)
}

func (e *Event) Get(ctx context.Context, app *application.Application) error {
	stmt := `
    SELECT name
		FROM events
    WHERE id = $1
  `
	return app.DB.Client.QueryRowContext(ctx, stmt, e.ID).Scan(&e.Name)
}

func (e *Event) Update(ctx context.Context, app *application.Application) error {
	stmt := `
		UPDATE events
		SET name = $2 
		WHERE id = $1
		RETURNING id
  `
	return app.DB.Client.QueryRowContext(ctx, stmt, e.ID, e.Name).Scan(&e.ID)
}

func (e *Event) Delete(ctx context.Context, app *application.Application) error {
	stmt := `DELETE FROM eventplayers WHERE event_id = $1`
	_, err := app.DB.Client.ExecContext(ctx, stmt, e.ID)
	if err != nil {
		return err
	}

	stmt = `DELETE FROM events WHERE id = $1`
	_, err = app.DB.Client.ExecContext(ctx, stmt, e.ID)
	return err
}

type Events []Event

func (es *Events) List(ctx context.Context, app *application.Application) error {
	stmt := `
    SELECT id, name
    FROM events
  `

	rows, err := app.DB.Client.QueryContext(ctx, stmt)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Name); err != nil {
			return err
		}
		*es = append(*es, e)
	}

	return rows.Err()
}

func GenEvent() Model             { return &Event{} }
func GenEvents() Models           { return &Events{} }
func AsEvent(i interface{}) Model { return i.(*Event) }
