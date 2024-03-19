package services

import "github.com/youth/product/pkg/store"

type Item struct {
	Id   int
	Name string
}

func (i *Item) Create() error {
	data := make(map[string]interface{})
	data[i.Name] = i.Name

	return store.CreateItem()
}
