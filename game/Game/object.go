package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ObjectController interface {
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
	Controller ObjectController

	// If to draw the image with xPos,yPos in the center
	// or if to draw the iamge with xPos,yPos in the
	// top left corner
	DrawInCenter bool
}

func (obj *Object) UpdatePositioning() {
	obj.XSpeed += obj.XAcceleration
	obj.YSpeed += obj.YAcceleration
	obj.XPos += obj.XSpeed
	obj.YPos += obj.YSpeed
}

func (obj *Object) Update() {
	obj.Controller.Update(obj)
}

func (obj *Object) Draw(screen *ebiten.Image) {
	obj.DrawOptions = &ebiten.DrawImageOptions{}
	if obj.DrawInCenter {
		obj.DrawOptions.GeoM.Translate(obj.XPos-float64(obj.Img.Bounds().Dx())/2, obj.YPos-float64(obj.Img.Bounds().Dy())/2)
	} else {
		obj.DrawOptions.GeoM.Translate(obj.XPos, obj.YPos)
	}
	obj.Controller.Draw(obj, screen)
	screen.DrawImage(obj.Img, obj.DrawOptions)
}

func (obj *Object) Destroy() {
	obj.Controller.Destroy(obj)
}
