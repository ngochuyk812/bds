package repositorybase

import (
	"context"
	"time"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/ngochuyk812/building_block/infrastructure/helpers"
	"gorm.io/gorm"
)

type Repository[T any] interface {
	Create(ctx context.Context, entity *T) error
	Update(ctx context.Context, entity *T) error
	DeleteByID(ctx context.Context, id int32) error
	DeleteByGuid(ctx context.Context, guid string) error
	GetByID(ctx context.Context, id int32) (*T, error)
	GetByGuid(ctx context.Context, guid string) (*T, error)
}

type genericRepository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) Repository[T] {
	return &genericRepository[T]{db: db}
}
func (r *genericRepository[T]) getSiteId(ctx context.Context) string {
	userContext, ok := helpers.AuthContext(ctx)
	if ok {
		return userContext.IdSite
	}
	return ""
}

func (r *genericRepository[T]) Create(ctx context.Context, entity *T) error {
	if guidSetter, ok := any(entity).(interface{ SetGuid(string) }); ok {
		newGuid, _ := guid.NewV4()
		guidSetter.SetGuid(newGuid.String())
	}
	if siteIdSetter, ok := any(entity).(interface{ SetSiteId(string) }); ok {
		siteIdSetter.SetSiteId(r.getSiteId(ctx))
	}
	if timeSetter, ok := any(entity).(interface{ SetCreatedAt(int64) }); ok {
		timeSetter.SetCreatedAt(time.Now().Unix())
	}

	return r.db.WithContext(ctx).Create(entity).Error
}

func (r *genericRepository[T]) Update(ctx context.Context, entity *T) error {
	if timeSetter, ok := any(entity).(interface{ SetUpdatedAt(int64) }); ok {
		timeSetter.SetUpdatedAt(time.Now().Unix())
	}
	return r.db.WithContext(ctx).Save(entity).Error
}

func (r *genericRepository[T]) DeleteByID(ctx context.Context, id int32) error {
	now := time.Now().Unix()
	return r.db.WithContext(ctx).
		Model(new(T)).
		Where("id = ?", id).
		Update("deleted_at", now).Error
}

func (r *genericRepository[T]) DeleteByGuid(ctx context.Context, guid string) error {
	now := time.Now().Unix()
	return r.db.WithContext(ctx).
		Model(new(T)).
		Where("guid = ?", guid).
		Update("deleted_at", now).Error
}

func (r *genericRepository[T]) GetByID(ctx context.Context, id int32) (*T, error) {
	var entity T
	siteId := ""
	userContext, ok := helpers.AuthContext(ctx)
	if ok {
		siteId = userContext.IdSite
	}
	err := r.db.WithContext(ctx).
		Where("id = ? AND deleted_at IS NULL AND site_id = ?", id, siteId).
		First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *genericRepository[T]) GetByGuid(ctx context.Context, guid string) (*T, error) {
	var entity T
	err := r.db.WithContext(ctx).
		Where("guid = ? AND deleted_at IS NULL", guid).
		First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
