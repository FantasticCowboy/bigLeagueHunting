package main

import (
	"log"
	"time"

	"github.com/FantasticCowboy/bigLeagueHunting/assets"
	"github.com/FantasticCowboy/bigLeagueHunting/configs"
	"github.com/FantasticCowboy/bigLeagueHunting/game"
	gameObjects "github.com/FantasticCowboy/bigLeagueHunting/game/gameObjects/basicObject"
	"github.com/FantasticCowboy/bigLeagueHunting/game/geometry"
	"github.com/FantasticCowboy/bigLeagueHunting/game/hitbox"
	sprite "github.com/FantasticCowboy/bigLeagueHunting/game/sprite"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	img := *assets.ReadImage("/Users/lukehobeika/Desktop/Projects/bigLeagueHunting/assets/duck.png")

	spriteControler := sprite.CreateSpriteController(&img, configs.Width/2, configs.Height/2, 0, 0)

	g := game.CreateGame()
	g.AddUpdatable(spriteControler)
	ebiten.SetWindowSize(960, 540)

	hitbox1 := hitbox.CreateHitBoxCorners(0, 10, 0, 10)
	hitbox2 := hitbox.CreateHitBoxCorners(0, 10, 0, 10)
	hitbox3 := hitbox.CreateHitBoxCorners(100, 1000, 100, 1000, geometry.CreatePoint(0, 0))

	tmp := gameObjects.BasicObject{}

	log.Printf("%v", tmp)

	log.Printf("%v", hitbox1.DetectHit(hitbox2))
	log.Printf("%v", hitbox1.DetectHit(hitbox3))

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
