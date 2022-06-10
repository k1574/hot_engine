package behaviorer

import(
	"github.com/k1574/hot/m/engine/object"
)

type Behaviorer interface {
	Update()
	Start()
	GetO() *object.Object
}