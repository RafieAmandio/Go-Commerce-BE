package handlers

import (
	"encoding/json"
	"mooi/library/response"
	"mooi/src/example/dto"
	"mooi/src/example/services"
	"net/http"
)

type ExampleHandlers interface {
	GetExample(http.ResponseWriter, *http.Request)
}

type exampleHandlers struct {
	ExampleService services.ExampleService
}

func NewHttpHandler(exampleService services.ExampleService) ExampleHandlers {
	return &exampleHandlers{
		ExampleService: exampleService,
	}
}

func (handler *exampleHandlers) GetExample(rw http.ResponseWriter, r *http.Request) {

	var exampleRequest dto.ExampleRequestBody
	err := json.NewDecoder(r.Body).Decode(&exampleRequest)
	if err != nil {
		//http.Error(rw, "invalid request body", http.StatusBadRequest)
		response.JsonResponse(rw, http.StatusBadRequest, "invalid request", nil)
		return
	}

	example := handler.ExampleService.GetExample(exampleRequest)
	exampleResponse := dto.ParseFromEntity(example)

	response.JsonResponse(rw, 200, "success", exampleResponse)
}
