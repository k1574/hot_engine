package vector

import(
	"github.com/faiface/pixel"
)

type Vector = pixel.Vec

var(
	Z = pixel.ZV
)

func
New(X, Y float64) Vector {
	v := Vector {X, Y,}
	return v
}