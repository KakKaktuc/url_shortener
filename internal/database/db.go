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
	// 1️⃣ Проверяем, есть ли запись с таким original
	var existing model.URL
	err := d.db.GetContext(ctx, &existing, `
		SELECT id, original, short, visits
		FROM shorten
		WHERE original = $1
	`, u.Original)

	if err == nil {
		// уже есть — возвращаем существующую
		u.ID = existing.ID
		u.Short = existing.Short
		u.Visits = existing.Visits
		return nil
	}

	// 2️⃣ Если GetContext вернул ошибку, проверим, действительно ли это "нет строк"
	// sqlx возвращает ту же ErrNoRows из database/sql, так что можно сделать:
	if err != nil && err.Error() != "sql: no rows in result set" {
		return err // другая ошибка
	}

	// 3️⃣ Добавляем новую запись
	query := `
		INSERT INTO shorten (original, short, visits)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	return d.db.QueryRowxContext(ctx, query, u.Original, u.Short, u.Visits).Scan(&u.ID)
}

func (d *Database) GetOriginal(ctx context.Context, short string) (string, error) {
	var original string
	query := `SELECT original FROM shorten WHERE short = $1`
	if err := d.db.GetContext(ctx, &original, query, short); err != nil {
		return "", err
	}
	return original, nil
}
