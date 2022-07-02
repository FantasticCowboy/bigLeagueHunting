package game

import (
	"github.com/FantasticCowboy/bigLeagueHunting/configs"
	"github.com/FantasticCowboy/bigLeagueHunting/game/interfaces"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	idToObjects map[int64]interfaces.Updatable
}

func CreateGame() (g *Game) {
	g = new(Game)
	g.idToObjects = make(map[int64]interfaces.Updatable)
	return
}

func (g *Game) AddObject(obj interfaces.Updatable) {
	g.idToObjects[obj.GetId()] = obj
}

func (g *Game) Update() error {
	for _, obj := range g.idToObjects {
		obj.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, obj := range g.idToObjects {
		obj.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return configs.Width, configs.Height
}
