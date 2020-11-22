package models

import (
	"context"
	"database/sql"
	"errors"

	"github.com/urchincolley/swiss-pair/pkg/application"
	errs "github.com/urchincolley/swiss-pair/pkg/errors"
)

type Player struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type Players []Player

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
		return err
	}

	return nil
}

func (p *Player) GetByID(ctx context.Context, app *application.Application) error {
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
			return errs.NotFound
		}
		return err
	}

	return nil
}

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
		return errs.NotFound
	}

	return nil
}
