package game

type State interface {
	Update()
	Destroy()
	Initialize()
}
