package main

import (
	"mooi/library/logger"
	"mooi/server"
	"mooi/storage/db"
)

func main() {

	database := db.NewPostgreSQLClient()

	server := server.NewServer()
	server.SetDatabase(database)

	logger.Log.Fatal(server.StartHttpServer())
}
