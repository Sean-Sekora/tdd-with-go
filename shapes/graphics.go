package shapes

type Graphics interface {
	drawText(text string)
	drawHorizontalLine(width int)
}
