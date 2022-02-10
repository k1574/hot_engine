package main

import(
	"hot_engine/m/engine"
	"hot_engine/m/sprite"
	"math"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
)

type GopherPlayer struct {
	O *engine.Object
	MoveSpeed float64
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
		if win.Pressed(pixelgl.KeyUp){
			cam.T.P.Y += movSpeed*dt
		}
		if win.Pressed(pixelgl.KeyDown){
			cam.T.P.Y -= movSpeed*dt
		}
		if win.Pressed(pixelgl.KeyRight){
			cam.T.P.X += movSpeed*dt
		}
		if win.Pressed(pixelgl.KeyLeft){
			cam.T.P.X -= movSpeed*dt
		}
}

func (g *GopherPlayer)GetO() *engine.Object {return g.O}

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

	eng.Run()
}
