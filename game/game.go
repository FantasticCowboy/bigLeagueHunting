package game

import (
	"github.com/FantasticCowboy/bigLeagueHunting/configs"
	"github.com/FantasticCowboy/bigLeagueHunting/game/interfaces"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	idToUpdatable map[int64]interfaces.Updatable
}

// Note. Returns nil if there is no object with the ID
func (g *Game) GetUpdatable(id int64) interfaces.Updatable {
	return CreateGame().idToUpdatable[id]
}

func (g *Game) RemoveUpdateable(id int64) {
	delete(g.idToUpdatable, id)
}

func (g *Game) AddUpdatable(obj interfaces.Updatable) {
	g.idToUpdatable[obj.GetId()] = obj
}

func CreateGame() (g *Game) {
	g = new(Game)
	g.idToUpdatable = make(map[int64]interfaces.Updatable)
	return
}

func (g *Game) Update() error {
	for _, obj := range g.idToUpdatable {
		obj.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, obj := range g.idToUpdatable {
		obj.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return configs.Width, configs.Height
}
