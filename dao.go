package cachedservice

import (
	"cloud.google.com/go/firestore"
)

type Dao struct {
	firestore  Firestore
	collection string
	idConsumer func(string, interface{})
	idProvider func(interface{}) string
}

func NewDao(firestore Firestore, collection string, idConsumer func(string, interface{}), idProvider func(interface{}) string) *Dao {
	return &Dao{firestore: firestore, collection: collection, idConsumer: idConsumer, idProvider: idProvider}
}

func (dao Dao) Get(id string) (map[string]interface{}, error) {
	cl, ctx := dao.firestore.Connect()
	doc, err := cl.Collection(dao.collection).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	} else {
		return doc.Data(), nil
	}
}

func (dao Dao) Update(id string, entity interface{}) (interface{}, error) {
	cl, ctx := dao.firestore.Connect()
	_, err := cl.Collection(dao.collection).Doc(id).Set(ctx, entity)
	return entity, err
}

func (dao Dao) Create(entity interface{}) (interface{}, error) {
	cl, ctx := dao.firestore.Connect()
	id := dao.idProvider(entity)
	if id != "" {
		_, err := cl.Collection(dao.collection).Doc(id).Set(ctx, entity)
		return entity, err
	} else {
		ref := cl.Collection(dao.collection).NewDoc()
		dao.idConsumer(ref.ID, entity)
		_, err := ref.Set(ctx, entity)
		return entity, err
	}

}

func (dao Dao) Delete(id string) (interface{}, error) {
	cl, ctx := dao.firestore.Connect()
	return cl.Collection(dao.collection).Doc(id).Delete(ctx)

}

func (dao Dao) GetByIds(ids []string) ([]map[string]interface{}, error) {
	cl, ctx := dao.firestore.Connect()
	docRefs := make([]*firestore.DocumentRef, 0)
	for _, id := range ids {
		docRefs = append(docRefs, cl.Collection(dao.collection).Doc(id))
	}

	dosSnapshots, err := cl.GetAll(ctx, docRefs)
	if err == nil {
		res := make([]map[string]interface{}, 0)
		for _, doc := range dosSnapshots {
			res = append(res, doc.Data())
		}
		return res, nil
	} else {
		return nil, err
	}

}

func (dao Dao) GetByString(path, query string) *firestore.DocumentIterator {
	cl, ctx := dao.firestore.Connect()
	return cl.Collection(dao.collection).Where(path, ">=", query).Where(path, "<=", query+"\uf8ff").Documents(ctx)
}
