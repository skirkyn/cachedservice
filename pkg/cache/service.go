package cache

import (
	"context"
	"github.com/go-redis/redis"
)

type Service struct {
	cl *redis.Client
	ctx context.Context
}

func NewService(cl *redis.Client, ctx context.Context) *Service {
	return &Service{cl: cl, ctx: ctx}
}


func (srv Service) GetById(id string) ([]byte, error) {

	return srv.cl.Get(srv.ctx, id).Bytes()
}

func (srv Service) GetByIds(ids []string) ([]interface{}, error) {
	return srv.cl.MGet(srv.ctx, ids...).Result()
}

func (srv Service) Set(id string, obj interface{}) (string, error) {
	return srv.cl.Set(srv.ctx, id, obj, 0).Result()
}

func (srv Service) Del(id string) (int64, error) {
	return srv.cl.Del(srv.ctx, id).Result()
}
