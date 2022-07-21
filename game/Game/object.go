package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Controller interface {
	Destroy(obj *Object)
	Update(obj *Object)
	Draw(obj *Object, screen *ebiten.Image)
}

type Object struct {

	// Id is a uid and objectType is the type of object
	Id         int64
	ObjectType string

	Img         *ebiten.Image
	DrawOptions *ebiten.DrawImageOptions

	// Kinematic attributes that effect where an image is going to be drawn
	XPos          float64
	YPos          float64
	XSpeed        float64
	YSpeed        float64
	XAcceleration float64
	YAcceleration float64

	// Manipulate the object's attributes and draws the object
	Controller Controller

	// If to draw the image with xPos,yPos in the center
	// or if to draw the iamge with xPos,yPos in the
	// top left corner
	DrawInCenter   bool
	RenderLevel    float64
	XScale, YScale float64
	Roation        float64
	Visible        bool
}

func (obj *Object) UpdatePositioning() {
	obj.XSpeed += obj.XAcceleration
	obj.YSpeed += obj.YAcceleration
	obj.XPos += obj.XSpeed
	obj.YPos += obj.YSpeed
}

func (obj *Object) TranslateImage() {

	if obj.DrawInCenter {
		cos := math.Cos(obj.Roation)
		sin := math.Sin(obj.Roation)

		orgDx := float64(obj.Img.Bounds().Dx()) * obj.XScale / 2
		orgDy := float64(obj.Img.Bounds().Dy()) * obj.YScale / 2

		dx := obj.XPos - (orgDx*cos - orgDx*sin)
		dy := obj.YPos - (orgDy*sin + orgDy*cos)

		obj.DrawOptions.GeoM.Translate(
			dx,
			dy,
		)
	} else {
		obj.DrawOptions.GeoM.Translate(obj.XPos, obj.YPos)
	}
}

func (obj *Object) FormatDrawOptions() {
	obj.DrawOptions.GeoM.Rotate(obj.Roation)

	obj.DrawOptions.GeoM.Scale(obj.XScale, obj.YScale)

	obj.TranslateImage()
}

func (obj *Object) Update() {
	obj.Controller.Update(obj)
}

func (obj *Object) Draw(screen *ebiten.Image) {
	if !obj.Visible {
		return
	}
	obj.DrawOptions = &ebiten.DrawImageOptions{}
	obj.FormatDrawOptions()
	obj.Controller.Draw(obj, screen)

}

func (obj *Object) Destroy() {
	obj.Controller.Destroy(obj)
}
