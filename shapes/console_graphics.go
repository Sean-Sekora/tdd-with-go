package shapes

import (
	"fmt"
	"go.uber.org/zap"
	"io"
	"os"
	"tdd_with_go/zapadapter"
)

type ConsoleGraphics struct {
	writer io.Writer
}

var logger = zapadapter.NewZapAdapter()

func NewConsoleGraphics() ConsoleGraphics {
	return ConsoleGraphics{writer: os.Stdout}
}

func (c ConsoleGraphics) WithWriter(w io.Writer) ConsoleGraphics {
	c.writer = w
	return c
}

func (c ConsoleGraphics) DrawText(text string) {
	_, err := fmt.Fprintln(c.writer, text)
	if err != nil {
		return
	}
	logger.Info("Successfully drew text")
}

func (c ConsoleGraphics) DrawHorizontalLine(width int) {
	text := ""
	for i := 0; i < width; i++ {
		text += "X"
	}
	_, err := fmt.Fprintln(c.writer, text)
	if err != nil {
		logger.Error("Error while drawing horizontal line", zap.Error(err))
	} else {
		logger.Info("Successfully drew horizontal line")
	}
}

func (c ConsoleGraphics) ToString() {
	fmt.Print(c.writer)
}
