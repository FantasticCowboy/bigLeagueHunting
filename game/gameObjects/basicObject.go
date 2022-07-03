package gameObjects

import (
	"github.com/FantasticCowboy/bigLeagueHunting/game"
	"github.com/FantasticCowboy/bigLeagueHunting/game/hitbox"
	"github.com/FantasticCowboy/bigLeagueHunting/game/interfaces"
	"github.com/FantasticCowboy/bigLeagueHunting/game/logger"
	"github.com/FantasticCowboy/bigLeagueHunting/game/sprite"
	"github.com/hajimehoshi/ebiten/v2"
)

type BasicObject struct {
	id            int64
	sprite        *sprite.SpriteController
	hitbox        *hitbox.Hitbox
	gameState     *game.Game
	pathGenerator *interfaces.PathGenerator

	collidables        []int64
	collissionHandlers map[int64]func(*BasicObject)
}

func (obj *BasicObject) Update() {

}

func (obj *BasicObject) Draw(screen *ebiten.Image) {

}

func (obj *BasicObject) Destroy() {

}

func (obj *BasicObject) GetId() int64 {
	return obj.id
}

func (obj *BasicObject) CheckAndHandleCollisions() {
	for _, id := range obj.collidables {
		if obj.detectCollision(id) {
			obj.collisionHandler(id)
		}
	}
}

func (obj *BasicObject) detectCollision(id int64) bool {
	updatable := obj.gameState.GetUpdatable(id)
	if updatable == nil {
		return false
	}
	// Todo. This seems sloppy, doing it because it seems logical to have
	// only one mapping
	otherObj, ok := (updatable).(interfaces.BasicObject)
	if !ok {
		return false
	}
	return obj.hitbox.DetectHit(otherObj.GetHitbox())
}

func (obj *BasicObject) collisionHandler(id int64) {
	handler, ok := obj.collissionHandlers[id]
	if !ok {
		logger.Logf("%s", "error")
		return
	}
	handler(obj)
}
