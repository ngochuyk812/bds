package repositorysite

import (
	"auth_service/internal/entities"
	repositorybase "auth_service/internal/repository/base"
	"context"

	"github.com/ngochuyk812/building_block/pkg/dtos"
	"gorm.io/gorm"
)

type SiteRepository interface {
	GetSiteBySiteId(ctx context.Context, siteId string) (*entities.Site, error)
	GetSitesPaging(ctx context.Context, page, size int32) (*dtos.PagingModel[*entities.Site], error)
	GetBaseRepository() repositorybase.Repository[entities.Site]
}

type siteRepository struct {
	db   *gorm.DB
	base repositorybase.Repository[entities.Site]
}

func NewSiteRepository(db *gorm.DB) SiteRepository {
	return &siteRepository{
		db:   db,
		base: repositorybase.NewRepository[entities.Site](db),
	}
}
func (r *siteRepository) GetBaseRepository() repositorybase.Repository[entities.Site] {
	return r.base
}

func (r *siteRepository) GetSiteBySiteId(ctx context.Context, siteId string) (*entities.Site, error) {
	var site entities.Site
	err := r.db.WithContext(ctx).Where("siteId = ? AND deleted_at IS NULL", siteId).First(&site).Error
	if err != nil {
		return nil, err
	}
	return &site, nil
}

func (r *siteRepository) GetSitesPaging(ctx context.Context, page, size int32) (*dtos.PagingModel[*entities.Site], error) {
	res := &dtos.PagingModel[*entities.Site]{}

	var sites []*entities.Site
	var total int64

	offset := (page - 1) * size

	err := r.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		Order("created_at DESC").
		Limit(int(size)).
		Offset(int(offset)).
		Find(&sites).Error

	if err != nil {
		return nil, err
	}

	err = r.db.WithContext(ctx).
		Model(&entities.Site{}).
		Where("deleted_at IS NULL").
		Count(&total).Error

	if err != nil {
		return nil, err
	}

	res.Items = sites
	res.Total = int(total)

	return res, nil
}
