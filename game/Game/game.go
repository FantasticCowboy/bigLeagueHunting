package game

import (
	"image/color"
	"log"
	"sort"

	"github.com/FantasticCowboy/bigLeagueHunting/configs"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	// List of all things that can be updated
	idToUpdatable map[int64]*Object
}

func (g *Game) RemoveObject(id int64) {
	delete(g.idToUpdatable, id)
}

func (g *Game) AddObject(obj *Object) {
	if _, ok := g.idToUpdatable[obj.Id]; ok {
		log.Fatalf("Duplicate add for the object")
	}
	g.idToUpdatable[obj.Id] = obj
}

func CreateGame() (g *Game) {
	g = new(Game)
	g.idToUpdatable = make(map[int64]*Object)
	return
}

func (g *Game) Update() error {
	for _, obj := range g.idToUpdatable {
		obj.Update()
	}
	xPos, yPos := ebiten.CursorPosition()
	log.Printf("%v %v", xPos, yPos)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// TODO: sorting on every draw is probably very inefficient
	screen.Fill(color.White)
	objs := make([]*Object, 0)
	for _, obj := range g.idToUpdatable {
		objs = append(objs, obj)
	}
	sort.Slice(objs, func(i, j int) bool {
		return objs[i].RenderLevel < objs[j].RenderLevel
	})
	for _, obj := range objs {
		obj.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return configs.Width, configs.Height
}
