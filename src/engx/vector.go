package enginex

import(
	"github.com/faiface/pixel"
)

type Vector = pixel.Vec

var(
	ZV = pixel.ZV
	V = NewVector
)


func
NewVector(X, Y float64) Vector {
	v := Vector {X, Y,}
	return v
}

