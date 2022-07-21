package game

import (
	"image"
	"math/rand"
	"time"
)

type Play struct {
	Emitters                map[*Emitter]bool
	MaxAmountOfDuckEmitters int
	EmitterLifespan         time.Duration
	EmitterSpawnLocations   image.Rectangle
}

func Initialize() {
	GetGameState().AddObject(CreateReticle())
}

func (play *Play) spawnNewEmitter() {

	obj, emitter := CreateEmitter(time.Second, 2*time.Second,
		func() *Object {
			xPos := float64(play.EmitterSpawnLocations.Dx())*rand.Float64() + float64(play.EmitterSpawnLocations.Min.X)
			yPos := float64(play.EmitterSpawnLocations.Dy())*rand.Float64() + float64(play.EmitterSpawnLocations.Min.Y)
			return CreateDuck(xPos, yPos, 10, 0)
		})

	play.Emitters[emitter] = true
	GetGameState().AddObject(obj)

}

func (play *Play) removeOldEmitters() {
	for emitter := range play.Emitters {
		if play.checkIfExpired(emitter) {
			play.deleteEmitter(emitter)
		}
	}
}

func (play *Play) checkIfExpired(emitter *Emitter) bool {
	return time.Since(emitter.BirthTime) > play.EmitterLifespan
}

func (play *Play) deleteEmitter(emitter *Emitter) {
	GetGameState().RemoveObject(emitter.ObjectPtr.Id)
	delete(play.Emitters, emitter)
}

func (play *Play) Destroy() {
	for emitter := range play.Emitters {
		GetGameState().RemoveObject(emitter.ObjectPtr.Id)
		play.deleteEmitter(emitter)
	}
}

func (play *Play) Update() {

	play.removeOldEmitters()
	for len(play.Emitters) < play.MaxAmountOfDuckEmitters {
		play.spawnNewEmitter()
	}
}
