package engine

import (
	"container/list"
	"time"
	//_ "image/png"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"github.com/surdeus/hot/src/engine/camera"
	"github.com/surdeus/hot/src/engine/transform"
	"github.com/surdeus/hot/src/engine/vector"
	"github.com/surdeus/hot/src/engine/behaviorer"
	"github.com/surdeus/hot/src/engine/matrix"
	"github.com/surdeus/hot/src/engine/object"
)

type Engine struct {
	lastTime time.Time
	Objects *list.List
	DT float64
	WinCfg pixelgl.WindowConfig
	Win *pixelgl.Window
	Cam *camera.Camera
}

func
(eng *Engine)update(){
		var finmat = matrix.I
		eng.Win.Clear(colornames.Whitesmoke)
		eng.setNewDT()
		for e := eng.Objects.Front() ; e != nil ; e = e.Next() {
			o := e.Value.(behaviorer.Behaviorer)
			o.Update()

			od := o.GetO()
			if od == nil {
				continue
			}

			if od.S != nil {
				finmat = eng.FromAbsToRealMatrix(od)
				od.S.Draw(eng.Win, finmat)
			}
		}
		eng.Win.Update()
}

func
(eng *Engine)FromAbsToRealMatrix(od *object.Object) matrix.Matrix {
	finmat := pixel.IM.ScaledXY(pixel.ZV, od.T.S).
		Rotated(vector.Z, od.T.R).
		Moved(od.T.P)

	/* To be able work with GUI or floating shit. */
	if !od.Floating {
		finmat = finmat.Chained(eng.Cam.FromAbsToRealMatrix())
	}

	return finmat
}


func
(eng *Engine)setNewDT(){
	eng.DT = time.Since(eng.lastTime).Seconds()
	eng.lastTime = time.Now()
}

func
(eng *Engine)AddBehaviorer(v behaviorer.Behaviorer) {
	eng.Objects.PushBack(v)
	v.Start()
}

func
New(cfg pixelgl.WindowConfig) (*Engine) {
	eng := Engine {
		Objects: list.New(),
		WinCfg: cfg,
		Cam: camera.New(
			transform.New(
				vector.New(1, 1),
				vector.New(1, 1),
				0),
			cfg.Bounds.Center()),
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

