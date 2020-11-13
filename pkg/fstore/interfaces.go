package fstore

import (
	"cloud.google.com/go/firestore"
	"context"
)

type Firestore interface {
	Connect() (*firestore.Client, context.Context)
}


