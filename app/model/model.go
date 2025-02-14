package model

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type Model struct {
	DBConn *pgx.Conn
}

func (model *Model) ConnectDB(databaseURL string) {
	conn, err := pgx.Connect(context.Background(), databaseURL)

	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	model.DBConn = conn
}
