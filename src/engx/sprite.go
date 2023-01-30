package sprite

import (
	"github.com/faiface/pixel"
	"os"
	"image"
	_ "image/png"
)

type Sprite = pixel.Sprite

func 
LoadSprite(path string) (*Sprite, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	pic := pixel.PictureDataFromImage(img)	
	sprite := pixel.NewSprite(pic, pic.Bounds())

	return sprite, nil
}
