package services

import (
	"assignment/config"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ctx context.Context
)

var bookCollection *mongo.Collection = config.GetCollection(config.DB, "book")
var promotionCollection *mongo.Collection = config.GetCollection(config.DB, "promotion")
var orderCollection *mongo.Collection = config.GetCollection(config.DB, "order")
