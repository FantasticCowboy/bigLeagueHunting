package assets

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/FantasticCowboy/bigLeagueHunting/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

// TODO: embed images in native go code
func ReadImage(file string) *ebiten.Image {
	f, err := os.Open(file)
	utils.FailIfError(err)

	img, _, err := image.Decode(f)

	utils.FailIfError(err)

	return ebiten.NewImageFromImage(img)
}

var DuckImage *ebiten.Image = ReadImage("/Users/lukehobeika/Desktop/Projects/bigLeagueHunting/assets/duck.png")
