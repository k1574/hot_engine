import(
	"hot_engine/m/engine"
)

func
(g *GopherPlayer)Start() {
}

func
(g *GopherPlayer)Update() {
		pos := &(g.O.T.P)
		angle := &(g.O.T.R)
		movSpeed := g.MoveSpeed
		if win.Pressed(pixelgl.MouseButtonLeft){
			click := win.MousePosition()
			direction := click.Sub(*pos) 
			*angle = math.Atan(direction.Y/direction.X)
			if direction.X < 0 { *angle += math.Pi }

		}
		if win.Pressed(pixelgl.KeyUp){
			pos.Y += movSpeed*dt
		}
		if win.Pressed(pixelgl.KeyDown){
			pos.Y -= movSpeed*dt
		}
		if win.Pressed(pixelgl.KeyRight){
			pos.X += movSpeed*dt
		}
		if win.Pressed(pixelgl.KeyLeft){
			pos.X -= movSpeed*dt
		}
}

func
(g *GopherPlayer)GetOD() *Object {
	return &g.O
}

func
main(){
	eng = engine.New(pixelgl.WindowConfig{
		Title: "Gopher Test",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync: true,
	})

	eng.Run()
}