package cachedservice

import (
	"cloud.google.com/go/firestore"
	"context"
)

type GenericCache interface {
	GetById(id string) ([]byte, error)
	GetByIds(ids []string) ([]interface{}, error)
	Set(id string, obj interface{}) (string, error)
	Del(id string) (int64, error)
}


type GenericService interface {
	Get(id string) (interface{}, error)
	Update(id string, entity interface{}) (interface{}, error)
	Create(entity interface{}) (interface{}, error)
	Delete(id string) (interface{}, error)
	GetByIds(ids []string) (interface{}, error)
	GetByString(path, query string) (interface{}, error)
}


type GenericConverter interface {
	FromMap(map[string]interface{}) (interface{}, error)
	FromBinaryJson([]byte) (interface{}, error)
	FromMapSlice([]map[string]interface{}) ([]interface{}, error)
	FromBinaryJsonSlice([]interface{}) ([]interface{}, error)
}



type GenericDao interface {
	Get(id string) (map[string]interface{}, error)
	Update(id string, entity interface{}) (interface{}, error)
	Create(entity interface{}) (interface{}, error)
	Delete(id string) (interface{}, error)
	GetByIds(ids []string) ([]map[string]interface{}, error)
	GetByString(path, query string) *firestore.DocumentIterator
}

type Firestore interface {
	Connect() (*firestore.Client, context.Context)
}


