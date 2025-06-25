package mongo

import (
	"context"
	"errors"

	"com.int.recipies/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type DataSource struct {
	client     *mongo.Client
	database   string
	collection string
	logger     *zap.Logger
}

func NewDataSource(client *mongo.Client, database string, collection string, logger *zap.Logger) *DataSource {
	return &DataSource{
		client:     client,
		database:   database,
		collection: collection,
		logger:     logger,
	}
}

func (d *DataSource) CreateRecipie(ctx context.Context, recipie models.Recipie) (int, error) {
	collection := d.client.Database(d.database).Collection(d.collection)
	_, err := collection.InsertOne(ctx, recipie)
	if err != nil {
		d.logger.Error("error when creating recipie", zap.Error(err))
		return -1, err
	}
	return recipie.ID, nil
}

func (d *DataSource) GetRecipie(ctx context.Context, id int) (*models.Recipie, error) {
	collection := d.client.Database(d.database).Collection(d.collection)
	filter := bson.M{"id": id}
	var recipie models.Recipie
	err := collection.FindOne(ctx, filter).Decode(&recipie)
	if errors.Is(err, mongo.ErrNoDocuments) {
		d.logger.Error("recipie not found", zap.Int("id", id))
		return nil, err
	} else if err != nil {
		d.logger.Error("error while fetching recipie", zap.Int("id", id), zap.Error(err))
		return nil, err
	}
	return &recipie, nil
}

