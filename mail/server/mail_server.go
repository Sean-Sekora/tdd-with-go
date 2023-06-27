package server

type MailServer interface {
	Send(recipient string, subject string, body string)
}
