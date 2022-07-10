package gameObjects

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
