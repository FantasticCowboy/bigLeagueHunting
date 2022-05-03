package assets

import (
	"image"
	"os"
)

// TODO: embed images in native go code

func ReadImage(file string) image.Image {
	f, err := os.Open(file)

}
