package commands_site

import (
	"auth_service/internal/entities"
	"auth_service/internal/infra"
	"auth_service/internal/repository"
	"context"
	"database/sql"
	"time"

	bus_core "github.com/ngochuyk812/building_block/pkg/mediator/bus"
)

type CreateSiteHandler struct {
	Cabin infra.Cabin
}

var _ bus_core.IHandler[CreateSiteCommand, CreateSiteCommandResponse] = (*CreateSiteHandler)(nil)

func (h *CreateSiteHandler) Handle(ctx context.Context, cmd CreateSiteCommand) (CreateSiteCommandResponse, error) {
	res := CreateSiteCommandResponse{}
	err := h.Cabin.GetUnitOfWork().ExecTx(ctx, func(uow repository.UnitOfWork) error {
		siteEntity := &entities.Site{
			Name:      cmd.Name,
			Siteid:    cmd.SiteId,
			Createdat: time.Now().Unix(),
		}
		err := uow.GetSiteRepository().CreateSite(ctx, siteEntity)
		return err
	})

	return res, err
}

type UpdateSiteHandler struct {
	Cabin infra.Cabin
}

var _ bus_core.IHandler[UpdateSiteCommand, UpdateSiteCommandResponse] = (*UpdateSiteHandler)(nil)

func (h *UpdateSiteHandler) Handle(ctx context.Context, cmd UpdateSiteCommand) (UpdateSiteCommandResponse, error) {
	res := UpdateSiteCommandResponse{}
	err := h.Cabin.GetUnitOfWork().ExecTx(ctx, func(uow repository.UnitOfWork) error {
		siteEntity := &entities.Site{
			Guid:      cmd.Guid,
			Name:      cmd.Name,
			Siteid:    cmd.SiteId,
			Updatedat: sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
		}
		err := uow.GetSiteRepository().UpdateSite(ctx, siteEntity)
		return err
	})
	return res, err
}

type DeleteSiteHandler struct {
	Cabin infra.Cabin
}

var _ bus_core.IHandler[DeleteSiteCommand, DeleteSiteCommandResponse] = (*DeleteSiteHandler)(nil)

func (h *DeleteSiteHandler) Handle(ctx context.Context, cmd DeleteSiteCommand) (DeleteSiteCommandResponse, error) {
	res := DeleteSiteCommandResponse{}
	err := h.Cabin.GetUnitOfWork().ExecTx(ctx, func(uow repository.UnitOfWork) error {
		err := uow.GetSiteRepository().DeleteSiteByGuid(ctx, cmd.Guid)
		return err
	})
	return res, err
}
