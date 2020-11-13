package fstore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/google/wire"
	"log"
)

func NewContext() context.Context {
	return context.Background()
}
func NewFirestore(ctx context.Context, projectId string) Firestore {

	var client *firestore.Client
	var err error

	client, err = firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatal("can't initialize the firestore client", err)
	}
	log.Println("DS initialized")
	return FStore{ctx, client}

}

var FirestoreProvidersSet = wire.NewSet(NewContext, NewFirestore)