package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/FreezeSnail/urlShortener/src/db/sqlite"
	domain "github.com/FreezeSnail/urlShortener/src/http/domain"
	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	db *sql.DB
	q  *sqlite.Queries
}

func NewSQLite(url string) (*SQLite, error) {
	db, err := sql.Open("sqlite3", url)
	if err != nil {
		return nil, fmt.Errorf("opening db: %w", err)
	}

	if _, err := db.Exec(sqlite.Schema); err != nil {
		return nil, fmt.Errorf("executing schema: %w", err)
	}

	return &SQLite{
		db: db,
		q:  sqlite.New(db),
	}, nil
}

func (db *SQLite) Close() {
	db.db.Close()
}

func (db *SQLite) AddUrl(ctx context.Context, url string, short string) (*domain.ShortenURLResponse, error) {
	resp, err := db.q.CreateURL(ctx, sqlite.CreateURLParams{
		Url:        url,
		Shorturl:   short,
		Userid:     sql.NullInt64{Int64: 1, Valid: true},
		Createdate: sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to add url to db: %v", err)
	}

	return &domain.ShortenURLResponse{
		CreatedAt: resp.Createdate.Int64,
		ID:        resp.ID,
		URL:       resp.Url,
		ShortURL:  resp.Shorturl,
		User:      resp.Userid.Int64,
	}, nil
}

func (db *SQLite) GetLongUrl(ctx context.Context, url string) (*domain.ShortURLResponse, error) {
	resp, err := db.q.GetLongURLFromShort(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to add url to db: %v", err)
	}

	return &domain.ShortURLResponse{
		URL: resp,
	}, nil
}

func (db *SQLite) CreateUser(ctx context.Context, user, pass, key string) error {
	params := sqlite.CreateUserParams{
		Name:     user,
		Password: pass,
		Apikey:   key,
	}
	err := db.q.CreateUser(ctx, params)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

func (db *SQLite) GetAPIKey(ctx context.Context, name, hashedPass string) (string, error) {
	params := sqlite.GetAPIKeyParams{
		Name:     name,
		Password: hashedPass,
	}
	resp, err := db.q.GetAPIKey(ctx, params)
	if err != nil {
		return "", fmt.Errorf("failed to get password hash: %v", err)
	}

	return resp, nil
}

func (db *SQLite) GetHashedPassword(ctx context.Context, name string) (string, error) {
	resp, err := db.q.GetHashPassword(ctx, name)
	if err != nil {
		return "", fmt.Errorf("failed to get password hash: %v", err)
	}

	return resp, nil
}
