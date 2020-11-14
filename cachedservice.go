package cachedservice

import (
	"encoding/json"
	"fmt"
	"log"
)

type CachedService struct {
	dao        *GenericDao
	cache      *GenericCache
	cnv        *GenericConverter
	idProvider func(interface{}) string
}

func NewCachedService(dao *GenericDao, cache *GenericCache, cnv *GenericConverter, idProvider func(interface{}) string) CachedService {
	return CachedService{dao: dao, cache: cache, cnv: cnv, idProvider: idProvider}
}

func (srv CachedService) Get(id string) (interface{}, error) {

	cached, err := (*srv.cache).GetById(id)
	if err == nil && cached != nil {
		return (*srv.cnv).FromBinaryJson(cached)
	} else {
		fromDS, err := (*srv.dao).Get(id)
		if err == nil {
			return (*srv.cnv).FromMap(fromDS)
		}
	}
	return cached, err
}

func (srv CachedService) Update(id string, entity interface{}) (interface{}, error) {

	prevState, err := (*srv.dao).Get(id)
	if err == nil {
		updated, err := (*srv.dao).Update(id, entity)
		if err == nil {
			jsonStr, err := json.Marshal(updated)
			if err == nil {
				_, err := (*srv.cache).Set(id, jsonStr)
				return updated, err
			} else {
				_, err := (*srv.dao).Update(id, prevState)
				return updated, err
			}
		} else {
			return updated, err
		}
	} else {
		return entity, err
	}

}

func (srv CachedService) Create(entity interface{}) (interface{}, error) {
	log.Println("creating a new entity")

	created, err := (*srv.dao).Create(entity)
	if err == nil {
		log.Println("created in the firestore")

		jsonStr, err := json.Marshal(created)
		if err == nil {
			_, err := (*srv.cache).Set(srv.idProvider(created), jsonStr)
			if err != nil {
				log.Println("creation in the redis failed", err)
				created, errDel := (*srv.dao).Delete(srv.idProvider(created))
				if errDel != nil {
					log.Println("deletion error", errDel)
				}
				return created, err

			} else {
				log.Println("created in redis ")
				return created, nil
			}
		} else {
			log.Println("error creating in the firestore")
			return created, err
		}

	} else {
		return created, err
	}

}

func (srv CachedService) Delete(id string) (interface{}, error) {
	deleted, err := (*srv.dao).Delete(id)
	if err == nil {
		_, err := (*srv.cache).Del(id)
		if err == nil {
			return deleted, nil
		} else {
			deleted, _ := (*srv.dao).Update(id, deleted)
			return deleted, err
		}
	} else {
		return deleted, err
	}
}

func (srv CachedService) GetByIds(ids []string) (interface{}, error) {
	fromCache, err := (*srv.cache).GetByIds(ids)
	if err == nil {
		return (*srv.cnv).FromBinaryJsonSlice(fromCache)
	} else {
		fromDs, err := (*srv.dao).GetByIds(ids)
		if err == nil {
			return (*srv.cnv).FromMapSlice(fromDs)
		} else {
			return nil, err
		}
	}
}

func (srv CachedService) GetByString(path, query string) (interface{}, error) {
	iter := (*srv.dao).GetByString(path, query)
	res := make([]interface{}, 0)
	resErr := make([]error, 0)
	for {
		next, err := iter.Next()
		if err != nil {
			log.Println(err)
			break
		}
		data, err := (*srv.cnv).FromMap(next.Data())
		if err == nil {
			res = append(res, data)
		} else {
			resErr = append(resErr, err)
		}

	}
	if len(resErr) == 0 {
		return res, nil
	}
	return res, fmt.Errorf("errors %+q\n", resErr)
}
