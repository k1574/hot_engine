package behaviorer

import(
	"github.com/surdeus/hot/src/engine/object"
)

type Behaviorer interface {
	Update()
	Start()
	O() *object.Object
}
