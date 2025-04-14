package queries_site

import "github.com/ngochuyk812/building_block/pkg/dtos"

type FetchSitesQuery struct {
	*dtos.PagingRequest
	Name   string
	SiteId string
}
