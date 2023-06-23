package shapes

import "io"

type Graphics interface {
	DrawText(text string)
	DrawHorizontalLine(width int)
	WithWriter(w io.Writer) Graphics
}
