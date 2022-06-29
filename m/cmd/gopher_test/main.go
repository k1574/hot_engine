package main

import(
	"math"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"github.com/k1574/hot/m/engine"
	"github.com/k1574/hot/m/engine/sprite"
	"github.com/k1574/hot/m/engine/object"
	"github.com/k1574/hot/m/engine/vector"
	"github.com/k1574/hot/m/engine/transform"
)

var(
	eng *engine.Engine
	goph_player1 GopherPlayer
	goph_sprite *sprite.Sprite
)
type Gopher struct {
	O *object.Object
}

type GopherPlayer struct {
	O *object.Object
	MoveSpeed float64
}

type CameraPlayer struct {
	RotationSpeed float64
	PositionSpeed float64
}

func (g *Gopher)Start() {
}

func (g *Gopher)Update() {
}

func
(g *GopherPlayer)Start() {
}

func
(g *GopherPlayer)Update() {
		pos := &(g.O.T.P)
		angle := &(g.O.T.R)
		movSpeed := g.MoveSpeed
		win := eng.Win
		cam := eng.Cam
		dt := eng.DT
	if win.Pressed(pixelgl.MouseButtonLeft){
		var real vector.Vector
		click := win.MousePosition()
		if !g.O.Floating {
			real = cam.FromRealToAbsVector(click)
		} else {
		real = click
		}
		direction := real.Sub(*pos)
		*angle = math.Atan(direction.Y/direction.X)
		if direction.X < 0 { *angle += math.Pi }
	}
	if win.Pressed(pixelgl.KeyW){
		pos.Y += movSpeed*dt
	}
	if win.Pressed(pixelgl.KeyS){
		pos.Y -= movSpeed*dt
	}
	if win.Pressed(pixelgl.KeyD){
		pos.X += movSpeed*dt
	}
	if win.Pressed(pixelgl.KeyA){
		pos.X -= movSpeed*dt
	}
}

func
(c *CameraPlayer)Start() {
}

func
(c *CameraPlayer)Update() {
	var(
	)
	t := &(eng.Cam.T)
	cam := eng.Cam
	dt := eng.DT
	win := eng.Win

	if win.Pressed(pixelgl.KeyPageUp){
		t.R += dt * c.RotationSpeed
	}
	if win.Pressed(pixelgl.KeyPageDown){
		t.R -= dt * c.RotationSpeed
	}

	if win.Pressed(pixelgl.MouseButtonLeft){
		var real vector.Vector
		click := win.MousePosition()
		real = cam.FromRealToAbsVector(click)
		goph := Gopher{}
		goph.O = &object.Object{}
		goph.O.S = goph_sprite
		goph.O.T = transform.Transform{P: real, S:vector.Vector{1, 1}, R: 0}
		eng.AddBehaviorer(&goph)
	}
	x, y := goph_player1.GetO().T.P.XY()
	t.P = vector.V(x-512, y-384)
}

func (g *GopherPlayer)GetO() *object.Object {return g.O}
func (g *CameraPlayer)GetO() *object.Object {return nil}
func (g *Gopher)GetO() *object.Object {return g.O}

func
main(){
	var err error
	eng = engine.New(pixelgl.WindowConfig{
		Title: "Gopher Test",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync: true,
	})

	goph_sprite, err = sprite.Load("media/player1.png")
	if err != nil {
		panic(err)
	}

	goph_player := GopherPlayer{
		O: object.New(
			transform.New(
				eng.WinCfg.Bounds.Center(),
				pixel.Vec{1, 0.33},
				0),
			goph_sprite,
			false),
		MoveSpeed: 0.,
	}
	eng.AddBehaviorer(&goph_player)
	
	goph_player1 = GopherPlayer{
		O: &object.Object {
			T: transform.Transform {
				P: pixel.ZV,
				S: pixel.Vec{0.35, 0.35},
			},
			S: goph_sprite,
		},
		MoveSpeed: 100.0,
	}
	eng.AddBehaviorer(&goph_player1)
	
	cam := CameraPlayer {
		RotationSpeed: 4,
		PositionSpeed: 200,
	}
	eng.AddBehaviorer(&cam)

	eng.Run()
}
