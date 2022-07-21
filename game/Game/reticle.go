package game

import (
	"github.com/FantasticCowboy/bigLeagueHunting/assets"
	"github.com/FantasticCowboy/bigLeagueHunting/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Reticle struct {
}

func CreateReticle() *Object {
	Reticle := Reticle{}

	return &Object{
		Id:            utils.GenerateUid(),
		ObjectType:    "reticle",
		Img:           assets.Reticle,
		DrawOptions:   &ebiten.DrawImageOptions{},
		XPos:          0,
		YPos:          0,
		XSpeed:        0,
		YSpeed:        0,
		XAcceleration: 0,
		YAcceleration: 0,
		Controller:    &Reticle,
		DrawInCenter:  true,
		RenderLevel:   1,
		XScale:        .5,
		YScale:        .5,
		Roation:       0,
		Visible:       true,
	}

}

func (reticle *Reticle) Destroy(obj *Object) {
	GetGameState().RemoveObject(obj.Id)
}

func (reticle *Reticle) Update(obj *Object) {
	newXpos, newYpos := ebiten.CursorPosition()
	obj.XPos, obj.YPos = float64(newXpos), float64(newYpos)
	obj.UpdatePositioning()
}

func (reticle *Reticle) Draw(obj *Object, screen *ebiten.Image) {
	screen.DrawImage(obj.Img, obj.DrawOptions)
}
