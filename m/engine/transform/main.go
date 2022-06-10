package transform

import(
	"github.com/k1574/hot/m/engine/vector"
)

type Transform struct {
	P, S vector.Vector
	R float64
}

func
New(P, S vector.Vector, R float64) Transform {
	t := Transform {P, S, R, }
	return t
}