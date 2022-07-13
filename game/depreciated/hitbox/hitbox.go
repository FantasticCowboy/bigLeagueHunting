package hitbox

import "github.com/FantasticCowboy/bigLeagueHunting/game/geometry"

type Hitbox struct {
	box      *geometry.Rectangle
	position *geometry.Point
}

func (hitbox *Hitbox) UpdatePosition(newPos *geometry.Point) {
	xStart, yStart := hitbox.box.GetTopLeftCorner().GetCords()
	xEnd, yEnd := hitbox.box.GetBottomRight().GetCords()
	xOffset := hitbox.position.GetX() - newPos.GetX()
	yOffset := hitbox.position.GetY() - newPos.GetY()

	hitbox.box = geometry.CreateRectangeCorners(
		xStart+xOffset, xEnd+xOffset,
		yStart+yOffset, yEnd+yOffset,
	)
	hitbox.position = newPos
}

func (hitbox *Hitbox) GetRectangle() *geometry.Rectangle {
	return hitbox.box
}

func CreateHitbox(box *geometry.Rectangle, position *geometry.Point) *Hitbox {
	h := Hitbox{}
	h.position = position
	xWidth := box.GetTopLeftCorner().GetX() - box.GetTopRight().GetX()
	yLength := box.GetTopLeftCorner().GetY() - box.GetBottomLeft().GetY()
	xStart, yStart := box.GetTopLeftCorner().GetCords()
	xEnd, yEnd := box.GetBottomRight().GetCords()

	h.box = geometry.CreateRectangeCorners(
		xStart-(xWidth/2), yStart-(yLength/2),
		xEnd+(xWidth/2), yEnd+(yLength/2),
	)
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
