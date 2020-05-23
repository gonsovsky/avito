package db

import (
	. "avito/shared"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func AllOrders() ([]AvitoOrder, error) {
	clientOptions := options.Client().ApplyURI(conStr)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	ctx := context.TODO()
	col := client.Database("db").Collection("order")
	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var orders []AvitoOrder
	for cursor.Next(ctx) {
		var x = AvitoOrder{}
		err := cursor.Decode(&x)
		if err != nil {
			fmt.Println("cursor() ERROR:", err)
		} else {
			//x.Id = cursor.Current.Index(0).Value().String()
			//x.Id = strings.Replace(x.Id, "{\"$oid\"", "", -1)
			//x.Id = strings.Replace(x.Id, ":\"", "", -1)
			//x.Id = strings.Replace(x.Id, "\"}", "", -1)
			orders = append(orders, x)
		}
	}
	return orders, nil
}

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
