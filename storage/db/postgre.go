package db

import (
	"context"
	"log"
	"mooi/config"

	"github.com/jackc/pgx/v4"
)

type PostgreSQLDB interface {
	ConnectDatabase() *pgx.Conn
}

type postgresDB struct {
	conn *pgx.Conn
}

func NewPostgreSQLClient() PostgreSQLDB {
	dbConfig := config.Get()

	conn, err := pgx.Connect(context.Background(), dbConfig.PostgreSQLConnectionString())
	if err != nil {
		log.Fatal("error while creating new PostgreSQL connection. error: ", err)
		return nil
	}

	return &postgresDB{
		conn: conn,
	}
}

func (p *postgresDB) ConnectDatabase() *pgx.Conn {
	return p.conn
}
