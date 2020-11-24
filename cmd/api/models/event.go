package models

import (
	"context"

	"github.com/urchincolley/swiss-pair/pkg/application"
)

type Event struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (e *Event) WithId(id int) { e.ID = id }

func (e *Event) Create(ctx context.Context, app *application.Application) error {
	stmt := `
    INSERT INTO events (
      name
    ) VALUES ($1)
    RETURNING id
  `
	return app.DB.Client.QueryRowContext(ctx, stmt, e.Name).Scan(&e.ID)
}

func (e *Event) GetById(ctx context.Context, app *application.Application) error {
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
	stmt := `
		DELETE FROM events
    WHERE id = $1
		RETURNING id
  `
	return app.DB.Client.QueryRowContext(ctx, stmt, e.ID).Scan(&e.ID)

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

func GenEvent() SingleIndexModel             { return &Event{} }
func GenEvents() NoIndexModel                { return &Events{} }
func AsEvent(i interface{}) SingleIndexModel { return i.(*Event) }
