package shapes

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Draw(g Graphics) {
	for i := 0; i < r.Height; i++ {
		g.DrawHorizontalLine(r.Width)
	}
}
