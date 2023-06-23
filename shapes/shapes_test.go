package shapes

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestShapes(t *testing.T) {
	buffer := new(bytes.Buffer)
	console := NewConsoleGraphics().WithWriter(buffer)
	shapes := NewShapes(console).
		Add(&TextBox{Text: "Hello from the Shapes SOLID Demo"}).
		Add(&Rectangle{Height: 1, Width: 5})

	shapes = shapes.Draw()
	output := buffer.String()
	fmt.Print(output)
	if !strings.Contains(output, "Hello from the Shapes SOLID Demo") || !strings.Contains(output, "XXXXX") {
		t.Errorf("")
	}
}
