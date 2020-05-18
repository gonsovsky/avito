package db

import (
	"avito/grabber"
	. "avito/shared"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var conStr = "mongodb://mongo-root:passw0rd@" + MainHost

func AllPages() ([]AvitoPage, error) {
	clientOptions := options.Client().ApplyURI(conStr)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	ctx := context.TODO()
	col := client.Database("db").Collection("avito")
	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var Pages []AvitoPage
	for cursor.Next(ctx) {
		var x = AvitoPage{}
		err := cursor.Decode(&x)
		if err != nil {
			fmt.Println("cursor() ERROR:", err)
		} else {
			//x.Id = cursor.Current.Index(0).Value().String()
			//x.Id = strings.Replace(x.Id, "{\"$oid\"", "", -1)
			//x.Id = strings.Replace(x.Id, ":\"", "", -1)
			//x.Id = strings.Replace(x.Id, "\"}", "", -1)
			Pages = append(Pages, x)
		}
	}
	return Pages, nil
}

func NewPage(page AvitoPage) (AvitoPage, error) {
	clientOptions := options.Client().ApplyURI(conStr)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return AvitoPage{}, err
	}
	ctx := context.TODO()
	col := client.Database("db").Collection("avito")
	page.Date = time.Now()
	res, err := col.InsertOne(ctx, page)
	if err != nil {
		return AvitoPage{}, err
	} else {
		return OnePageByObjectId(res.InsertedID)
	}
}

func OnePageByObjectId(id interface{}) (AvitoPage, error) {
	clientOptions := options.Client().ApplyURI(conStr)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return AvitoPage{}, err
	}
	var result = AvitoPage{}
	ctx := context.TODO()
	col := client.Database("db").Collection("avito")
	x := id
	var rx = col.FindOne(ctx, bson.M{"_id": x})
	rx.Decode(&result)
	return result, nil
}

func OnePage(id string) (AvitoPage, error) {
	x, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return AvitoPage{}, err
	}
	return OnePageByObjectId(x)
}

func UpdatePage(id string, page AvitoPage) (AvitoPage, error) {
	clientOptions := options.Client().ApplyURI(conStr)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return AvitoPage{}, err
	}
	ctx := context.TODO()
	col := client.Database("db").Collection("avito")
	x, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return AvitoPage{}, err
	}

	update := bson.D{{"$set",
		page,
	}}

	_, err = col.UpdateOne(ctx, bson.M{"_id": x}, update)
	if err != nil {
		return AvitoPage{}, err
	} else {
		return OnePage(id)
	}
}

func UpdatePageLight(id string, page AvitoLitePage) (AvitoPage, error) {
	clientOptions := options.Client().ApplyURI(conStr)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return AvitoPage{}, err
	}
	ctx := context.TODO()
	col := client.Database("db").Collection("avito")
	x, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return AvitoPage{}, err
	}

	update := bson.D{{"$set",
		page,
	}}

	_, err = col.UpdateOne(ctx, bson.M{"_id": x}, update)
	if err != nil {
		return AvitoPage{}, err
	} else {
		return OnePage(id)
	}
}

func DelPage(id string) error {
	clientOptions := options.Client().ApplyURI(conStr)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	ctx := context.TODO()
	col := client.Database("db").Collection("avito")
	x, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	col.DeleteOne(ctx, bson.M{"_id": x})
	return nil
}

func Parse(id string) (AvitoPage, error) {
	page, err := OnePage(id)
	if err != nil {
		return page, err
	}
	page, err = grabber.Grab(page)
	if err != nil {
		return page, err
	}
	page.Id = ""
	page, err = UpdatePage(id, page)
	if err != nil {
		return page, err
	}
	return page, nil
}
