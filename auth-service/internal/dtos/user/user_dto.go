package userdto

type CreateSiteCommand struct {
	Name   string `json:"name" validate:"required"`
	SiteId string `json:"site_id" validate:"required"`
}

type CreateSiteCommandResponse struct {
	SiteId string `json:"site_id"`
}

type UpdateSiteCommand struct {
	Name   string `json:"name" validate:"required"`
	SiteId string `json:"site_id" validate:"required"`
	Guid   string `json:"guid" validate:"required"`
}

type UpdateSiteCommandResponse struct {
	Guid string `json:"guid"`
}

type DeleteSiteCommand struct {
	Guid string `json:"guid" validate:"required"`
}

type DeleteSiteCommandResponse struct {
	Success bool `json:"success"`
}
