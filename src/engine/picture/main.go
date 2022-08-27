package picture

import (
	"os"
	"image"
	"github.com/faiface/pixel"
)

func Load(p string) (pixel.Picture, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}
