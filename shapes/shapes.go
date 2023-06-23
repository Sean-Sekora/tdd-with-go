package shapes

type Shapes struct {
	Shapes   []Shape
	Graphics Graphics
}

func NewShapes(g Graphics) Shapes {
	return Shapes{Graphics: g}
}

func (s Shapes) Add(shape Shape) Shapes {
	ns := s
	ns.Shapes = append(ns.Shapes, shape)
	return ns
}

func (s Shapes) Draw() Shapes {
	for _, shape := range s.Shapes {
		shape.Draw(s.Graphics)
	}
	return s
}
