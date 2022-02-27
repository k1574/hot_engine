package vector

import(
	"github.com/faiface/pixel"
)

type Vector = pixel.Vec

var(
	Z = pixel.ZV
	V = New
)

/*func
(v1 Vector)Add(v2 Vector) Vector {
	return Vector(pixel.Vec(v1).Add(pixel.Vec(v2)))
}*/

func
New(X, Y float64) Vector {
	v := Vector {X, Y,}
	return v
}

