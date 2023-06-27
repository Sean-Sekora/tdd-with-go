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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			mail.SendMail(tt.args.recipient, tt.args.subject, tt.args.body)

			// Assert
			mockedMailServer.AssertCalled(t, "Send", tt.args.recipient, tt.args.subject, tt.args.body)
			mockedMailServer.AssertNumberOfCalls(t, "Send", 1)
		})
	}
}
