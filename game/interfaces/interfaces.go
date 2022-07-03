package interfaces

import (
	"github.com/FantasticCowboy/bigLeagueHunting/game/hitbox"
	"github.com/hajimehoshi/ebiten/v2"
)

type Updatable interface {
	Update()
	Draw(screen *ebiten.Image)
	Destroy()
	GetId() int64
}

type BasicObject interface {
	Update()
	Draw(screen *ebiten.Image)
	Destroy()
	GetId() int64
	AddCollidable(id int64)
	RemoveCollidable(id int64)
	GetHitbox() *hitbox.Hitbox
}

type PathGenerator interface {
	GetNext() (int64, int64)
}
