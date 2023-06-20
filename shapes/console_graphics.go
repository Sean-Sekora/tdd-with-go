package shapes

import "fmt"

func drawText(text string) {
	fmt.Print(text)
}

func drawHorizontalLine(width int) {
	text := ""
	for i := 0; i < width; i++ {
		text += "X"
	}

	fmt.Print(text)
}

func print(text string) {
	fmt.Print(text)
}
