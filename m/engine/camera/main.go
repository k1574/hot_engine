package camera

import(
	"hot/m/engine/transform"
	"hot/m/engine/vector"
	"hot/m/engine/matrix"
)

type Camera struct {
	T transform.Transform
	Around vector.Vector
}

func
New(T transform.Transform, Around vector.Vector) *Camera {
	c := &Camera{T, Around, }
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

func
(cam *Camera)FromAbsToRealMatrix() matrix.Matrix {
	return matrix.IM.Rotated(cam.T.P.Add(cam.Around),
			cam.T.R).
		Moved(vector.Z.Sub(cam.T.P)).
		ScaledXY(cam.T.P.Add(cam.Around),
		cam.T.S)
}