package commands_site

type CreateSiteCommand struct {
	Name   string
	SiteId string
}
type CreateSiteCommandResponse struct {
}

type UpdateSiteCommand struct {
	Name   string
	SiteId string
	Guid   string
}
type UpdateSiteCommandResponse struct {
}

type DeleteSiteCommand struct {
	Guid string
}
type DeleteSiteCommandResponse struct {
}
