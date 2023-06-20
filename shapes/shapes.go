package shapes

type Shapes struct {
	All      []Shape
	Graphics Graphics
}

func (s *Shapes) Add(shape Shape) {
	(*s).All = append((*s).All, shape)
}

func (s *Shapes) Draw() {
	for _, shape := range s.All {
		shape.Draw(s.Graphics)
	}
}
