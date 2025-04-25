package repositorysite

import (
	"auth_service/internal/db/site"
	"auth_service/internal/entity"
	"context"
	"database/sql"
	"time"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/ngochuyk812/building_block/pkg/dtos"
	"golang.org/x/sync/errgroup"
)

type SiteRepository interface {
	CreateSite(ctx context.Context, siteEntity *entity.Site) error
	GetSiteById(ctx context.Context, id int32) (*entity.Site, error)
	GetSiteByGuid(ctx context.Context, guid string) (*entity.Site, error)
	UpdateSite(ctx context.Context, siteEntity *entity.Site) error
	DeleteSiteById(ctx context.Context, id int32) error
	DeleteSiteByGuid(ctx context.Context, guid string) error
	GetSitesPaging(ctx context.Context, page, size int32) (*dtos.PagingModel[*entity.Site], error)
}

type siteRepository struct {
	readQueries  *site.Queries
	writeQueries *site.Queries
}

func NewSiteRepository(readDB, writeDB *sql.DB, tx *sql.Tx) SiteRepository {

	if tx != nil {
		txDB := site.New(tx)
		return &siteRepository{
			readQueries:  txDB,
			writeQueries: txDB,
		}
	}
	return &siteRepository{
		readQueries:  site.New(readDB),
		writeQueries: site.New(writeDB),
	}

}

func (r *siteRepository) CreateSite(ctx context.Context, siteEntity *entity.Site) error {
	guid, err := guid.NewV4()
	if err != nil {
		return err
	}
	siteEntity.Guid = guid.String()

	return r.writeQueries.CreateSite(ctx, site.CreateSiteParams{
		Guid:      siteEntity.Guid,
		Siteid:    siteEntity.Siteid,
		Name:      siteEntity.Name,
		Createdat: siteEntity.Createdat,
	})
}

func (r *siteRepository) GetSiteById(ctx context.Context, id int32) (*entity.Site, error) {
	result, err := r.readQueries.GetSiteById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.Site{
		ID:        result.ID,
		Guid:      result.Guid,
		Siteid:    result.Siteid,
		Name:      result.Name,
		Createdat: result.Createdat,
		Updatedat: result.Updatedat,
		Deletedat: result.Deletedat,
	}, nil
}

func (r *siteRepository) GetSiteByGuid(ctx context.Context, guid string) (*entity.Site, error) {
	result, err := r.readQueries.GetSiteByGuid(ctx, guid)
	if err != nil {
		return nil, err
	}

	return &entity.Site{
		ID:        result.ID,
		Guid:      result.Guid,
		Siteid:    result.Siteid,
		Name:      result.Name,
		Createdat: result.Createdat,
		Updatedat: result.Updatedat,
		Deletedat: result.Deletedat,
	}, nil
}

func (r *siteRepository) UpdateSite(ctx context.Context, siteEntity *entity.Site) error {
	if siteEntity.ID > 0 {
		return r.writeQueries.UpdateSiteById(ctx, site.UpdateSiteByIdParams{
			Siteid:    siteEntity.Siteid,
			Name:      siteEntity.Name,
			Updatedat: siteEntity.Updatedat,
			ID:        siteEntity.ID,
		})
	} else {
		return r.writeQueries.UpdateSiteByGuid(ctx, site.UpdateSiteByGuidParams{
			Siteid:    siteEntity.Siteid,
			Name:      siteEntity.Name,
			Updatedat: siteEntity.Updatedat,
			Guid:      siteEntity.Guid,
		})
	}
}

func (r *siteRepository) DeleteSiteById(ctx context.Context, id int32) error {
	return r.writeQueries.DeleteSiteById(ctx, site.DeleteSiteByIdParams{
		Deletedat: sql.NullInt64{Int64: time.Now().Unix()},
		ID:        id,
	})
}

func (r *siteRepository) DeleteSiteByGuid(ctx context.Context, guid string) error {
	return r.writeQueries.DeleteSiteByGuid(ctx, site.DeleteSiteByGuidParams{
		Deletedat: sql.NullInt64{Int64: time.Now().Unix()},
		Guid:      guid,
	})
}

func (r *siteRepository) GetSitesPaging(ctx context.Context, page, size int32) (*dtos.PagingModel[*entity.Site], error) {
	res := &dtos.PagingModel[*entity.Site]{}

	var (
		limit  = size
		offset = (page - 1) * size
	)
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		var err error
		data, err := r.readQueries.GetSitesPaging(ctx, site.GetSitesPagingParams{
			Limit:  limit,
			Offset: offset,
		})
		if err == nil {
			entityItems := make([]*entity.Site, len(data))
			for i, item := range data {
				entityItems[i] = &entity.Site{
					ID:        item.ID,
					Guid:      item.Guid,
					Siteid:    item.Siteid,
					Name:      item.Name,
					Createdat: item.Createdat,
					Updatedat: item.Updatedat,
					Deletedat: item.Deletedat,
				}
			}
			res.Items = entityItems
		}
		return err
	})
	g.Go(func() error {
		var err error
		total, err := r.readQueries.CountSites(ctx)
		if err == nil {
			res.Total = int(total)
		}
		return err
	})
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return res, nil
}
