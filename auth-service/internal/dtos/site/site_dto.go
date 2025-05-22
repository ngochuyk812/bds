package sitedto

import "github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"

type CreateSiteCommand struct {
	Name   string `json:"name" validate:"required"`
	SiteId string `json:"site_id" validate:"required"`
}

type CreateSiteCommandResponse struct {
	*statusmsg.StatusMessage
}

type UpdateSiteCommand struct {
	Name string `json:"name" validate:"required"`
	Guid string `json:"guid" validate:"required"`
}
type UpdateSiteCommandResponse struct {
	*statusmsg.StatusMessage
}

type DeleteSiteCommand struct {
	Guid string `json:"guid" validate:"required"`
}
type DeleteSiteCommandResponse struct {
	*statusmsg.StatusMessage
}

type FetchSitesQuery struct {
	Page     int32
	PageSize int32
	Name     string
	SiteId   int32
}

type SiteModel struct {
	ID     int64  `json:"id"`
	Guid   string `json:"guid"`
	Name   string `json:"name"`
	SiteId string `json:"site_id"`
}

type FetchSitesResponse struct {
	Items      []SiteModel
	Total      int
	Page       int32
	PageSize   int32
	TotalPages int32
}
