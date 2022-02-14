package main

import(
	"hot_engine/m/engine"
	"hot_engine/m/sprite"
	"math"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"fmt"
)

type GopherPlayer struct {
	O *engine.Object
	MoveSpeed float64
}

type Camera struct {
	O *engine.Object
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
		win := g.O.E.Win
		cam := g.O.E.Cam
		dt := g.O.E.DT
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
(c *Camera)Start() {
}

func
(c *Camera)Update() {
	var(
		vec pixel.Vec
		r float64
	)
	t := &(c.O.E.Cam.T)
	dt := c.O.E.DT
	win := c.O.E.Win
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

func (g *GopherPlayer)GetO() *engine.Object {return g.O}
func (g *Camera)GetO() *engine.Object {return g.O}

func
main(){
	eng := engine.New(pixelgl.WindowConfig{
		Title: "Gopher Test",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync: true,
	})

	goph_sprite, err := sprite.LoadSprite("media/hiking.png")
	if err != nil {
		panic(err)
	}

	goph_player := GopherPlayer{
		O: &engine.Object {
			T: engine.Transform {
				P: eng.WinCfg.Bounds.Center(),
				S: pixel.Vec{1, 0.33},
			},
			S: goph_sprite,
		},
		MoveSpeed: 200.0,
	}
	eng.AddObject(&goph_player)
	
	goph_player1 := GopherPlayer{
		O: &engine.Object {
			T: engine.Transform {
				P: pixel.ZV,
				S: pixel.Vec{0.35, 0.35},
			},
			S: goph_sprite,
		},
		MoveSpeed: 100.0,
	}
	eng.AddObject(&goph_player1)
	
	cam := Camera {
		O: &engine.Object{},
		RotationSpeed: 4,
		PositionSpeed: 200,
	}
	eng.AddObject(&cam)

	eng.Run()
}
