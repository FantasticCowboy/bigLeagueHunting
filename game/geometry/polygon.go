package geometry

type Polygon struct {
	points []*Point
}

func CreatePolygon(points []*Point) *Polygon {
	p := Polygon{}
	p.points = points
	return &p
}
