package gameObjects

import (
	"github.com/FantasticCowboy/bigLeagueHunting/game/logger"
)

func (obj *BasicObject) AddCollidable(id int64, handler func(*BasicObject)) {
	obj.collidables[id] = handler
}

func (obj *BasicObject) RemoveCollidable(id int64) {
	if _, ok := obj.collidables[id]; ok {
		delete(obj.collidables, id)
	}
}

func (obj *BasicObject) CheckAndHandleCollisions() {
	for id := range obj.collidables {
		if obj.detectCollision(id) {
			obj.collisionHandler(id)
		}
	}
}

func (obj *BasicObject) detectCollision(id int64) bool {
	otherObj := obj.gameState.GetCollidable(id)
	if otherObj == nil {
		return false
	}
	return obj.hitbox.DetectHit(otherObj.GetHitbox())
}

func (obj *BasicObject) collisionHandler(id int64) {
	handler, ok := obj.collidables[id]
	if !ok {
		logger.Logf("%s", "error")
		return
	}
	handler(obj)
}
