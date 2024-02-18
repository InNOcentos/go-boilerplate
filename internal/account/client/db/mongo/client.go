package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/InNOcentos/Doggo-track/backend/internal/account/client/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type cl struct {
	client *mongo.Client
}

func (c *cl) Database(name string) db.Database {
	return db.Database{Mongo: c.client.Database(name)}
}

func (c *cl) Ping(ctx context.Context) {
	if err := c.client.Ping(ctx, nil); err != nil {
		log.Fatal("Cannot ping mongo")
	}

	fmt.Println("Mongodb ping successfull")
}

func New(ctx context.Context, connectionString string) (*cl, error) {
	opts := options.Client().
		ApplyURI(connectionString).
		SetServerSelectionTimeout(4 * time.Second)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		fmt.Println("err connect mongo")
		return nil, err
	}

	return &cl{client: client}, nil
}
