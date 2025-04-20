package events

import (
	"context"
	"fmt"
	"sender_service/internal/infra"

	"github.com/ngochuyk812/building_block/infrastructure/eventbus"
	"github.com/ngochuyk812/building_block/infrastructure/helpers"
)

type IntegrationEventSendMail struct {
	Body    string
	Subject string
	To      []string
	Cc      []string
}

func (e *IntegrationEventSendMail) Key() string {
	return "sendmail"
}

func NewIntegrationEventSendMailHandler(cabin infra.Cabin) eventbus.IntegrationEventHandler {
	return &integrationEventSendMailHandler{
		cabin: cabin,
	}
}

type integrationEventSendMailHandler struct {
	cabin infra.Cabin
}

func (e *integrationEventSendMailHandler) NewEvent() eventbus.IntegrationEvent {
	return &IntegrationEventSendMail{}
}
func (e *integrationEventSendMailHandler) Handle(ctx context.Context, event eventbus.IntegrationEvent) error {
	sendMailEvent, oke := event.(*IntegrationEventSendMail)
	if oke == false {
		e.cabin.GetInfra().GetLogger().Error(fmt.Sprintf("cannot parse event to %v", event))
		return nil
	}
	userContext, _ := helpers.AuthContext(ctx)
	e.cabin.GetInfra().GetLogger().Info(fmt.Sprintf("User: %s\n", userContext.UserName))

	e.cabin.GetInfra().GetLogger().Info(fmt.Sprintf("To: %s\nSubject: %s\nBody: %s\n", sendMailEvent.To, sendMailEvent.Subject, sendMailEvent.Body))
	return nil
}
