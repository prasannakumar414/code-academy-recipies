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

func (d *DataSource) UpdateRecipie(ctx context.Context, recipie models.Recipie) (error) {
	collection := d.client.Database(d.database).Collection(d.collection)
	filter := bson.M{"id": recipie.ID}
	update := bson.M{"$set": recipie}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		d.logger.Error("error while updating recipie", zap.Int("id", recipie.ID), zap.Error(err))
		return err
	}
	return nil
}
func (d *DataSource) DeleteRecipie(ctx context.Context, id int) (error) {
	collection := d.client.Database(d.database).Collection(d.collection)
	filter := bson.M{"id": id}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		d.logger.Error("error while deleting recipie", zap.Int("id", id), zap.Error(err))
		return err
	}
	return nil
}

func (d *DataSource) SearchRecipies(ctx context.Context, query string, skip int, limit int) ([]models.Recipie, error) {
	collection := d.client.Database(d.database).Collection(d.collection)
	filter := bson.M{"$text": bson.M{"$search": query}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		d.logger.Error("error while searching recipies", zap.Error(err))
		return nil, err
	}
	var recipies []models.Recipie
	err = cursor.All(ctx, &recipies)
	if err != nil {
		d.logger.Error("error while searching recipies", zap.Error(err))
		return nil, err
	}
	return recipies, nil
}

func (d *DataSource) ListRecipies(ctx context.Context, skip int, limit int) ([]models.Recipie, error) {
	collection := d.client.Database(d.database).Collection(d.collection)
	filter := bson.M{}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		d.logger.Error("error while listing recipies", zap.Error(err))
		return nil, err
	}
	var recipies []models.Recipie
	err = cursor.All(ctx, &recipies)
	if err != nil {
		d.logger.Error("error while listing recipies", zap.Error(err))
		return nil, err
	}
	return recipies, nil
}

func (d *DataSource) GetRecipiesCount(ctx context.Context) (int, error) {
	collection := d.client.Database(d.database).Collection(d.collection)
	filter := bson.M{}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		d.logger.Error("error while getting recipies count", zap.Error(err))
		return -1, err
	}
	return int(count), nil
}