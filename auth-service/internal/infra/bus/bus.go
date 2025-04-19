package bus

import (
	commands_auth "auth_service/internal/app/commands/auth"
	commands_site "auth_service/internal/app/commands/site"
	queries_site "auth_service/internal/app/queries/site"
	"auth_service/internal/infra"

	bus_core "github.com/ngochuyk812/building_block/pkg/mediator/bus"
)

func InjectBus(c infra.Cabin) {
	mediator := c.GetInfra().GetMediator()
	bus_core.RegisterHandler(mediator, commands_site.CreateSiteCommand{}, &commands_site.CreateSiteHandler{
		Cabin: c,
	})
	bus_core.RegisterHandler(mediator, commands_site.UpdateSiteCommand{}, &commands_site.UpdateSiteHandler{
		Cabin: c,
	})
	bus_core.RegisterHandler(mediator, commands_site.DeleteSiteCommand{}, &commands_site.DeleteSiteHandler{
		Cabin: c,
	})
	bus_core.RegisterHandler(mediator, queries_site.FetchSitesQuery{}, &queries_site.FetchSitesHandler{
		Cabin: c,
	})
	bus_core.RegisterHandler(mediator, commands_auth.LoginCommand{}, &commands_auth.LoginHandler{
		Cabin: c,
	})
	bus_core.RegisterHandler(mediator, commands_auth.SignUpCommand{}, &commands_auth.SignUpCommandHandler{
		Cabin: c,
	})
	bus_core.RegisterHandler(mediator, commands_auth.VerifySignUpCommand{}, &commands_auth.VerifySignUpCommandHandler{
		Cabin: c,
	})
}
