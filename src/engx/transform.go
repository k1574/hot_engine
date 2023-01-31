package engx

type Transform struct {
	P, S Vector
	R float64
}

var (
	T = NewTransform
)

func
NewTransform(P, S Vector, R float64) Transform {
	t := Transform {P, S, R, }
	return t
}
