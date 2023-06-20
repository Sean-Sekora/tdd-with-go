package shapes

import "fmt"

type ConsoleGraphics struct{}

func (c ConsoleGraphics) DrawText(text string) {
	Print(text)
}

func (c ConsoleGraphics) DrawHorizontalLine(width int) {
	text := ""
	for i := 0; i < width; i++ {
		text += "X"
	}

	Print(text)
}

func Print(text string) {
	fmt.Println(text)
}
