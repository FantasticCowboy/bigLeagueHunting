package game

import (
	"image"
	"log"

	"github.com/FantasticCowboy/bigLeagueHunting/assets"
	"github.com/FantasticCowboy/bigLeagueHunting/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type DuckController struct {
	Hitbox    Hitbox
	animation *AnimationController
	alpha     float64
}

func (duck *DuckController) Destroy(obj *Object) {
	GetGameState().RemoveObject(obj.Id)
}

func (duck *DuckController) Update(obj *Object) {
	obj.UpdatePositioning()
	obj.Img = duck.animation.GetFrame()
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

func (duck *DuckController) Draw(obj *Object, screen *ebiten.Image) {
	duck.Hitbox.Draw(screen)

	screen.DrawImage(obj.Img, obj.DrawOptions)

}
func CreateDuck(xPos, yPos, xSpeed, ySpeed float64) *Object {
	controller := DuckController{}
	minPoint := image.Point{int(xPos) - 100, int(yPos) - 100}
	maxPoint := image.Point{int(xPos) + 100, int(yPos) + 100}
	rec := image.Rectangle{minPoint, maxPoint}
	controller.Hitbox = Hitbox{
		Box:     rec,
		Visible: false,
	}

	controller.animation = CreateAnimationController(
		16,
		16,
		4,
		assets.DuckImage,
		4,
	)

	return &Object{
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
		XScale:        10,
		YScale:        10,
		Roation:       0,
		Visible:       true,
	}
}
