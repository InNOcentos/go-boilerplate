package user

import (
	"context"
	"fmt"
	"time"

	"github.com/InNOcentos/Doggo-track/backend/internal/account/model"
	"github.com/InNOcentos/Doggo-track/backend/internal/account/repository/user/converter"
	"github.com/InNOcentos/Doggo-track/backend/internal/account/repository/user/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const USERS_COLLECTION_NAME = "users"

func (r *repo) Get(ctx context.Context, id string) (*model.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	var user entity.User

	err = r.db.Collection(USERS_COLLECTION_NAME).FindOne(ctx, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}

	return converter.EntityToModel(&user), nil
}

func (r *repo) Create(ctx context.Context, dto *model.CreateUser) (string, error) {
	now := time.Now()

	user := entity.User{
		Name:      dto.Name,
		Surname:   dto.Surname,
		CreatedAt: now,
		UpdatedAt: now,
	}

	fmt.Println(now)

	res, err := r.db.Collection(USERS_COLLECTION_NAME).InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}
