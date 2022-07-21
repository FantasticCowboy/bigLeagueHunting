package main

import (
	"image"
	"log"
	"time"

	"github.com/FantasticCowboy/bigLeagueHunting/configs"
	game "github.com/FantasticCowboy/bigLeagueHunting/game/Game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.CreateGame()
	game.SetGameState(g)

	ebiten.SetWindowSize(960, 540)

	g.CurrentState = &game.Play{
		Emitters:                make(map[*game.Emitter]bool),
		MaxAmountOfDuckEmitters: 5,
		EmitterLifespan:         time.Minute,
		EmitterSpawnLocations:   image.Rect(0, 0, 16, configs.Height-16),
	}

	g.AddObject(game.CreateReticle())

	if err := ebiten.RunGame(game.GetGameState()); err != nil {
		log.Fatal(err)
	}
}
