package object

import(
	"hot/m/engine/transform"
	"hot/m/engine/sprite"
)

type Object struct {
	T transform.Transform
	S *sprite.Sprite
}

func
New(t transform.Transform, s *sprite.Sprite) *Object {
	o := &Object {
		T: t,
		S: s,
	}
	return o
}