package shapes

import (
	"fmt"
	"io"
	"os"
)

type ConsoleGraphics struct {
	writer io.Writer
}

func NewConsoleGraphics() ConsoleGraphics {
	return ConsoleGraphics{writer: os.Stdout}
}

func (c ConsoleGraphics) WithWriter(w io.Writer) ConsoleGraphics {
	c.writer = w
	return c
}

func (c ConsoleGraphics) DrawText(text string) {
	fmt.Fprintln(c.writer, text)
}

func (c ConsoleGraphics) DrawHorizontalLine(width int) {
	text := ""
	for i := 0; i < width; i++ {
		text += "X"
	}

	fmt.Fprintln(c.writer, text)
}

func (c ConsoleGraphics) ToString() {
	println(c.writer)
}
