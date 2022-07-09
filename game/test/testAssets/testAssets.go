package testAssets

import (
	"github.com/FantasticCowboy/bigLeagueHunting/assets"
	"github.com/FantasticCowboy/bigLeagueHunting/game"
	"github.com/FantasticCowboy/bigLeagueHunting/game/gameObjects"
)

var DuckImage = *assets.ReadImage("/Users/lukehobeika/Desktop/Projects/bigLeagueHunting/assets/duck.png")

func CreateTestDuck(gameState *game.Game, xPos float64, yPos float64) (obj *gameObjects.BasicObject) {
	obj = gameObjects.CreateEmptyBasicObj(gameState)
}
