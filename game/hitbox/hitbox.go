package hitbox

import "github.com/FantasticCowboy/bigLeagueHunting/game/geometry"

type Hitbox struct {
	box geometry.Rectangle
}

func CreateHitBox(box geometry.Rectangle) *Hitbox {
	h := Hitbox{}
	h.box = box
	return &h
}

func CreateHitBoxCorners(xStart float64,
	xEnd float64,
	yStart float64,
	yEnd float64) *Hitbox {
	h := Hitbox{}
	h.box = *geometry.CreateRectangeCorners(xStart, yStart, xEnd, yEnd)
	return &h
}

func (hitbox *Hitbox) DetectHitPoint(p *geometry.Point) bool {
	xStart, yStart := hitbox.box.GetTopLeftCorner().GetCords()
	xEnd, yEnd := hitbox.box.GetBottomRight().GetCords()
	x, y := p.GetCords()
	return (x >= xStart && x <= xEnd) && (y >= yStart && y <= yEnd)
}

func (hitbox *Hitbox) DetectHit(other *Hitbox) bool {
	return (hitbox.DetectHitPoint(other.box.GetBottomLeft()) ||
		hitbox.DetectHitPoint(other.box.GetBottomRight()) ||
		hitbox.DetectHitPoint(other.box.GetTopLeftCorner()) ||
		hitbox.DetectHitPoint(other.box.GetTopRight()))
}
