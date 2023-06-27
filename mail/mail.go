package mail

import "tdd_with_go/mail/server"

type Mail struct {
	server server.MailServer
}

func (m *Mail) SendMail(recipient string, subject string, body string) {
	m.server.Send(recipient, subject, body)
}
