package vector

import(
	"github.com/faiface/pixel"
)

type Vectorer2 interface {
	XY() (x, y float64)
}
type Vector = pixel.Vec

var(
	Z = pixel.ZV
	V = New
)

func Mul(v Vectorer2, m float64) Vector {
	x, y := v.XY()
	return Vector{x*m, y*m}
}

func
New(X, Y float64) Vector {
	v := Vector {X, Y,}
	return v
}

