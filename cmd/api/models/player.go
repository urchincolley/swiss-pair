package models

import (
	"context"
	"database/sql"
	"errors"
	"strings"

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

	err := app.DB.Client.QueryRowContext(
		ctx, stmt, p.FirstName, p.LastName, p.Email,
	).Scan(&p.ID)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return err
		}
		return err
	}

	return nil
}

func (p *Player) GetById(ctx context.Context, app *application.Application) error {
	stmt := `
    SELECT first_name, last_name, email
    FROM players
    WHERE id = $1
  `

	err := app.DB.Client.QueryRowContext(
		ctx, stmt, p.ID,
	).Scan(&p.FirstName, &p.LastName, &p.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return err
			//return errs.NotFound
		}
		return err
	}

	return nil
}

func (p *Player) Update(ctx context.Context, app *application.Application) error {
	stmt := `
		UPDATE players SET (
			first_name, last_name, email
		) = ($2, $3, $4)
		WHERE id = $1
		RETURNING id
  `

	res, err := app.DB.Client.ExecContext(
		ctx, stmt, p.ID, p.FirstName, p.LastName, p.Email,
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

func (p *Player) Delete(ctx context.Context, app *application.Application) error {
	stmt := `
		DELETE FROM players
    WHERE id = $1
  `

	res, err := app.DB.Client.ExecContext(ctx, stmt, p.ID)
	if err != nil {
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

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func GenPlayer() Model             { return &Player{} }
func GenPlayers() Models           { return &Players{} }
func AsPlayer(i interface{}) Model { return i.(*Player) }
