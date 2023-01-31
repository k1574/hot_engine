package engx

import(
	//"github.com/faiface/pixel"
)

type Camera struct {
	T Transform
	Around Vector
}

func
NewCamera(T Transform, Around Vector) *Camera {
	c := &Camera{T, Around, }
	return c
}

/*func
(cam *Camera)Rotate(r float64){
}*/

/* Returns physical(in engine) vector of camera view plus the r value. */
func
(cam *Camera)VecOfView(r float64) Vector {
	vec := V(0, 1)
	r -= cam.T.R
	vec = vec.Rotated(r)
	return vec
}

func
(cam *Camera)FromAbsToRealMatrix() Matrix {
	return IM.Rotated(cam.T.P.Add(cam.Around),
			cam.T.R).
		Moved(ZV.Sub(cam.T.P)).
		ScaledXY(cam.T.P.Add(cam.Around), cam.T.S)
}

/* Convert position from real to absolute.
	Real values mean that it how sprites are drawn.
	Absolute are values which the engine works with. */
func
(cam *Camera)FromRealToAbsVector(v Vector) Vector {
	return cam.FromAbsToRealMatrix().Unproject(v)
}
