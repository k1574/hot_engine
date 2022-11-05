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
	Batches []*pixel.Batch
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

		for _, b := range eng.Batches {
			b.Clear()
		}

		// Drawing objects to it's batches.
		// Could be way simpler if we used just one batch and spritesheet,
		// but I like keeping sprites in different files, so...
		for e := eng.Objects.Front() ; e != nil ; e = e.Next() {
			o := e.Value.(behaviorer.Behaviorer)
			o.Update()

			od := o.O()
			if od == nil {
				continue
			}

			if od.P != nil {
				finmat = eng.FromAbsToRealMatrix(od)
				od.S.Draw(od.B, finmat)
			}
		}

		// Drawing batches themselves.
		for _, b := range eng.Batches {
			b.Draw(eng.Win)
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

func (eng *Engine)AddBatch(b *pixel.Batch) {
	eng.Batches = append(eng.Batches, b)
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

