package user

import (
	"github.com/InNOcentos/Doggo-track/backend/internal/account/client/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	db *mongo.Database
}

func NewRepository(db db.Database) *repo {
	return &repo{db: db.Mongo}
}
