package models

import (
	"context"

	"github.com/urchincolley/swiss-pair/pkg/application"
)

type Player struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (p *Player) WithId(id int) { p.ID = id }

func (p *Player) Create(ctx context.Context, app *application.Application) error {
	stmt := `
    INSERT INTO players (
      first_name, last_name, email
    ) VALUES ($1, $2, $3)
    RETURNING id
  `
	return app.DB.Client.QueryRowContext(
		ctx, stmt, p.FirstName, p.LastName, p.Email,
	).Scan(&p.ID)
}

func (p *Player) GetById(ctx context.Context, app *application.Application) error {
	stmt := `
    SELECT first_name, last_name, email
    FROM players
    WHERE id = $1
  `
	return app.DB.Client.QueryRowContext(
		ctx, stmt, p.ID,
	).Scan(&p.FirstName, &p.LastName, &p.Email)
}

func (p *Player) Update(ctx context.Context, app *application.Application) error {
	stmt := `
		UPDATE players SET (
			first_name, last_name, email
		) = ($2, $3, $4)
		WHERE id = $1
		RETURNING id
  `
	return app.DB.Client.QueryRowContext(
		ctx, stmt, p.ID, p.FirstName, p.LastName, p.Email,
	).Scan(&p.ID)
}

func (p *Player) Delete(ctx context.Context, app *application.Application) error {
	stmt := `DELETE FROM eventplayers WHERE player_id = $1`
	_, err := app.DB.Client.ExecContext(ctx, stmt, p.ID)
	if err != nil {
		return err
	}

	stmt = `DELETE FROM players WHERE id = $1`
	_, err = app.DB.Client.ExecContext(ctx, stmt, p.ID)
	return err
}

type Players []Player

func (ps *Players) List(ctx context.Context, app *application.Application) error {
	stmt := `
    SELECT id, first_name, last_name, email
    FROM players
  `

	rows, err := app.DB.Client.QueryContext(ctx, stmt)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var p Player
		if err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Email); err != nil {
			return err
		}
		*ps = append(*ps, p)
	}

	return rows.Err()
}

func GenPlayer() SingleIndexModel             { return &Player{} }
func GenPlayers() NoIndexModel                { return &Players{} }
func AsPlayer(i interface{}) SingleIndexModel { return i.(*Player) }
