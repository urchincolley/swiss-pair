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
	return e.list(ctx, app)
}

// Add the provided set of EventPlayers
func (e *EventPlayers) Update(ctx context.Context, app *application.Application) error {
	var err error
	if ctx.Value(CtxKey("method")).(string) == "PATCH" {
		err = e.add(ctx, app)
	} else {
		err = e.drop(ctx, app)
	}
	if err != nil {
		return err
	}

	return e.list(ctx, app)
}

// Drop also uses Update
func (e *EventPlayers) Delete(ctx context.Context, app *application.Application) error {
	return errors.New("not implemented")
}

func (e *EventPlayers) add(ctx context.Context, app *application.Application) error {
	vs := make([]string, len(e.PlayerIds))
	for i, pid := range e.PlayerIds {
		vs[i] = fmt.Sprintf("(%d, %d)", e.EventId, pid)
	}

	stmt := fmt.Sprintf(`
		INSERT INTO eventplayers (event_id, player_id)
		VALUES %s ON CONFLICT DO NOTHING
	`, strings.Join(vs, ", "))

	_, err := app.DB.Client.ExecContext(ctx, stmt)
	return err
}

func (e *EventPlayers) drop(ctx context.Context, app *application.Application) error {
	vs := make([]string, len(e.PlayerIds))
	for i, pid := range e.PlayerIds {
		vs[i] = fmt.Sprintf("%d", pid)
	}

	stmt := fmt.Sprintf(`
		DELETE FROM eventplayers
		WHERE event_id = $1
		AND player_id IN (%s)
	`, strings.Join(vs, ", "))

	_, err := app.DB.Client.ExecContext(ctx, stmt, e.EventId)
	return err
}

func (e *EventPlayers) list(ctx context.Context, app *application.Application) error {
	stmt := `SELECT player_id FROM eventplayers WHERE event_id = $1`

	rows, err := app.DB.Client.QueryContext(ctx, stmt, e.EventId)
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
