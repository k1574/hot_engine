package main

import (
	"container/list"
	"image"
	"os"
	"time"
	"math"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	//"fmt"
)

type Transform struct {
	P, S pixel.Vec
	R float64
}

type Behaviorer interface {
	Start()
	Update()
	GetOD() *Object
}

type Object struct {
	T Transform
	S *pixel.Sprite
}

type Camera struct {
	T Transform
}

type GopherPlayer struct {
	O Object
	MoveSpeed float64
}

type Engine struct {
	lasttime time.Time
	objects *list.List
	DT float64
	wcfg pixelgl.WindowConfig
	win *pixelgl.Window
}

func
(engine *Engine)Update(){
		for e := engine.objects.Front() ; e != nil ; e = e.Next() {
			o := e.Value.(Behaviorer)
			o.Update()
			od := o.GetOD()
			finmat := pixel.IM.ScaledXY(pixel.ZV, od.T.S).Rotated(pixel.ZV, od.T.R).Moved(od.T.P)
			od.S.Draw(win, finmat)
		}
}

func
(engine *Engine)setNewDT(){
	engine.dt = time.Since(engine.lasttime).Seconds()
	engine.lasttime = time.Now()
}


func
AddObject(o *Behaviorer) {
	objects.PushBack(o)
	o.Start()
}

func run() {
	win, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)

	pic, err := loadPicture("../hiking.png")
	if err != nil {
		panic(err)
	}

	goph_sprite = pixel.NewSprite(pic, pic.Bounds())
	objects = list.New()

	goph_player := GopherPlayer{
		O: Object {
			T: Transform {
				P: win.Bounds().Center(),
				S: pixel.Vec{3, 1},
			},
			S: goph_sprite,
		},
		MoveSpeed: 200.0,
	}
	AddObject(&goph_player)
	
	goph_player1 := GopherPlayer{
		O: Object {
			T: Transform {
				P: pixel.ZV,
				S: pixel.Vec{1, 1},
			},
			S: goph_sprite,
		},
		MoveSpeed: 100.0,
	}
	AddObject(&goph_player1)

	lasttime = time.Now()
	for !win.Closed() {
		setDT()
		win.Clear(colornames.Whitesmoke)


		win.Update()
	}
}



func Run {
	pixelgl.Run(run)
}

