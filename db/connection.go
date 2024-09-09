package db

import (
	"database/sql"
	"fmt"

	"github.com/pytsx/api-postgresql/config"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	cfg := config.GetDB()

	strConn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Pass,
		cfg.Database,
	)

	conn, err := sql.Open("postgres", strConn)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()


	return conn, err 
}