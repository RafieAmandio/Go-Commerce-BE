package db

import (
	"database/sql"
	"log"
	"mooi/config"

	_ "github.com/lib/pq"
)

type PostgreSQLDB interface {
	ConnectDatabase() *sql.DB
}

type postgresDB struct {
	db *sql.DB
}

func NewPostgreSQLClient() PostgreSQLDB {
	dbConfig := config.Get()

	db, err := sql.Open("postgres", dbConfig.PostgreSQLConnectionString())
	if err != nil {
		log.Fatal("error while creating new PostgreSQL connection. error: ", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("error while pinging the database. error: ", err)
		return nil
	}

	return &postgresDB{
		db: db,
	}
}

func (p *postgresDB) ConnectDatabase() *sql.DB {
	return p.db
}
