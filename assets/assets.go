package assets

import (
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// TODO: embed images in native go code
func ReadImage(file string) *ebiten.Image {

	img, _, err := ebitenutil.NewImageFromFile(file)

	if err != nil {
		log.Fatal("Could not read file " + file)
	}

	return img
}

var DuckImage *ebiten.Image = ReadImage("/Users/lukehobeika/Desktop/Projects/bigLeagueHunting/assets/reticle.png")
var Reticle *ebiten.Image = ReadImage("/Users/lukehobeika/Desktop/Projects/bigLeagueHunting/assets/reticle.png")
