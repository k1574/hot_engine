package object

import(
	"github.com/surdeus/hot/src/engine/transform"
	"github.com/surdeus/hot/src/engine/sprite"
)

type Object struct {
	T transform.Transform
	S *sprite.Sprite
	Floating bool
}

func
New(t transform.Transform, s *sprite.Sprite, f bool) *Object {
	o := &Object {
		T: t,
		S: s,
		Floating: f,
	}
	return o
}
