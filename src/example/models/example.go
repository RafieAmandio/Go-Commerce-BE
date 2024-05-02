package models

type Example struct {
	Name string
}

func (t *Example) CollectionName() string {
	return "example"
}
