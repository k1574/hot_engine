package engx

import(
	"github.com/faiface/pixel"
)

type Object struct {
	T Transform
	P *pixel.Picture
	S *Sprite
	B *pixel.Batch
	Floating bool
}

var (
	O = NewObject
)

func
NewObject(t Transform, p *pixel.Picture, s  *Sprite,
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

func (o *Object)O() *Object {
	return o
}

func (o *Object) Update() {
}

func (o *Object) Start() {
}
