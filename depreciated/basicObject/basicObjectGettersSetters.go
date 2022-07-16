package gameObjects

import (
	"github.com/FantasticCowboy/bigLeagueHunting/game/geometry"
	"github.com/FantasticCowboy/bigLeagueHunting/game/sprite"
)

func (obj *BasicObject) SetSpritePosition(xPos, yPos float64) {
	obj.spritePosition = geometry.CreatePoint(xPos, yPos)
}
func (obj *BasicObject) SetHitboxPosition(xPos, yPos float64) {
	obj.hitboxPosition = geometry.CreatePoint(xPos, yPos)
}

func (obj *BasicObject) SetSprite(newSprite *sprite.SpriteController) {
	obj.spriteController = newSprite
}

func (obj *BasicObject) GetId() int64 {
	return obj.id
}
