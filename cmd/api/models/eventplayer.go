package models

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/urchincolley/swiss-pair/pkg/application"
)

type EventPlayers struct {
	EventId   int   `json:"event_id"`
	PlayerIds []int `json:"player_ids"`
}

func (e *EventPlayers) WithId(id int) {
	e.EventId = id
}

func (e *EventPlayers) Create(ctx context.Context, app *application.Application) error {
	return errors.New("not implemented")
}

// List all players for Event
func (e *EventPlayers) GetById(ctx context.Context, app *application.Application) error {
	stmt := fmt.Sprintf(`
    SELECT player_id
    FROM eventplayers
    WHERE event_id = %v
  `, e.EventId)
	return e.execute(ctx, app, stmt)
}

// Add the provided set of EventPlayers
func (e *EventPlayers) Update(ctx context.Context, app *application.Application) error {
	vs := make([]string, len(e.PlayerIds))

	stmt := ""
	if ctx.Value(CtxKey("method")).(string) == "PATCH" {
		for i, pid := range e.PlayerIds {
			vs[i] = fmt.Sprintf("(%d, %d)", e.EventId, pid)
		}

		stmt = fmt.Sprintf(`
			INSERT INTO eventplayers (event_id, player_id)
			VALUES %s ON CONFLICT DO NOTHING
			RETURNING player_id
		`, strings.Join(vs, ", "))
	} else {
		for i, pid := range e.PlayerIds {
			vs[i] = fmt.Sprintf("%d", pid)
		}

		stmt = fmt.Sprintf(`
			DELETE FROM eventplayers
			WHERE event_id = %v
			AND player_id IN (%s)
			RETURNING player_id
		`, e.EventId, strings.Join(vs, ", "))
	}

	return e.execute(ctx, app, stmt)
}

// Drop also uses Update
func (e *EventPlayers) Delete(ctx context.Context, app *application.Application) error {
	return errors.New("not implemented")
}

func (e *EventPlayers) execute(ctx context.Context, app *application.Application, stmt string) error {
	fmt.Printf("\n\n**** execute stmt %v ****\n\n", stmt)
	rows, err := app.DB.Client.QueryContext(ctx, stmt)
	if err != nil {
		return err
	}
	defer rows.Close()

	pids := []int{}
	for rows.Next() {
		var pid int
		if err := rows.Scan(&pid); err != nil {
			return err
		}
		pids = append(pids, pid)
	}
	e.PlayerIds = pids

	return rows.Err()
}

func GenEventPlayers() SingleIndexModel             { return &EventPlayers{} }
func AsEventPlayers(i interface{}) SingleIndexModel { return i.(*EventPlayers) }
