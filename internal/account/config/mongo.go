package config

import (
	"errors"
	"os"
)

const (
	MongoConnectionUrl = "MONGO_CONNECTION_URL"
	DatabaseName       = "account"
)

type MongoConfig interface {
	ConnectionUrl() string
	DatabaseName() string
}

type mongoConfig struct {
	connectionUrl string
	databaseName  string
}

func NewMongoConfig() (MongoConfig, error) {
	connectionUrl := os.Getenv(MongoConnectionUrl)
	if len(connectionUrl) == 0 {
		return nil, errors.New("mongo connection url not found")
	}

	return &mongoConfig{
		connectionUrl: connectionUrl,
		databaseName:  DatabaseName,
	}, nil
}

func (cfg *mongoConfig) ConnectionUrl() string {
	return cfg.connectionUrl
}

func (cfg *mongoConfig) DatabaseName() string {
	return cfg.databaseName
}
