package shapes

type Shapes struct {
	All      []Shape
	Graphics Graphics
}

func NewShapes(g Graphics) Shapes {
	return Shapes{Graphics: g}
}

func (s Shapes) Add(shape Shape) Shapes {
	ns := s
	ns.All = append(s.All, shape)
	return ns
}

func (s Shapes) Draw() Shapes {
	for _, shape := range s.All {
		shape.Draw(s.Graphics)
	}
	return s
}
