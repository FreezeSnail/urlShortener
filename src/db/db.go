package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/FreezeSnail/urlShortener/src/db/sqlite"
	domain "github.com/FreezeSnail/urlShortener/src/domain"

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

func (db *SQLite) AddUrl(url string, short string) (*domain.UrlResponse, error) {
	resp, err := db.q.CreateURL(context.TODO(), sqlite.CreateURLParams{
		Url:        url,
		Shorturl:   short,
		Userid:     sql.NullInt64{Int64: 1, Valid: true},
		Createdate: sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to add url to db: %v", err)
	}

	return &domain.UrlResponse{
		CreatedAt: resp.Createdate.Int64,
		Id:        resp.ID,
		Url:       resp.Url,
		ShortURL:  resp.Shorturl,
		User:      resp.Userid.Int64,
	}, nil
}

func (db *SQLite) GetLongUrl(url string) (*domain.ShortUrlResponse, error) {
	resp, err := db.q.GetLongURLFromShort(context.TODO(), url)
	if err != nil {
		return nil, fmt.Errorf("failed to add url to db: %v", err)
	}

	return &domain.ShortUrlResponse{
		Url: resp,
	}, nil
}
