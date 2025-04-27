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
