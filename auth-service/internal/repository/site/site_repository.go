package repositorysite

import (
	"auth_service/internal/domain/site"
	"context"
	"database/sql"
	"time"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/ngochuyk812/building_block/pkg/dtos"
	"golang.org/x/sync/errgroup"
)

type SiteRepository interface {
	CreateSite(ctx context.Context, arg site.CreateSiteParams) error
	GetSiteById(ctx context.Context, id int32) (site.Site, error)
	GetSiteByGuid(ctx context.Context, guid string) (site.Site, error)
	UpdateSiteById(ctx context.Context, arg site.UpdateSiteByIdParams) error
	UpdateSiteByGuid(ctx context.Context, arg site.UpdateSiteByGuidParams) error
	DeleteSiteById(ctx context.Context, id int32) error
	DeleteSiteByGuid(ctx context.Context, guid string) error
	GetSitesPaging(ctx context.Context, page, size int32) (*dtos.PagingModel[site.Site], error)
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

func (r *siteRepository) CreateSite(ctx context.Context, arg site.CreateSiteParams) error {
	arg.Createdat = time.Now().Unix()
	guid, err := guid.NewV4()
	if err != nil {
		return err
	}
	arg.Guid = guid.String()

	return r.writeQueries.CreateSite(ctx, arg)
}

func (r *siteRepository) GetSiteById(ctx context.Context, id int32) (site.Site, error) {
	return r.readQueries.GetSiteById(ctx, id)
}

func (r *siteRepository) GetSiteByGuid(ctx context.Context, guid string) (site.Site, error) {
	return r.readQueries.GetSiteByGuid(ctx, guid)
}

func (r *siteRepository) UpdateSiteById(ctx context.Context, arg site.UpdateSiteByIdParams) error {
	return r.writeQueries.UpdateSiteById(ctx, arg)
}

func (r *siteRepository) UpdateSiteByGuid(ctx context.Context, arg site.UpdateSiteByGuidParams) error {
	return r.writeQueries.UpdateSiteByGuid(ctx, arg)
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

func (r *siteRepository) GetSitesPaging(ctx context.Context, page, size int32) (*dtos.PagingModel[site.Site], error) {
	res := &dtos.PagingModel[site.Site]{}

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
			Name:   2,
		})
		if err == nil {
			res.Items = data
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
