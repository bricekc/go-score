package base

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// type générique pour représenter un repository
type Repository[Entity EntityConstraint] struct {
	Collection *mongo.Collection
}

func (repository Repository[Entity]) FindAll(filter ...bson.D) ([]Entity, error) {
    var cur *mongo.Cursor
    var err error

    if len(filter) > 0 {
        cur, err = repository.Collection.Find(context.Background(), filter[0])
    } else {
        cur, err = repository.Collection.Find(context.Background(), bson.D{})
    }

    if err != nil {
        return nil, err
    }

    var results []Entity
    if err = cur.All(context.Background(), &results); err != nil {
		log.Fatal(err)
    }

    return results, nil
}

func (repository Repository[Entity]) InsertOne(entity Entity) error {
	_, err := repository.Collection.InsertOne(context.Background(), entity)

	return err
}

func (repository Repository[Entity]) FindById(id string) (Entity, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	var result Entity
	err := repository.Collection.FindOne(context.Background(), bson.D{{"_id", objectId}}).Decode(&result)
	return result, err
}

func (repository Repository[Entity]) Delete(id string) (*mongo.DeleteResult, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	result, err := repository.Collection.DeleteOne(context.Background(), bson.D{{"_id", objectId}})
	return result, err
}
