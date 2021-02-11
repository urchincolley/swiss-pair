package models

import (
	"context"
	"errors"

	"github.com/urchincolley/swiss-pair/pkg/application"
)

type Pairing struct {
	Table        int `json:"table"`
	FirstPlayer  int `json:"first_player"`
	SecondPlayer int `json:"second_player"`
}

type Pairings struct {
	EventId       int       `json:"event_id"`
	Round         int       `json:"round"`
	EventPairings []Pairing `json:"pairings"`
}

func (ps *Pairings) PopulateFromContext(ctx context.Context) {
	ps.EventId = ctx.Value(CtxKey("id")).(int)
	ps.Round = ctx.Value(CtxKey("round")).(int)
}

func (ps *Pairings) Create(ctx context.Context, app *application.Application) error {
	return errors.New("not implemented")
}

func (ps *Pairings) Get(ctx context.Context, app *application.Application) error {
	stmt := `
		SELECT
			tble,
			first_player,
			second_player
		FROM pairings p
		WHERE event_id = $1 AND rnd = $2`

	rows, err := app.DB.Client.QueryContext(ctx, stmt, ps.EventId, ps.Round)
	if err != nil {
		return err
	}
	defer rows.Close()

	eps := []Pairing{}
	for rows.Next() {
		var p Pairing
		err := rows.Scan(
			&p.Table, p.FirstPlayer, p.SecondPlayer,
		)
		if err != nil {
			return err
		}
		eps = append(eps, p)
	}
	ps.EventPairings = eps

	return rows.Err()
}

func (ps *Pairings) Update(ctx context.Context, app *application.Application) error {
	return errors.New("not implemented")
}

func (ps *Pairings) Delete(ctx context.Context, app *application.Application) error {
	return errors.New("not implemented")
}

func GenPairings() Model { return &Pairings{} }
