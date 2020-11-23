package models

import (
	"context"
	"database/sql"
	"errors"
	"strings"

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

	err := app.DB.Client.QueryRowContext(ctx, stmt, e.Name).Scan(&e.ID)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return err
		}
		return err
	}

	return nil
}

func (e *Event) GetById(ctx context.Context, app *application.Application) error {
	stmt := `
    SELECT name
		FROM events
    WHERE id = $1
  `

	err := app.DB.Client.QueryRowContext(ctx, stmt, e.ID).Scan(&e.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return err
			//return errs.NotFound("event not found")
		}
		return err
	}

	return nil
}

func (e *Event) Update(ctx context.Context, app *application.Application) error {
	stmt := `
		UPDATE events
		SET name = $2 
		WHERE id = $1
		RETURNING id
  `

	res, err := app.DB.Client.ExecContext(
		ctx, stmt, e.ID, e.Name,
	)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return err
			//return errs.Conflict(fmt.Sprintf("event with name %s already exists", e.Name))
		}
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return err
	}

	return nil
}

func (e *Event) Delete(ctx context.Context, app *application.Application) error {
	stmt := `
		DELETE FROM events
    WHERE id = $1
  `

	res, err := app.DB.Client.ExecContext(ctx, stmt, e.ID)
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errors.New("event not found")
	}

	return nil
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

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func GenEvent() Model             { return &Event{} }
func GenEvents() Models           { return &Events{} }
func AsEvent(i interface{}) Model { return i.(*Event) }
