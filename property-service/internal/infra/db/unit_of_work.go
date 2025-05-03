package db

import (
	"context"
	"property_service/internal/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

type UnitOfWork interface {
	Execute(ctx context.Context, fn func(ctx context.Context, uow UnitOfWork) error) error
	GetAmenityRepository() repositories.AmenityRepositoryInterface
	GetCategoryRepository() repositories.CategoryRepositoryInterface
	GetMediaRepository() repositories.MediaRepositoryInterface
	GetPropertyRepository() repositories.PropertyRepositoryInterface
	GetPropertyTypeRepository() repositories.PropertyTypeRepositoryInterface
	GetRoomRepository() repositories.RoomRepositoryInterface
	GetRoomTypeRepository() repositories.RoomTypeRepositoryInterface
	GetRoomPricingRepository() repositories.RoomPricingRepositoryInterface
	GetCityRepository() repositories.CityRepositoryInterface
	GetAreaRepository() repositories.AreaRepositoryInterface
}

type unitOfWork struct {
	client       *mongo.Client
	dbName       string
	repositories map[string]interface{}
}

func NewUnitOfWork(client *mongo.Client, dbName string) UnitOfWork {
	return &unitOfWork{client: client, dbName: dbName, repositories: make(map[string]interface{})}
}

func (u *unitOfWork) GetAmenityRepository() repositories.AmenityRepositoryInterface {
	const key = "AmenityRepository"
	if repo, ok := u.repositories[key]; ok {
		return repo.(repositories.AmenityRepositoryInterface)
	}
	repo := repositories.NewAmenityRepository(u.client.Database(u.dbName).Collection("amenities"))
	u.repositories[key] = repo
	return repo
}

func (u *unitOfWork) GetCategoryRepository() repositories.CategoryRepositoryInterface {
	const key = "CategoryRepository"
	if repo, ok := u.repositories[key]; ok {
		return repo.(repositories.CategoryRepositoryInterface)
	}
	repo := repositories.NewCategoryRepository(u.client.Database(u.dbName).Collection("categories"))
	u.repositories[key] = repo
	return repo
}

func (u *unitOfWork) GetMediaRepository() repositories.MediaRepositoryInterface {
	const key = "MediaRepository"
	if repo, ok := u.repositories[key]; ok {
		return repo.(repositories.MediaRepositoryInterface)
	}
	repo := repositories.NewMediaRepository(u.client.Database(u.dbName).Collection("media"))
	u.repositories[key] = repo
	return repo
}

func (u *unitOfWork) GetPropertyRepository() repositories.PropertyRepositoryInterface {
	const key = "PropertyRepository"
	if repo, ok := u.repositories[key]; ok {
		return repo.(repositories.PropertyRepositoryInterface)
	}
	repo := repositories.NewPropertyRepository(u.client.Database(u.dbName).Collection("properties"))
	u.repositories[key] = repo
	return repo
}

func (u *unitOfWork) GetPropertyTypeRepository() repositories.PropertyTypeRepositoryInterface {
	const key = "PropertyTypeRepository"
	if repo, ok := u.repositories[key]; ok {
		return repo.(repositories.PropertyTypeRepositoryInterface)
	}
	repo := repositories.NewPropertyTypeRepository(u.client.Database(u.dbName).Collection("property_types"))
	u.repositories[key] = repo
	return repo
}

func (u *unitOfWork) GetRoomRepository() repositories.RoomRepositoryInterface {
	const key = "RoomRepository"
	if repo, ok := u.repositories[key]; ok {
		return repo.(repositories.RoomRepositoryInterface)
	}
	repo := repositories.NewRoomRepository(u.client.Database(u.dbName).Collection("rooms"))
	u.repositories[key] = repo
	return repo
}

func (u *unitOfWork) GetRoomTypeRepository() repositories.RoomTypeRepositoryInterface {
	const key = "RoomTypeRepository"
	if repo, ok := u.repositories[key]; ok {
		return repo.(repositories.RoomTypeRepositoryInterface)
	}
	repo := repositories.NewRoomTypeRepository(u.client.Database(u.dbName).Collection("room_types"))
	u.repositories[key] = repo
	return repo
}

func (u *unitOfWork) GetRoomPricingRepository() repositories.RoomPricingRepositoryInterface {
	const key = "RoomPricingRepository"
	if repo, ok := u.repositories[key]; ok {
		return repo.(repositories.RoomPricingRepositoryInterface)
	}
	repo := repositories.NewRoomPricingRepository(u.client.Database(u.dbName).Collection("room_pricing"))
	u.repositories[key] = repo
	return repo
}

func (u *unitOfWork) GetCityRepository() repositories.CityRepositoryInterface {
	const key = "CityRepository"
	if repo, ok := u.repositories[key]; ok {
		return repo.(repositories.CityRepositoryInterface)
	}
	repo := repositories.NewCityRepository(u.client.Database(u.dbName).Collection("cities"))
	u.repositories[key] = repo
	return repo
}

func (u *unitOfWork) GetAreaRepository() repositories.AreaRepositoryInterface {
	const key = "AreaRepository"
	if repo, ok := u.repositories[key]; ok {
		return repo.(repositories.AreaRepositoryInterface)
	}
	repo := repositories.NewAreaRepository(u.client.Database(u.dbName).Collection("areas"))
	u.repositories[key] = repo
	return repo
}

func (uow *unitOfWork) Execute(ctx context.Context, fn func(ctx context.Context, uow UnitOfWork) error) error {
	return uow.client.UseSession(ctx, func(sessionCtx mongo.SessionContext) error {
		if err := sessionCtx.StartTransaction(); err != nil {
			return err
		}

		if err := fn(sessionCtx, uow); err != nil {
			_ = sessionCtx.AbortTransaction(sessionCtx)
			return err
		}

		return sessionCtx.CommitTransaction(sessionCtx)
	})
}
