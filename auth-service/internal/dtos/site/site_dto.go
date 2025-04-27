package sitedto

type CreateSiteCommand struct {
	Name   string `json:"name" validate:"required"`
	SiteId string `json:"site_id" validate:"required"`
}

type CreateSiteCommandResponse struct {
	Guid   string `json:"guid"`
	Name   string `json:"name"`
	ID     int64  `json:"id"`
	SiteId string `json:"site_id"`
}

type UpdateSiteCommand struct {
	Name string `json:"name" validate:"required"`
	Guid string `json:"guid" validate:"required"`
}

type DeleteSiteCommand struct {
	Guid string `json:"guid" validate:"required"`
}

type FetchSitesQuery struct {
	Page     int32
	PageSize int32
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
