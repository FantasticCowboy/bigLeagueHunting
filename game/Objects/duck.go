package objects

import (
	"image"
	"log"

	"github.com/FantasticCowboy/bigLeagueHunting/assets"
	game "github.com/FantasticCowboy/bigLeagueHunting/game/Game"
	"github.com/FantasticCowboy/bigLeagueHunting/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type DuckController struct {
	Hitbox game.Hitbox
	alpha  float64
}

func (duck *DuckController) Destroy(obj *game.Object) {
	game.GetGameState().RemoveObject(obj.Id)
}

func (duck *DuckController) Update(obj *game.Object) {
	obj.UpdatePositioning()

	duck.Hitbox.Box = image.Rectangle{image.Pt(int(obj.XPos)-100, int(obj.YPos)-100), image.Pt(100+int(obj.XPos), 100+int(obj.YPos))}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		p := image.Point{x, y}
		if duck.Hitbox.DetectHit(&p) {
			log.Printf("Hit! %v %v", x, y)
			obj.Destroy()
		}
	}
}

func (duck *DuckController) Draw(obj *game.Object, screen *ebiten.Image) {
	duck.Hitbox.Draw(screen)

	obj.DrawOptions.ColorM.Scale(1, 1, 1, duck.alpha)
	duck.alpha -= .01
	screen.DrawImage(obj.Img, obj.DrawOptions)

}
func CreateDuck(xPos, yPos, xSpeed, ySpeed float64) *game.Object {
	controller := DuckController{}

	minPoint := image.Point{int(xPos) - 100, int(yPos) - 100}
	maxPoint := image.Point{int(xPos) + 100, int(yPos) + 100}
	rec := image.Rectangle{minPoint, maxPoint}
	controller.Hitbox = game.Hitbox{
		Box: rec,
	}
	controller.alpha = 1

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
		RenderLevel:   0,
		XScale:        1,
		YScale:        1,
		Roation:       0,
	}
}
