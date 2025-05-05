package repositories

import (
	"context"
	"errors"
	"fmt"
	dtos "property_service/internal/dtos/shared"
	db_helper "property_service/internal/infra/db/helpers"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrIdInvalid = "invalid id"
	ErrNotFound  = "not found"
)

func NewRepository[T any](collection *mongo.Collection) Repository[T] {
	return &repository[T]{
		collection: collection,
	}
}

type Repository[T any] interface {
	Create(ctx context.Context, doc *T) (*T, error)
	GetAll(ctx context.Context) ([]*T, error)
	GetByID(ctx context.Context, id string) (*T, error)
	GetByGuid(ctx context.Context, id string) (*T, error)
	Update(ctx context.Context, id string, updateDoc *T) (*T, error)
	Delete(ctx context.Context, id string) error
	UpdateByGuid(ctx context.Context, guid string, updateDoc *T) (*T, error)
	DeleteByGuid(ctx context.Context, guid string) error
	SearchAdvance(ctx context.Context, query dtos.SearchAdvanceModel) (*dtos.SearchAdvanceResponse[T], error)
}
type repository[T any] struct {
	collection *mongo.Collection
}

func (r *repository[T]) SearchAdvance(ctx context.Context, query dtos.SearchAdvanceModel) (*dtos.SearchAdvanceResponse[T], error) {
	res := &dtos.SearchAdvanceResponse[T]{}
	filter := buildMongoFilter(query.Filters)
	skip := int64(query.StartRow)
	limit := int64(query.EndRow - query.StartRow)
	sort := buildMongoSort(query.Sort)

	opts := options.Find().SetSkip(skip).SetLimit(limit).SetSort(sort)
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []T
	for cursor.Next(ctx) {
		var item T
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		results = append(results, item)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	res.Rows = results
	res.Total = int(count)
	return res, nil
}

func (r *repository[T]) Create(ctx context.Context, doc *T) (*T, error) {
	_, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return nil, fmt.Errorf("failed to insert document: %w", err)
	}
	return doc, nil
}

func (r *repository[T]) GetAll(ctx context.Context) ([]*T, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get documents: %v", err)
	}
	defer cursor.Close(ctx)

	var results []*T
	_ = cursor.All(ctx, &results)

	return results, nil
}

func (r *repository[T]) GetByID(ctx context.Context, id string) (*T, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New(ErrIdInvalid)
	}

	var result T
	filter := db_helper.BuildFilter(ctx, bson.M{"_id": objectID})
	err = r.collection.FindOne(ctx, filter).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	return &result, err
}

func (r *repository[T]) Update(ctx context.Context, id string, updateDoc *T) (*T, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New(ErrIdInvalid)
	}
	updateDocMap := bson.M{}
	data, err := bson.Marshal(updateDoc)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal update document: %w", err)
	}
	err = bson.Unmarshal(data, &updateDocMap)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal update document: %w", err)
	}
	updateDocMap["updated_at"] = time.Now().Unix()
	filter := db_helper.BuildFilter(ctx, bson.M{"_id": objectID})

	update := bson.M{"$set": updateDocMap}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil || result.MatchedCount == 0 {
		return nil, errors.New(ErrNotFound)
	}

	return updateDoc, nil
}

func (r *repository[T]) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New(ErrIdInvalid)
	}

	update := bson.M{"$set": bson.M{"deleted_at": time.Now().Unix()}}
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return fmt.Errorf("failed to update deleted_at: %w", err)
	}
	return nil
}

func (r *repository[T]) GetByGuid(ctx context.Context, guid string) (*T, error) {
	var result T
	filter := db_helper.BuildFilter(ctx, bson.M{"guid": guid})

	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	return &result, err
}

func (r *repository[T]) UpdateByGuid(ctx context.Context, guid string, updateDoc *T) (*T, error) {
	filter := db_helper.BuildFilter(ctx, bson.M{"guid": guid})
	updateDocMap := bson.M{}

	data, err := bson.Marshal(updateDoc)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal update document: %w", err)
	}
	err = bson.Unmarshal(data, &updateDocMap)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal update document: %w", err)
	}
	updateDocMap["updated_at"] = time.Now().Unix()
	update := bson.M{"$set": updateDocMap}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil || result.MatchedCount == 0 {
		return nil, errors.New(ErrNotFound)
	}
	return updateDoc, nil
}

func (r *repository[T]) DeleteByGuid(ctx context.Context, guid string) error {

	update := bson.M{"$set": bson.M{"deleted_at": time.Now().Unix()}}
	_, err := r.collection.UpdateOne(ctx, bson.M{"guid": guid}, update)
	if err != nil {
		return fmt.Errorf("failed to update deleted_at: %w", err)
	}
	return nil
}

func buildMongoFilter(filters map[string]dtos.FilterModel) bson.M {
	mongoFilter := db_helper.BuildFilter(context.Background(), bson.M{})

	for field, cond := range filters {
		switch cond.Type {
		case "equals":
			mongoFilter[field] = cond.Filter
		case "contains":
			mongoFilter[field] = bson.M{
				"$regex":   cond.Filter,
				"$options": "i",
			}
		case "greaterThan":
			mongoFilter[field] = bson.M{"$gt": cond.Filter}
		case "lessThan":
			mongoFilter[field] = bson.M{"$lt": cond.Filter}
		default:
			continue
		}
	}

	return mongoFilter
}

func buildMongoSort(sortModel []dtos.SortModelItem) bson.D {
	sort := bson.D{}
	for _, s := range sortModel {
		order := 1
		if s.Sort == "desc" {
			order = -1
		}
		sort = append(sort, bson.E{Key: s.ColId, Value: order})
	}
	return sort
}
