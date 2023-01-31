package engx

type Behaviorer interface {
	Update()
	Start()
	O() *Object
}
