package shapes

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShapes_Draw(t *testing.T) {
	type testCase struct {
		shapes   Shapes
		expected string
	}

	tests := map[string]testCase{
		"single shape": {
			shapes:   NewShapes(NewConsoleGraphics().WithWriter(new(bytes.Buffer))).Add(&TextBox{Text: "I'm a string"}),
			expected: "I'm a string\n",
		},
		"multiple shapes": {
			shapes:   NewShapes(NewConsoleGraphics().WithWriter(new(bytes.Buffer))).Add(&TextBox{Text: "I'm a string"}).Add(&Rectangle{Height: 1, Width: 5}),
			expected: "I'm a string\nXXXXX\n",
		},
		"no shapes": {
			shapes:   NewShapes(NewConsoleGraphics().WithWriter(new(bytes.Buffer))),
			expected: "",
		},
	}

	for name, rtc := range tests {
		tc := rtc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// Arrange
			buffer := new(bytes.Buffer)
			console := NewConsoleGraphics().WithWriter(buffer)
			NewConsoleGraphics().WithWriter(buffer)
			shapes := Shapes{
				Shapes:   tc.shapes.Shapes,
				Graphics: console,
			}

			// Act
			shapes.Draw()

			// Assert
			output := buffer.String()
			assert.Equal(t, tc.expected, output)
		})
	}
}
