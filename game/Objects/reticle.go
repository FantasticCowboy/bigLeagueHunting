package objects

import (
	"github.com/FantasticCowboy/bigLeagueHunting/assets"
	game "github.com/FantasticCowboy/bigLeagueHunting/game/Game"
	"github.com/FantasticCowboy/bigLeagueHunting/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Reticle struct {
}

func CreateReticle() *game.Object {
	Reticle := Reticle{}

	return &game.Object{
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
	}

}

func (reticle *Reticle) Destroy(obj *game.Object) {
	game.GetGameState().RemoveObject(obj.Id)
}

func (reticle *Reticle) Update(obj *game.Object) {
	newXpos, newYpos := ebiten.CursorPosition()
	obj.XPos, obj.YPos = float64(newXpos), float64(newYpos)
	obj.UpdatePositioning()
}

func (reticle *Reticle) Draw(obj *game.Object, screen *ebiten.Image) {

}
