package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func DbCon() (*sql.DB, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", "root", "admin", "localhost", 1234, "inventory")

	fmt.Println("url", url)
	dbClient, err := sql.Open("postgres", url)
	if err != nil {
		fmt.Printf("cannot connect psql using sql driver, error:, %+v", err)
		return nil, err
	}

	if err = dbClient.Ping(); err != nil {
		fmt.Errorf("ping test failed to psql using sql driver, error: %+v", err)
		return nil, err
	}

	return dbClient, nil
}
