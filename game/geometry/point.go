package geometry

type Point struct {
	x float64
	y float64
}

func CreatePoint(x float64, y float64) *Point {
	pTmp := CreatePointObject(x, y)
	return &pTmp
}

func CreatePointObject(x float64, y float64) (p Point) {
	p.x = x
	p.y = y
	return
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func (p *Point) GetCords() (x float64, y float64) {
	return p.x, p.y
}
