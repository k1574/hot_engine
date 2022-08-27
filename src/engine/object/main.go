package object

import(
	"github.com/surdeus/hot/src/engine/transform"
	"github.com/surdeus/hot/src/engine/sprite"
	"github.com/faiface/pixel"
)

type Object struct {
	T transform.Transform
	P *pixel.Picture
	S *pixel.Sprite
	B *pixel.Batch
	Floating bool
}

func
New(t transform.Transform, p *pixel.Picture, s  *sprite.Sprite,
		b *pixel.Batch, f bool) *Object {
	o := &Object {
		T: t,
		P: p,
		S: s,
		B: b,
		Floating: f,
	}
	return o
}
