package shapes

type Graphics interface {
	DrawText(text string)
	DrawHorizontalLine(width int)
}
