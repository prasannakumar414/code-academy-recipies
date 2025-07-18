package main

import (
	"context"

	"com.int.recipies/datasource/mongo"
	"com.int.recipies/datasource/postgresql"
	"com.int.recipies/http"
	"com.int.recipies/http/handlers"
	"com.int.recipies/service"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	mongoClient, err := mongo.Connect("mongodb://localhost:27017")
	if err != nil {
		logger.Error("could not initialize mongo", zap.Error(err))
	}
	logger.Info("initialized mongo")
	defer func() {
		if err = mongoClient.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		logger.Error("Unable to connect to database", zap.Error(err))
	}
	defer conn.Close(context.Background())

	recipiesDB := postgresql.NewRecipiesDB(conn, logger)

	_ = mongo.NewDataSource(mongoClient, "recipies-db", "recipies", logger)
	recipiesService := service.NewService(recipiesDB, logger)
	recipieHandler := handlers.NewRecipieHandler(recipiesService)
	handlers := http.Handlers{
		RecipieHandlers: *recipieHandler,
	}
	server := http.NewServer(handlers)
	server.ListenAndServe("localhost:8080")
}
