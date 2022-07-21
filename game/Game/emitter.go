package game

import (
	"time"

	"github.com/FantasticCowboy/bigLeagueHunting/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Emitter struct {
	basePeriod    time.Duration
	randomness    time.Duration // range of randomnes
	currentPeriod time.Duration
	lastEmit      time.Time
	BirthTime     time.Time
	objectCreator func() *Object
	ObjectPtr     Object
}

func (duckEmitter *Emitter) generateNewCurrentPeriod() {
	duckEmitter.currentPeriod = duckEmitter.basePeriod + time.Duration(utils.GenerateRandomNumber(float64(duckEmitter.randomness)))
}

func (emitter *Emitter) Destroy(obj *Object) {
	GetGameState().RemoveObject(obj.Id)
}

func (duckEmmiter *Emitter) Update(obj *Object) {
	if time.Since(duckEmmiter.lastEmit) > duckEmmiter.currentPeriod {
		GetGameState().AddObject(duckEmmiter.objectCreator())
		duckEmmiter.lastEmit = time.Now()
		duckEmmiter.generateNewCurrentPeriod()
	}
}

func (duckEmmiter *Emitter) Draw(obj *Object, screen *ebiten.Image) {

}

func CreateEmitter(
	period time.Duration,
	randomness time.Duration,
	objectCreator func() *Object,
) (*Object, *Emitter) {

	controller := &Emitter{
		lastEmit:      time.Now(),
		basePeriod:    period,
		randomness:    randomness,
		BirthTime:     time.Now(),
		objectCreator: objectCreator,
	}

	controller.generateNewCurrentPeriod()

	obj := Object{
		Id:            utils.GenerateUid(),
		ObjectType:    "DuckEmitter",
		Img:           nil,
		DrawOptions:   nil,
		XPos:          0,
		YPos:          0,
		XSpeed:        0,
		YSpeed:        0,
		XAcceleration: 0,
		YAcceleration: 0,
		Controller:    controller,
		DrawInCenter:  true,
		RenderLevel:   1,
		XScale:        0,
		YScale:        0,
		Roation:       0,
		Visible:       false,
	}

	controller.ObjectPtr = obj

	return &obj, controller
}
