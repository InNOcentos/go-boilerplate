package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Client interface {
	Ping(ctx context.Context)
	Database(name string) Database
}

type Database struct {
	Mongo *mongo.Database
}
