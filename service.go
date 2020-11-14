package cachedservice

import (
	"fmt"
	"log"
)

type Service struct {
	dao        *GenericDao
	cnv        *GenericConverter
	idProvider func(interface{}) string
}

func NewService(dao *GenericDao, cnv *GenericConverter, idProvider func(interface{}) string) *Service {
	return &Service{dao: dao, cnv: cnv, idProvider: idProvider}
}

func (srv Service) Get(id string) (interface{}, error) {

	return (*srv.dao).Get(id)
}

func (srv Service) Update(id string, entity interface{}) (interface{}, error) {

	return (*srv.dao).Update(id, entity)
}

func (srv Service) Create(entity interface{}) (interface{}, error) {
	log.Println("creating a new entity")

	return (*srv.dao).Create(entity)

}

func (srv Service) Delete(id string) (interface{}, error) {
	return (*srv.dao).Delete(id)

}

func (srv Service) GetByIds(ids []string) (interface{}, error) {
	return (*srv.dao).GetByIds(ids)
}

func (srv Service) GetByString(path, query string) (interface{}, error) {
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
