package fstore

import (
	"cloud.google.com/go/firestore"
	"context"
)

type FStore struct {
	ctx    context.Context
	client *firestore.Client
}

func NewFStore(ctx context.Context, client *firestore.Client) *FStore {
	return &FStore{ctx: ctx, client: client}
}

func (ds FStore) DS() (*firestore.Client, context.Context) {
	return ds.client, ds.ctx
}


func (ds FStore) Connect() (*firestore.Client, context.Context) {
	return ds.client, ds.ctx
}
