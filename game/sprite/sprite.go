package sprite

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	drawOptions ebiten.DrawImageOptions
	img         *ebiten.Image
	uid         int64
	visible     bool
}

// Description: draws the image in the center of the image
func Create(img *ebiten.Image, uid int64, xPos float64, yPos float64, xScale float64, yScale float64) (obj *Sprite) {
	drawOption := ebiten.DrawImageOptions{}
	drawOption.GeoM.Scale(xScale, yScale)
	drawOption.GeoM.Translate(xPos, yPos)
	return &Sprite{drawOption, img, uid, true}
}

func CreateInCenter(img *ebiten.Image, uid int64, xPos float64, yPos float64, xScale float64, yScale float64) (obj *Sprite) {
	return Create(
		img,
		uid,
		xPos-xScale*float64(img.Bounds().Dx()/2),
		yPos-yScale*float64(img.Bounds().Dy()/2),
		xScale,
		yScale,
	)
}

func (pic *Sprite) GetCenter() (int, int) {
	return int(pic.drawOptions.GeoM.Element(0, 2)) + pic.img.Bounds().Dx()/2,
		int(pic.drawOptions.GeoM.Element(1, 2)) + pic.img.Bounds().Dy()/2
}

func (pic *Sprite) GetTopLeftOfImage() (int, int) {
	return int(pic.drawOptions.GeoM.Element(0, 2)),
		int(pic.drawOptions.GeoM.Element(1, 2))
}

func (pic *Sprite) Update() {

}

func (pic *Sprite) SetVisibility(predicate bool) {
	pic.visible = predicate
}

func (pic *Sprite) Destroy() {
	pic.visible = false
}

func (pic *Sprite) Draw(screen *ebiten.Image) {
	if pic.visible {
		screen.DrawImage(pic.img, &pic.drawOptions)
	}
}

func (pic *Sprite) GetId() int64 {
	return pic.uid
}
