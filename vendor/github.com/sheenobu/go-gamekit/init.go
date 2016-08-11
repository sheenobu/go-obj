package gamekit

import "github.com/veandco/go-sdl2/sdl"

// Init initializes the gamekit engine
func Init() {
	sdl.Init(sdl.INIT_EVERYTHING)
}
