package repositorysite

import (
	"auth_service/internal/entities"
	"context"
	"time"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/ngochuyk812/building_block/pkg/dtos"
	"gorm.io/gorm"
)

type SiteRepository interface {
	CreateSite(ctx context.Context, siteEntity *entities.Site) error
	GetSiteById(ctx context.Context, id int32) (*entities.Site, error)
	GetSiteByGuid(ctx context.Context, guid string) (*entities.Site, error)
	GetSiteBySiteId(ctx context.Context, siteId string) (*entities.Site, error)
	UpdateSite(ctx context.Context, siteEntity *entities.Site) error
	DeleteSiteById(ctx context.Context, id int32) error
	DeleteSiteByGuid(ctx context.Context, guid string) error
	GetSitesPaging(ctx context.Context, page, size int32) (*dtos.PagingModel[*entities.Site], error)
}

type siteRepository struct {
	db *gorm.DB
}

func NewSiteRepository(db *gorm.DB) SiteRepository {
	return &siteRepository{
		db: db,
	}
}

func (r *siteRepository) GetSiteBySiteId(ctx context.Context, siteId string) (*entities.Site, error) {
	var site entities.Site
	err := r.db.WithContext(ctx).Where("siteId = ? AND deleted_at IS NULL", siteId).First(&site).Error
	if err != nil {
		return nil, err
	}
	return &site, nil
}

func (r *siteRepository) CreateSite(ctx context.Context, siteEntity *entities.Site) error {
	guid, err := guid.NewV4()
	if err != nil {
		return err
	}
	siteEntity.Guid = guid.String()
	siteEntity.CreatedAt = time.Now().Unix()

	return r.db.WithContext(ctx).Create(siteEntity).Error
}

func (r *siteRepository) GetSiteById(ctx context.Context, id int32) (*entities.Site, error) {
	var site entities.Site
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&site).Error
	if err != nil {
		return nil, err
	}
	return &site, nil
}

func (r *siteRepository) GetSiteByGuid(ctx context.Context, guid string) (*entities.Site, error) {
	var site entities.Site
	err := r.db.WithContext(ctx).Where("guid = ? AND deleted_at IS NULL", guid).First(&site).Error
	if err != nil {
		return nil, err
	}
	return &site, nil
}

func (r *siteRepository) UpdateSite(ctx context.Context, siteEntity *entities.Site) error {
	siteEntity.UpdatedAt = time.Now().Unix()

	if siteEntity.ID > 0 {
		return r.db.WithContext(ctx).Model(&entities.Site{}).
			Where("id = ?", siteEntity.ID).
			Updates(map[string]interface{}{
				"name":       siteEntity.Name,
				"siteId":     siteEntity.SiteId,
				"updated_at": siteEntity.UpdatedAt,
			}).Error
	} else {
		return r.db.WithContext(ctx).Model(&entities.Site{}).
			Where("guid = ?", siteEntity.Guid).
			Updates(map[string]interface{}{
				"name":       siteEntity.Name,
				"siteId":     siteEntity.SiteId,
				"updated_at": siteEntity.UpdatedAt,
			}).Error
	}
}

func (r *siteRepository) DeleteSiteById(ctx context.Context, id int32) error {
	now := time.Now().Unix()
	return r.db.WithContext(ctx).Model(&entities.Site{}).
		Where("id = ?", id).
		Update("deleted_at", now).Error
}

func (r *siteRepository) DeleteSiteByGuid(ctx context.Context, guid string) error {
	now := time.Now().Unix()
	return r.db.WithContext(ctx).Model(&entities.Site{}).
		Where("guid = ?", guid).
		Update("deleted_at", now).Error
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
