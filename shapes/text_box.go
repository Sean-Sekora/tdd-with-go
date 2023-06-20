package shapes

type TextBox struct {
	Text string
}

func (t *TextBox) Draw(g Graphics) {
	g.DrawText(t.Text)
}
