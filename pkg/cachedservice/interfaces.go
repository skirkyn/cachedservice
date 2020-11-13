package cachedservice



type GenericService interface {
	Get(id string) (interface{}, error)
	Update(id string, entity interface{}) (interface{}, error)
	Create(entity interface{}) (interface{}, error)
	Delete(id string) (interface{}, error)
	GetByIds(ids []string) (interface{}, error)
	GetByString(path, query string) (interface{}, error)
}

