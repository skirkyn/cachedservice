package cachedservice

import (
	"context"
	"github.com/go-redis/redis"
	"log"
)

type Cache struct {
	cl  *redis.Client
	ctx context.Context
}

type DbUrl string
type Db int

func NewClient(db Db, url DbUrl) *redis.Client {
	log.Println("connecting to the redis", url, "db", db)
	cl := redis.NewClient(&redis.Options{
		Addr: string(url),
		DB:   int(db),
	})
	log.Println("connected")
	return cl
}

func NewCache(cl *redis.Client, ctx context.Context) *Cache {
	return &Cache{cl: cl, ctx: ctx}
}

func (srv Cache) GetById(id string) ([]byte, error) {

	return srv.cl.Get(id).Bytes()
}

func (srv Cache) GetByIds(ids []string) ([]interface{}, error) {
	return srv.cl.MGet(ids...).Result()
}

func (srv Cache) Set(id string, obj interface{}) (string, error) {
	return srv.cl.Set(id, obj, 0).Result()
}

func (srv Cache) Del(id string) (int64, error) {
	return srv.cl.Del(id).Result()
}
