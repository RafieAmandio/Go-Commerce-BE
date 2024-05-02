package dto

import "mooi/src/example/models"

type ExampleResponse struct {
	Name string `json:"name"`
}

func ParseFromEntity(data models.Example) ExampleResponse {
	return ExampleResponse{
		Name: data.Name,
	}
}
