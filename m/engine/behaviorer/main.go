package behaviorer

import(
	"hot/m/engine/object"
)

type Behaviorer interface {
	Update()
	Start()
	GetO() *object.Object
}