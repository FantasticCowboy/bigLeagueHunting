package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Hitbox struct {
	Box     image.Rectangle
	Visible bool
}

func (hitbox Hitbox) DetectHit(p *image.Point) bool {
	return (hitbox.Box.Min.X <= p.X && p.X <= hitbox.Box.Max.X) && (hitbox.Box.Min.Y <= p.Y && p.Y <= hitbox.Box.Max.Y)
}

func (hitbox Hitbox) Draw(screen *ebiten.Image) {
	if hitbox.Visible {
		options := ebiten.DrawImageOptions{}
		rectangle := ebiten.NewImage(hitbox.Box.Dx(), hitbox.Box.Dy())
		options.GeoM.Translate(float64(hitbox.Box.Min.X), float64(hitbox.Box.Min.Y))
		rectangle.Fill(color.Black)
		screen.DrawImage(rectangle, &options)
	}
}
