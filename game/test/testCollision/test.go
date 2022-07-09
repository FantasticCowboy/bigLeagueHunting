package main

import (
	"log"
	"time"

	"github.com/FantasticCowboy/bigLeagueHunting/configs"
	"github.com/FantasticCowboy/bigLeagueHunting/game"
	"github.com/FantasticCowboy/bigLeagueHunting/game/test/testAssets"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	img := testAssets.DuckImage

	duck1 := gameObjects

	g := game.CreateGame()
	g.AddUpdatable(spriteControler)
	ebiten.SetWindowSize(960, 540)

	go func() {
		xPos := configs.Width / 2.0
		yPos := configs.Height / 2.0
		visible := false
		for {
			time.Sleep(time.Millisecond * 500)
			spriteControler.SetPosition(xPos, yPos)
			visible = !visible
		}
	}()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
