package cachedservice

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

type FStore struct {
	ctx    context.Context
	client *firestore.Client
}

func NewContext() context.Context {
	return context.Background()
}

func NewFStore(ctx context.Context, projectId string) *FStore {
	var client *firestore.Client
	var err error

	client, err = firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatal("can't initialize the firestore client", err)
	}
	log.Println("DS initialized")
	return &FStore{ctx: ctx, client: client}
}

func (ds FStore) DS() (*firestore.Client, context.Context) {
	return ds.client, ds.ctx
}


func (ds FStore) Connect() (*firestore.Client, context.Context) {
	return ds.client, ds.ctx
}
