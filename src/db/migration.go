package sqlite

import (
	"database/sql"

	"github.com/FreezeSnail/urlShortener/src/db/sqlite"
)

func RunMigrations(db *sql.DB) error {

	_, err := db.Exec(sqlite.Schema)
	if err != nil {
		return err
	}
}
