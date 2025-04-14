package bus

import (
	commands_site "auth_service/internal/app/commands/site"
	queries_site "auth_service/internal/app/queries/site"
	"auth_service/internal/infra"

	bus_core "github.com/ngochuyk812/building_block/pkg/mediator/bus"
)

func InjectBus(c infra.Cabin) {
	bus_core.RegisterHandler(c.GetInfra().GetMediator(), commands_site.CreateSiteCommand{}, &commands_site.CreateSiteHandler{
		Cabin: c,
	})
	bus_core.RegisterHandler(c.GetInfra().GetMediator(), commands_site.UpdateSiteCommand{}, &commands_site.UpdateSiteHandler{
		Cabin: c,
	})
	bus_core.RegisterHandler(c.GetInfra().GetMediator(), commands_site.DeleteSiteCommand{}, &commands_site.DeleteSiteHandler{
		Cabin: c,
	})
	bus_core.RegisterHandler(c.GetInfra().GetMediator(), queries_site.FetchSitesQuery{}, &queries_site.FetchSitesHandler{
		Cabin: c,
	})
}
