package models

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/urchincolley/swiss-pair/pkg/application"
)

type EventPlayersRequest struct {
	EventId   int   `json:"event_id"`
	PlayerIds []int `json:"player_ids"`
}

type EventPlayers struct {
	EventId int     `json:"event_id"`
	Players Players `json:"players"`
}

func (e *EventPlayers) PopulateFromContext(ctx context.Context) {
	e.EventId = ctx.Value(CtxKey("id")).(int)
}

func (e *EventPlayers) Create(ctx context.Context, app *application.Application) error {
	return errors.New("not implemented")
}

// List all players for Event
func (e *EventPlayers) Get(ctx context.Context, app *application.Application) error {
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
	vs := make([]string, len(e.Players))
	for i, p := range e.Players {
		vs[i] = fmt.Sprintf("(%d, %d)", e.EventId, p.ID)
	}

	stmt := fmt.Sprintf(`
		INSERT INTO eventplayers (event_id, player_id)
		VALUES %s ON CONFLICT DO NOTHING
	`, strings.Join(vs, ", "))

	_, err := app.DB.Client.ExecContext(ctx, stmt)
	return err
}

func (e *EventPlayers) drop(ctx context.Context, app *application.Application) error {
	vs := make([]string, len(e.Players))
	for i, p := range e.Players {
		vs[i] = fmt.Sprintf("%d", p.ID)
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
	stmt := `
		SELECT p.id, p.first_name, p.last_name, p.email
		FROM eventplayers e JOIN players p ON e.player_id = p.id
		WHERE e.event_id = $1`

	rows, err := app.DB.Client.QueryContext(ctx, stmt, e.EventId)
	if err != nil {
		return err
	}
	defer rows.Close()

	ps := []Player{}
	for rows.Next() {
		var p Player
		if err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Email); err != nil {
			return err
		}
		ps = append(ps, p)
	}
	e.Players = ps

	return rows.Err()
}

func GenEventPlayers() Model             { return &EventPlayers{} }
func AsEventPlayers(i interface{}) Model { return i.(*EventPlayers) }
