package gameObjects

import (
	"github.com/FantasticCowboy/bigLeagueHunting/game"
	"github.com/FantasticCowboy/bigLeagueHunting/game/geometry"
	"github.com/FantasticCowboy/bigLeagueHunting/game/hitbox"
	"github.com/FantasticCowboy/bigLeagueHunting/game/sprite"
	"github.com/FantasticCowboy/bigLeagueHunting/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type BasicObject struct {
	id                int64
	spriteController  *sprite.SpriteController
	hitbox            *hitbox.Hitbox
	hitboxProportions *geometry.Rectangle
	gameState         *game.Game
	collidables       map[int64]func(*BasicObject)
	spritePosition    *geometry.Point
	hitboxPosition    *geometry.Point
}

func CreateEmptyBasicObj(gameState *game.Game) *BasicObject {
	obj := BasicObject{}
	obj.id = utils.GenerateUid()
	obj.gameState = gameState
	obj.collidables = make(map[int64]func(*BasicObject))
	return &obj
}

func (obj *BasicObject) Update() {
	if obj.hitbox != nil {
		obj.hitbox.UpdatePosition(obj.hitboxPosition)
		obj.CheckAndHandleCollisions()
	}
	if obj.spriteController != nil {
		obj.spriteController.SetPosition(obj.spritePosition.GetCords())
		obj.spriteController.Update()
	}
}

func (obj *BasicObject) Draw(screen *ebiten.Image) {
	if obj.spriteController != nil {
		obj.spriteController.Draw(screen)
	}
}

func (obj *BasicObject) Destroy() {
	obj = nil
}
