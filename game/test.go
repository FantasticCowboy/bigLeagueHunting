package main

import (
	"log"

	game "github.com/FantasticCowboy/bigLeagueHunting/game/Game"
	objects "github.com/FantasticCowboy/bigLeagueHunting/game/Objects"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.CreateGame()
	game.SetGameState(g)
	ebiten.SetWindowSize(960, 540)
	duck := objects.CreateDuck(0, 0, 1, 1)
	reticle := objects.CreateReticle()

	//ebiten.SetCursorMode(ebiten.CursorModeHidden)

	g.AddObject(duck)
	g.AddObject(reticle)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
