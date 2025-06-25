package main

import (
	"context"

	"com.int.recipies/datasource/mongo"
	"com.int.recipies/http"
	"com.int.recipies/http/handlers"
	"com.int.recipies/service"
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
	dataSource := mongo.NewDataSource(mongoClient, "recipies-db", "recipies", logger)
	recipiesService := service.NewService(dataSource)
	recipieHandler := handlers.NewRecipieHandler(recipiesService)
	handlers := http.Handlers{
		RecipieHandlers: *recipieHandler,
	}
	server := http.NewServer(handlers)
	server.ListenAndServe("localhost:8080")
}
