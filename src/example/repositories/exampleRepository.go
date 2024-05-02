package repositories

import (
	"mooi/src/example/models"
	"mooi/storage/db"
)

type ExampleRepositoryInterface interface {
	Find(string) models.Example
}

type exampleRepository struct {
	database db.PostgreSQLDB
}

func NewExampleRepostiory(database db.PostgreSQLDB) ExampleRepositoryInterface {
	return &exampleRepository{
		database: database,
	}
}

func (ep *exampleRepository) Find(name string) models.Example {
	var example models.Example
	// fetch example from db
	// example = ep.database.ConnectDatabase().Collection(example.CollectionName()).FindOne(nil, bson.M{})

	// dummy example
	example.Name = "Hello " + name
	return example
}
