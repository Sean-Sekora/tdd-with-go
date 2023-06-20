package shapes

import "testing"

func TestShapes_Draw(t *testing.T) {
	console := ConsoleGraphics{}
	shapes := Shapes{Graphics: console}
	shapes.Add(&TextBox{Text: "Hello from the Shapes SOLID Demo"})
	shapes.Add(&Rectangle{Height: 1, Width: 32})
	shapes.Add(&TextBox{Text: "Using the SOLID principles to"})
	shapes.Add(&TextBox{Text: "create an extensible mini-framework."})
	shapes.Add(&TextBox{Text: "Draw shapes as ASCII art."})
	shapes.Add(&TextBox{Text: "Following is a 5 x 3 Rectangle:"})
	shapes.Add(&Rectangle{Height: 3, Width: 5})

	shapes.Draw()
}
