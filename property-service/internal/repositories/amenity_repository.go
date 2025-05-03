package repositories

import (
	"context"
	"errors"
	"property_service/internal/entities"
	db_helper "property_service/internal/infra/db/helpers"

	"github.com/ngochuyk812/building_block/pkg/dtos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AmenityRepositoryInterface interface {
	GetBaseRepo() Repository[entities.Amenity]
	GetByName(ctx context.Context, name string) (*entities.Amenity, error)
	GetAmenitiesPaging(ctx context.Context, name string, page int32, pageSize int32) (*dtos.PagingModel[entities.Amenity], error)
}

type amenityRepository struct {
	base       Repository[entities.Amenity]
	collection *mongo.Collection
}

func NewAmenityRepository(collection *mongo.Collection) AmenityRepositoryInterface {
	return &amenityRepository{
		base:       NewRepository[entities.Amenity](collection),
		collection: collection,
	}
}

func (r *amenityRepository) GetBaseRepo() Repository[entities.Amenity] {
	return r.base
}

func (r *amenityRepository) GetByName(ctx context.Context, name string) (*entities.Amenity, error) {
	var result entities.Amenity
	filter := db_helper.BuildFilter(ctx, bson.M{"name": name})
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}
	return &result, nil
}

func (r *amenityRepository) GetAmenitiesPaging(ctx context.Context, name string, page int32, pageSize int32) (*dtos.PagingModel[entities.Amenity], error) {
	res := &dtos.PagingModel[entities.Amenity]{}
	filter := bson.M{}
	if name != "" {
		filter["name"] = bson.M{"$regex": name, "$options": "i"}
	}
	filter = db_helper.BuildFilter(ctx, filter)

	totalCount, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	cursor, err := r.collection.Find(ctx, filter, &options.FindOptions{
		Skip:  func() *int64 { skip := int64((page - 1) * pageSize); return &skip }(),
		Limit: func() *int64 { limit := int64(pageSize); return &limit }(),
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var amenities []entities.Amenity
	if err := cursor.All(ctx, &amenities); err != nil {
		return nil, err
	}

	res.Items = amenities
	res.Total = int(totalCount)
	res.Page = int(page)
	res.PageSize = int(pageSize)
	return res, nil
}
