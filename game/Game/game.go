package game

import (
	"image/color"
	"log"
	"sort"

	"github.com/FantasticCowboy/bigLeagueHunting/assets"
	"github.com/FantasticCowboy/bigLeagueHunting/configs"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	// List of all things that can be updated
	IdToUpdatable map[int64]*Object

	// Buffer of things to be added to idToUpdatable
	AddBuffer map[int64]*Object

	CurrentState State
}

func (g *Game) RemoveObject(id int64) {
	delete(g.IdToUpdatable, id)
	delete(g.AddBuffer, id)
}

func (g *Game) TransferBuffer() {
	for _, obj := range g.AddBuffer {
		g.IdToUpdatable[obj.Id] = obj
	}
	for k := range g.AddBuffer {
		delete(g.AddBuffer, k)
	}
}

func (g *Game) GetObject(id int64) *Object {
	return g.IdToUpdatable[id]
}

func (g *Game) AddObject(obj *Object) {
	if _, ok := g.IdToUpdatable[obj.Id]; ok {
		log.Fatalf("Duplicate add for the object")
	}
	if _, ok := g.AddBuffer[obj.Id]; ok {
		log.Fatalf("Duplicate add for the object")
	}
	g.AddBuffer[obj.Id] = obj
}

func CreateGame() (g *Game) {
	g = new(Game)
	g.IdToUpdatable = make(map[int64]*Object)
	g.AddBuffer = make(map[int64]*Object)
	return
}

func (g *Game) Update() error {
	g.CurrentState.Update()
	g.TransferBuffer()

	for _, obj := range g.IdToUpdatable {
		obj.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// TODO: sorting on every draw is probably very inefficient
	screen.Fill(color.RGBA{11, 127, 171, 255})

	drawOptions := &ebiten.DrawImageOptions{}
	drawOptions.GeoM.Scale(
		configs.Width/float64(assets.Background.Bounds().Dx()),
		configs.Height/float64(assets.Background.Bounds().Dy()),
	)
	screen.DrawImage(assets.Background, drawOptions)

	objs := make([]*Object, 0)
	for _, obj := range g.IdToUpdatable {
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
