package repositories

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrIdInvalid = "invalid id"
	ErrNotFound  = "not found"
)

type Repository[T any] struct {
	collection *mongo.Collection
}

func (r *Repository[T]) Create(ctx context.Context, doc *T) (*T, error) {
	_, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return nil, fmt.Errorf("failed to insert document: %w", err)
	}
	return doc, nil
}

func (r *Repository[T]) GetAll(ctx context.Context) ([]*T, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get documents: %v", err)
	}
	defer cursor.Close(ctx)

	var results []*T
	_ = cursor.All(ctx, &results)

	return results, nil
}

func (r *Repository[T]) GetByID(ctx context.Context, id string) (*T, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New(ErrIdInvalid)
	}

	var result T
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	return &result, err
}

func (r *Repository[T]) Update(ctx context.Context, id string, updateDoc *T) (*T, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New(ErrIdInvalid)
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": updateDoc}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil || result.MatchedCount == 0 {
		return nil, errors.New(ErrNotFound)
	}

	return updateDoc, nil
}

func (r *Repository[T]) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New(ErrIdInvalid)
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil || result.DeletedCount == 0 {
		return errors.New(ErrNotFound)
	}

	return nil
}
