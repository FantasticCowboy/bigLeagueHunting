package gameObjects

import (
	"github.com/FantasticCowboy/bigLeagueHunting/game"
	"github.com/FantasticCowboy/bigLeagueHunting/game/geometry"
	"github.com/FantasticCowboy/bigLeagueHunting/game/hitbox"
	"github.com/FantasticCowboy/bigLeagueHunting/game/interfaces"
	"github.com/FantasticCowboy/bigLeagueHunting/game/sprite"
	"github.com/FantasticCowboy/bigLeagueHunting/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type BasicObject struct {
	id             int64
	sprite         *sprite.SpriteController
	hitbox         *hitbox.Hitbox
	gameState      *game.Game
	pathGenerator  *interfaces.PathGenerator
	collidables    map[int64]func(*BasicObject)
	imagePosition  *geometry.Point
	hitboxPosition *geometry.Point
}

func CreateEmptyBasicObj(gameState *game.Game) *BasicObject {
	obj := BasicObject{}
	obj.id = utils.GenerateUid()
	obj.gameState = gameState
	obj.collidables = make(map[int64]func(*BasicObject))
	return &obj
}

func (obj *BasicObject) Update() {
	obj.CheckAndHandleCollisions()

	if obj.sprite != nil {
		obj.sprite.Update()
	}
}

func (obj *BasicObject) Draw(screen *ebiten.Image) {
	if obj.sprite != nil {
		obj.sprite.Draw(screen)
	}
}

func (obj *BasicObject) Destroy() {
	obj = nil
}
