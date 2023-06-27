package mail

import (
	"github.com/stretchr/testify/mock"
	"tdd_with_go/mail/server"
	"testing"
)

func TestSendMail(t *testing.T) {
	type args struct {
		recipient string
		subject   string
		body      string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Send welcome mail",
			args: args{
				recipient: "test@example.com",
				subject:   "Welcome!",
				body:      "Welcome to your account",
			},
		},
	}
	// Arrange
	mockedMailServer := server.NewMailServerMock()
	mockedMailServer.On("Send", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	mail := Mail{server: mockedMailServer}

	for _, rt := range tests {
		test := rt
		t.Run(test.name, func(t *testing.T) {
			// Act
			mail.SendMail(test.args.recipient, test.args.subject, test.args.body)

			// Assert
			mockedMailServer.AssertCalled(t, "Send", test.args.recipient, test.args.subject, test.args.body)
			mockedMailServer.AssertNumberOfCalls(t, "Send", 1)
		})
	}
}
