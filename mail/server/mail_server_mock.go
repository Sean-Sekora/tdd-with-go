package server

import "github.com/stretchr/testify/mock"

type MailServerMock struct {
	mock.Mock
}

func NewMailServerMock() *MailServerMock {
	return &MailServerMock{}
}

func (m *MailServerMock) Send(recipient string, subject string, body string) {
	m.Called(recipient, subject, body)
}
