package camera

import(
	"hot/m/engine/transform"
	"hot/m/engine/vector"
)

type Camera struct {
	T transform.Transform
}

func
New(T transform.Transform) *Camera {
	c := &Camera {T,}
	return c
}

func
(cam *Camera)Rotate(r float64){
}

/* Returns physical(in engine) vector of camera view plus the r value. */
func
(cam *Camera)VecOfView(r float64) vector.Vector {
	vec := vector.New(0, 1)
	r -= cam.T.R
	vec = vec.Rotated(r)
	return vec
}
