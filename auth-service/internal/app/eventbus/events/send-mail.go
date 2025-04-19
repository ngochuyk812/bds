package events

type IntegrationEventSendMail struct {
	Body    string
	Subject string
	To      []string
	Cc      []string
}

func (e *IntegrationEventSendMail) Key() string {
	return "sendmail"
}
