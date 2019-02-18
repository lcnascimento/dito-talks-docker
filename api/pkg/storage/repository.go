package storage

import (
	"context"
	"errors"
	"os"
	"time"

	"talk-docker/pkg"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var errNotImplemented = errors.New("Funcionalidade ainda não implementada")

// Repository  ...
type Repository struct {
	db *mongo.Database
}

// NewRepository  ...
func NewRepository() (*Repository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, os.Getenv("MONGO_ENDPOINT"))

	if err != nil {
		return nil, err
	}

	db, err := client.Database(os.Getenv("MONGO_DATABASE")), nil

	if err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}

// ListTalks  ...
func (repo Repository) ListTalks() (*[]pkg.Talk, error) {
	c, err := repo.db.Collection("talks").Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	talks := []pkg.Talk{}

	for c.Next(context.Background()) {
		talk := &pkg.Talk{}
		if err := c.Decode(talk); err != nil {
			return nil, err
		}

		talks = append(talks, *talk)
	}

	return &talks, nil
}

// CreateTalk  ...
func (repo Repository) CreateTalk(t pkg.Talk) error {
	_, err := repo.db.Collection("talks").InsertOne(context.Background(), t)
	return err
}
