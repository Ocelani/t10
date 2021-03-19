package main

import (
	"context"

	"github.com/Ocelani/t10/pkg/auto"
	"github.com/Ocelani/t10/pkg/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoRepository struct {
	db *database.MongoDB
}

// NewMongoRepository is a constructor of type Repository.
func NewMongoRepository(mongoURI, collection string) auto.Repository {
	return &MongoRepository{
		database.NewMongoDB(mongoURI, collection),
	}
}

// Create just register a auto.Entity data in database.
func (r *MongoRepository) Create(data auto.Entity) (auto.Entity, error) {
	mg, err := r.db.Collection.InsertOne(
		context.Background(),
		data,
	)
	if err != nil {
		return auto.Entity{}, err
	}
	data.ID = mg.InsertedID.(primitive.ObjectID)

	return data, nil
}

// ReadAll returns the entire data found in MongoRepository collection.
func (r *MongoRepository) ReadAll() ([]auto.Entity, error) {
	cursor, err := r.db.Collection.Find(
		context.Background(),
		bson.M{},
	)
	if err != nil {
		return nil, err
	}
	var (
		v  auto.Entity
		vv []auto.Entity
	)
	for cursor.Next(context.TODO()) {
		err = cursor.Decode(&v)
		if err != nil {
			return nil, err
		}
		vv = append(vv, v)
	}

	return vv, nil
}

// QueryStatus finds and returns the data of a single auto.Entity.
func (r *MongoRepository) ReadAllWithStatus(status string) ([]auto.Entity, error) {
	cursor, err := r.db.Collection.Find(
		context.Background(),
		bson.M{"status": status},
	)
	if err != nil {
		return nil, err
	}
	var (
		v  auto.Entity
		vv []auto.Entity
	)
	for cursor.Next(context.TODO()) {
		err = cursor.Decode(&v)
		if err != nil {
			return nil, err
		}
		vv = append(vv, v)
	}

	return vv, nil
}

// ReadOneWithID finds and returns the data of a single auto.Entity.
func (r *MongoRepository) ReadOneWithID(id string) (auto.Entity, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return auto.Entity{}, err
	}
	mg := r.db.Collection.FindOne(
		context.Background(),
		bson.M{"_id": oID},
	)
	var v auto.Entity
	if err := mg.Decode(&v); err != nil {
		return auto.Entity{}, err
	}

	return v, nil
}

// ReadOneWithName finds and returns the data of a single auto.Entity.
func (r *MongoRepository) ReadOneWithName(name string) (auto.Entity, error) {
	mg := r.db.Collection.FindOne(
		context.Background(),
		bson.M{"name": name},
	)
	var v auto.Entity
	if err := mg.Decode(&v); err != nil {
		return auto.Entity{}, err
	}

	return v, nil
}

// Update searches the planet parameter ID, then, updates its data in database.
func (r *MongoRepository) Update(data auto.Entity) (auto.Entity, error) {
	_, err := r.db.Collection.UpdateOne(
		context.Background(),
		bson.M{"_id": data.ID},
		bson.M{"$set": data},
	)
	if err != nil {
		return auto.Entity{}, err
	}

	return data, nil
}

// Delete the specific auto.Entity data in database with its id as a parameter.
func (r *MongoRepository) Delete(id string) error {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.db.Collection.DeleteOne(
		context.Background(),
		bson.M{"_id": oID},
	)
	if err != nil {
		return err
	}

	return nil
}
