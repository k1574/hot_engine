package engine

import (
	"container/list"
	"time"
	//_ "image/png"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Transform struct {
	P, S pixel.Vec
	R float64
}

type Behaviorer interface {
	Start()
	Update()
	GetO() *Object
}

type Object struct {
	E *Engine
	T Transform
	S *pixel.Sprite
}

type Camera struct {
	T Transform
}

type Engine struct {
	lastTime time.Time
	Objects *list.List
	DT float64
	WinCfg pixelgl.WindowConfig
	Win *pixelgl.Window
	Cam *Camera
}

func
(eng *Engine)update(){
		eng.Win.Clear(colornames.Whitesmoke)
		eng.setNewDT()
		for e := eng.Objects.Front() ; e != nil ; e = e.Next() {
			o := e.Value.(Behaviorer)
			o.Update()
			od := o.GetO()
			finmat := pixel.IM.ScaledXY(pixel.ZV, od.T.S).
				Rotated(pixel.ZV, od.T.R).
				Moved(od.T.P).
				ScaledXY(pixel.ZV, od.E.Cam.T.S).
				Rotated(od.E.Win.Bounds().Center(), od.E.Cam.T.R).
				Moved(pixel.ZV.Sub(od.E.Cam.T.P))
			if od.S != nil {
				od.S.Draw(eng.Win, finmat)
			}
		}
		eng.Win.Update()
}

func
(eng *Engine)setNewDT(){
	eng.DT = time.Since(eng.lastTime).Seconds()
	eng.lastTime = time.Now()
}

func
(eng *Engine)AddObject(v Behaviorer) {
	eng.Objects.PushBack(v)
	v.GetO().E = eng
	v.Start()
}

func
New(cfg pixelgl.WindowConfig) (*Engine) {
	eng := Engine {
		Objects: list.New(),
		WinCfg: cfg,
		Cam: &Camera{
			T: Transform {
				S: pixel.Vec{1, 1},
				R: 0,
				P: pixel.Vec{0, 0},
			},
		},
	}

	return &eng
}

func
(eng *Engine)run() {
	var err error

	eng.Win, err = pixelgl.NewWindow(eng.WinCfg)
	if err != nil {
		panic(err)
	}
	eng.Win.SetSmooth(true)

	eng.lastTime = time.Now()
	for !eng.Win.Closed() {
		eng.update()
	}
}

func
(eng *Engine)Run() {
	pixelgl.Run(eng.run)
}

