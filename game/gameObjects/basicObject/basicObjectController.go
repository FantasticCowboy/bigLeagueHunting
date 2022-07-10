package gameObjects

import "github.com/FantasticCowboy/bigLeagueHunting/game"

// Controlller to change the state of a basic object
type basicObjectController struct {
	basicObject   *BasicObject
	XPos          float64
	YPos          float64
	XSpeed        float64
	YSpeed        float64
	XAcceleration float64
	YAcceleration float64
}

func CreateBasicObjectController(
	xPos, yPos, xSpeed, ySpeed, xAcceleartion, yAcceleration float64, gameState *game.Game,
) *basicObjectController {

	controller := basicObjectController{}
	controller.XPos = xPos
	controller.YPos = yPos
	controller.XSpeed = xSpeed
	controller.YSpeed = ySpeed
	controller.XAcceleration = xAcceleartion
	controller.YAcceleration = yAcceleration

	controller.basicObject = CreateEmptyBasicObj(gameState)

	return &controller
}

func (controler *basicObjectController) Update() {
	controler.XSpeed += controler.XAcceleration
	controler.YSpeed += controler.YAcceleration
	controler.XPos += controler.XSpeed
	controler.YPos += controler.YSpeed

	controler.basicObject.SetHitboxPosition(controler.XPos, controler.YPos)
	controler.basicObject.SetSpritePosition(controler.XPos, controler.YPos)
	controler.basicObject.Update()
}

func (controller *basicObjectController) SetBasicObject(basicObject *BasicObject, doDelete bool) {
	if doDelete {
		controller.basicObject.Destroy()
	}
	controller.basicObject = basicObject
}

func (controller *basicObjectController) SetPosition(xPos, yPos float64) {
	controller.XPos = xPos
	controller.YPos = yPos
}

func (controller *basicObjectController) SetSpeed(xSpeed, ySpeed float64) {
	controller.XSpeed = xSpeed
	controller.YSpeed = ySpeed
}

func (controller *basicObjectController) SetAcceleration(xAcceleration, yAcceleration float64) {
	controller.XAcceleration = xAcceleration
	controller.YAcceleration = yAcceleration
}
