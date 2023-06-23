package shapes

type Shapes struct {
	Shapes   []Shape
	Graphics Graphics
}

func NewShapes(g Graphics) Shapes {
	return Shapes{Graphics: g}
}

func (s Shapes) Add(shape Shape) Shapes {
	s.Shapes = append(s.Shapes, shape)
	return s
}

func (s Shapes) Draw() Shapes {
	for _, shape := range s.Shapes {
		shape.Draw(s.Graphics)
	}
	return s
}
