package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Hitbox struct {
	Box image.Rectangle
}

func (hitbox Hitbox) DetectHit(p *image.Point) bool {
	return (hitbox.Box.Min.X <= p.X && p.X <= hitbox.Box.Max.X) && (hitbox.Box.Min.Y <= p.Y && p.Y <= hitbox.Box.Max.Y)
}

func (hitbox Hitbox) Draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	rectangle := ebiten.NewImage(hitbox.Box.Dx(), hitbox.Box.Dy())

	rectangle.Fill(color.Black)
	options.GeoM.Translate(float64(hitbox.Box.Min.X), float64(hitbox.Box.Min.Y))

	screen.DrawImage(rectangle, &options)
}
