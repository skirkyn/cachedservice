package cache


type CacheService interface {
	GetById(id string) ([]byte, error)
	GetByIds(ids []string) ([]interface{}, error)
	Set(id string, obj interface{}) (string, error)
	Del(id string) (int64, error)
}


