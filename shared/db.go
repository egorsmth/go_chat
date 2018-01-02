package shared

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func Init(info string) error {
	var err error
	Db, err = sql.Open("postgres", info)
	if err != nil {
		return err
	}

	if err = Db.Ping(); err != nil {
		return err
	}
	return nil
}
