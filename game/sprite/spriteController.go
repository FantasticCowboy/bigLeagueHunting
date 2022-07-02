package sprite

import (
	"github.com/FantasticCowboy/bigLeagueHunting/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

// TODO: need to make the sprite controller only update the position
// and other aspects of the sprite
// Basically its responsiblity should be drawing the sprite given some set
// of options as well as cleaning up the sprite
// different than the sprite because the sprite cannot redraw itself
// Maybe the classes should be merged?

type SpriteController struct {
	image   *ebiten.Image
	sprite  *Sprite
	xPos    float64
	yPos    float64
	xSpeed  float64
	ySpeed  float64
	uid     int64
	visible bool
}

func CreateSpriteController(img *ebiten.Image, xPos float64, yPos float64, xSpeed float64, ySpeed float64) *SpriteController {
	sprite := CreateInCenter(img, utils.GenerateUid(), xPos, yPos, 1, 1)
	return &SpriteController{img, sprite, xPos, yPos, xSpeed, ySpeed, utils.GenerateUid(), true}
}

func (controller *SpriteController) SetImage(image *ebiten.Image) {
	controller.image = image
}

func (controller *SpriteController) SetPosition(xPos float64, yPos float64) {
	controller.xPos = xPos
	controller.yPos = yPos
}

func (controller *SpriteController) SetSpeed(xSpeed float64, ySpeed float64) {
	controller.xSpeed = xSpeed
	controller.ySpeed = ySpeed
}

func (controller *SpriteController) UpdateSpeed(xSpeed float64, ySpeed float64) {
	controller.xSpeed = xSpeed
	controller.ySpeed = ySpeed
}

func (controller *SpriteController) SetVisibility(predicate bool) {
	controller.visible = predicate
}

func (controller *SpriteController) Update() {
	controller.xPos += controller.xSpeed
	controller.yPos += controller.ySpeed

	// Should get rid of previously allocated sprite
	controller.sprite = CreateInCenter(
		controller.image,
		utils.GenerateUid(),
		controller.xPos,
		controller.yPos,
		1,
		1,
	)
	if !controller.visible {
		controller.sprite.SetVisibility(false)
	}
}

func (controller *SpriteController) Draw(screen *ebiten.Image) {
	controller.sprite.Draw(screen)
}

func (controller *SpriteController) Destroy() {
	controller.sprite.Destroy()
	controller.sprite = nil
}

func (controller *SpriteController) GetId() int64 {
	return controller.uid
}
