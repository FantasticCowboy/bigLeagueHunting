package geometry

import "log"

type Rectangle struct {
	topLeft     *Point
	topRight    *Point
	bottomRight *Point
	bottomLeft  *Point
}

func CreateQuadralateral(
	topLeft *Point,
	topRight *Point,
	bottomRight *Point,
	bottomLeft *Point) (s *Rectangle) {
	s.topLeft = topLeft
	s.bottomLeft = bottomLeft
	s.topRight = topRight
	s.bottomRight = bottomRight
	return
}

func CreateRectangeCorners(
	xStart float64,
	yStart float64,
	xEnd float64,
	yEnd float64,
) *Rectangle {

	if xStart >= xEnd || xStart >= yEnd {
		log.Panicf("Invalid starts and ends")
	}

	s := Rectangle{}
	s.topLeft = CreatePoint(xStart, yStart)
	s.bottomLeft = CreatePoint(xStart, yEnd)
	s.topRight = CreatePoint(xEnd, yStart)
	s.bottomRight = CreatePoint(xEnd, yEnd)
	return &s
}

func (s *Rectangle) GetTopLeftCorner() *Point {
	return s.topLeft
}

func (s *Rectangle) GetTopRight() *Point {
	return s.topRight
}

func (s *Rectangle) GetBottomRight() *Point {
	return s.bottomRight
}

func (s *Rectangle) GetBottomLeft() *Point {
	return s.bottomLeft
}
