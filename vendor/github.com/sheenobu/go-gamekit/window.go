package gamekit

import (
	"github.com/sheenobu/go-gamekit/pair"

	"github.com/veandco/go-sdl2/sdl"
)

// Window represents a managed window
type Window struct {
	ID       uint32
	Size     *pair.RxInt32Pair
	Window   *sdl.Window
	Renderer *sdl.Renderer
	Mouse    *Mouse

	wm *WindowManager
}
