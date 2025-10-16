package database

import (
	"context"
	"url_shortener/internal/model"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	db *sqlx.DB
}

// Конструктор, создаёт репозиторий поверх sqlx.DB
func NewRepo(db *sqlx.DB) *Database {
	return &Database{db: db}
}

// Создание новой записи в таблице shorten
func (d *Database) Create(ctx context.Context, u *model.URL) error {
	query := `
		INSERT INTO shorten (original, short, visits)
		VALUES ($1, $2, $3)
		RETURNING id;
	`

	// QueryRowxContext возвращает строку, из которой нужно считать значение
	if err := d.db.QueryRowxContext(ctx, query, u.Original, u.Short, u.Visits).Scan(&u.ID); err != nil {
		return err
	}

	return nil
}

func (d *Database) GetOriginal(ctx context.Context, short string) (string, error) {
	var original string
	query := `SELECT original FROM shorten WHERE short = $1`
	if err := d.db.GetContext(ctx, &original, query, short); err != nil {
		return "", err
	}
	return original, nil
}
