package main

import(
	"fmt"
	"math"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"github.com/surdeus/hot/src/engx"
	"github.com/surdeus/hot/src/engx/picx"
)

var(
	counter int
	eng *engx.Engine
	goph_player1 GopherPlayer
	goph_picture *pixel.Picture
	goph_batch *pixel.Batch
	goph_sprite *engx.Sprite
)
type Gopher struct {
	*engx.Object
}

type GopherPlayer struct {
	*engx.Object
	MoveSpeed float64
}

type CameraPlayer struct {
	*engx.Object
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
		pos := &(g.O().T.P)
		angle := &(g.O().T.R)
		movSpeed := g.MoveSpeed
		win := eng.Win
		cam := eng.Cam
		dt := eng.DT
	if win.Pressed(pixelgl.MouseButtonLeft){
		var real engx.Vector
		click := win.MousePosition()
		if !g.O().Floating {
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
		var real engx.Vector
		click := win.MousePosition()
		real = cam.FromRealToAbsVector(click)
		goph := Gopher{}
		goph.Object = &engx.Object{}
		goph.O().S = goph_sprite
		goph.O().P = goph_picture
		goph.O().B = goph_batch
		goph.O().T = engx.Transform{
			P: real,
			S: engx.Vector{.25, .25},
			R: 0,
		}
		eng.AddBehaviorer(&goph)
		counter++
		fmt.Println(counter)
	}
	x, y := goph_player1.O().T.P.XY()
	t.P = engx.V(x-512, y-384)
}

func
main(){
	var err error
	eng = engx.New(pixelgl.WindowConfig{
		Title: "Gopher Test",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync: true,
	})

	goph_pic, err := picx.Load("media/player1.png")
	goph_picture = &goph_pic
	goph_batch = pixel.NewBatch(&pixel.TrianglesData{}, *goph_picture)
	if err != nil {
		panic(err)
	}

	goph_sprite = pixel.NewSprite(*goph_picture, (*goph_picture).Bounds())

	goph_player := GopherPlayer{
		Object: engx.O(
			engx.T(
				eng.WinCfg.Bounds.Center(),
				pixel.Vec{1, 0.33},
				0,
			),
			goph_picture,
			goph_sprite,
			goph_batch,
			false,
		),
		MoveSpeed: 0. ,
	}
	eng.AddBehaviorer(&goph_player)
	
	goph_player1 = GopherPlayer{
		Object: &engx.Object {
			T: engx.Transform {
				P: pixel.ZV,
				S: pixel.Vec{0.35, 0.35},
			},
			P: goph_picture,
			S: goph_sprite,
			B: goph_batch,
		},
		MoveSpeed: 100.0,
	}
	eng.AddBehaviorer(&goph_player1)
	
	cam := CameraPlayer {
		RotationSpeed: 4,
		PositionSpeed: 200,
	}
	eng.AddBehaviorer(&cam)

	eng.AddBatch(goph_batch)

	eng.Run()
}
