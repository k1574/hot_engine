package main

import(
	"math"
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"hot/m/engine"
	"hot/m/engine/sprite"
	"hot/m/engine/object"
	"hot/m/engine/vector"
	"hot/m/engine/transform"
)

var(
	eng *engine.Engine
)

type GopherPlayer struct {
	O *object.Object
	MoveSpeed float64
}

type CameraPlayer struct {
	RotationSpeed float64
	PositionSpeed float64
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
			click := win.MousePosition().Add(cam.T.P)
			direction := click.Sub(*pos)
			*angle = math.Atan(direction.Y/direction.X)
			if direction.X < 0 { *angle += math.Pi }

		}
		if win.Pressed(pixelgl.KeyW){
			cam.T.P.Y += movSpeed*dt
		}
		if win.Pressed(pixelgl.KeyS){
			cam.T.P.Y -= movSpeed*dt
		}
		if win.Pressed(pixelgl.KeyD){
			cam.T.P.X += movSpeed*dt
		}
		if win.Pressed(pixelgl.KeyA){
			cam.T.P.X -= movSpeed*dt
		}
}

func
(c *CameraPlayer)Start() {
}

func
(c *CameraPlayer)Update() {
	var(
		vec vector.Vector
		r float64
	)
	t := &(eng.Cam.T)
	dt := eng.DT
	win := eng.Win
	vec = pixel.V(0, 1)
	one := float64(0)

	if win.Pressed(pixelgl.KeyPageUp){
		t.R += dt * c.RotationSpeed
	}
	if win.Pressed(pixelgl.KeyPageDown){
		t.R -= dt * c.RotationSpeed
	}


	r = -t.R
	if win.Pressed(pixelgl.KeyUp){
		one = 1
	}
	if win.Pressed(pixelgl.KeyDown){
		one = 1
		r += math.Pi
	}
	if win.Pressed(pixelgl.KeyLeft){
		one = 1
		r += math.Pi/2
	}
	if win.Pressed(pixelgl.KeyRight){
		one = 1
		r -= math.Pi/2
	}

	if one != 0 {
		vec = vec.Rotated(r)
		fmt.Println(r)
		fmt.Println(vec)
		t.P = t.P.Add(vec.Scaled(dt * c.PositionSpeed))
	}
}

func (g *GopherPlayer)GetO() *object.Object {return g.O}
func (g *CameraPlayer)GetO() *object.Object {return nil}

func
main(){
	eng = engine.New(pixelgl.WindowConfig{
		Title: "Gopher Test",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync: true,
	})

	goph_sprite, err := sprite.Load("media/hiking.png")
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
		),
		MoveSpeed: 200.0,
	}
	eng.AddBehaviorer(&goph_player)
	
	goph_player1 := GopherPlayer{
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
