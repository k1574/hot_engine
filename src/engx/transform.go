package enginex

import(
	"github.com/surdeus/hot/src/engine/vector"
)

type Transform struct {
	P, S vector.Vector
	R float64
}

func
NewT(P, S vector.Vector, R float64) Transform {
	t := Transform {P, S, R, }
	return t
}
