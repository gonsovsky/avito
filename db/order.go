package db

import (
	. "avito/shared"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewOrder(order AvitoOrder) (AvitoOrder, error) {
	clientOptions := options.Client().ApplyURI(conStr)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return AvitoOrder{}, err
	}
	ctx := context.TODO()
	col := client.Database("db").Collection("order")
	order.Date = time.Now()
	_, err = col.InsertOne(ctx, order)
	if err != nil {
		return AvitoOrder{}, err
	} else {
		return AvitoOrder{}, nil
	}
}
