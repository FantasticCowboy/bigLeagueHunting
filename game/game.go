package game

import (
	"github.com/FantasticCowboy/bigLeagueHunting/configs"
	"github.com/FantasticCowboy/bigLeagueHunting/game/interfaces"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	// List of all things that can be updated
	idToUpdatable map[int64]interfaces.Updatable

	// List of all things that are collidables
	idToCollidable map[int64]interfaces.BasicObject
}

func (g *Game) CheckAndHandleCollisions() {
	for _, obj := range g.idToCollidable {
		obj.CheckAndHandleCollisions()
	}
}

func (g *Game) GetCollidable(id int64) interfaces.BasicObject {
	return g.idToCollidable[id]
}

// Note. Returns nil if there is no object with the ID
func (g *Game) GetUpdatable(id int64) interfaces.Updatable {
	return CreateGame().idToUpdatable[id]
}

func (g *Game) RemoveUpdateable(id int64) {
	delete(g.idToUpdatable, id)
	delete(g.idToCollidable, id)
}

func (g *Game) AddCollidable(obj interfaces.BasicObject) {
	g.AddUpdatable(obj)
	g.idToCollidable[obj.GetId()] = obj
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
