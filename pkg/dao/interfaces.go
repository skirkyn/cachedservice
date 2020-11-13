package dao

import "cloud.google.com/go/firestore"

type GenericDao interface {
	Get(id string) (map[string]interface{}, error)
	Update(id string, entity interface{}) (interface{}, error)
	Create(entity interface{}) (interface{}, error)
	Delete(id string) (interface{}, error)
	GetByIds(ids []string) ([]map[string]interface{}, error)
	GetByString(path, query string) *firestore.DocumentIterator
}
