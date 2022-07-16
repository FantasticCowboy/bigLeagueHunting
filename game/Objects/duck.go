package objects

import (
	"github.com/FantasticCowboy/bigLeagueHunting/assets"
	game "github.com/FantasticCowboy/bigLeagueHunting/game/Game"
	"github.com/FantasticCowboy/bigLeagueHunting/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type DuckController struct {
}

func (duck *DuckController) Destroy(obj *game.Object) {
	game.GetGameState().RemoveObject(obj.Id)
}

func (duck *DuckController) Update(obj *game.Object) {
	obj.UpdatePositioning()
}

func (duck *DuckController) Draw(obj *game.Object, screen *ebiten.Image) {
	screen.DrawImage(obj.Img, obj.DrawOptions)
}
func CreateDuck(xPos, yPos, xSpeed, ySpeed float64) *game.Object {
	controller := DuckController{}
	return &game.Object{
		Id:            utils.GenerateUid(),
		ObjectType:    "duck",
		Img:           assets.DuckImage,
		DrawOptions:   &ebiten.DrawImageOptions{},
		XPos:          xPos,
		YPos:          yPos,
		XSpeed:        xSpeed,
		YSpeed:        ySpeed,
		XAcceleration: 0,
		YAcceleration: 0,
		Controller:    &controller,
		DrawInCenter:  true,
	}
}
