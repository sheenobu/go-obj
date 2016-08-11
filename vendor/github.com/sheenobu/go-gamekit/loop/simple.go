package loop

import (
	"github.com/sheenobu/go-gamekit"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/net/context"
)

type simpleLoop struct {
	wm      *gamekit.WindowManager
	running bool
	render  func()
	ctx     context.Context
}

// Simple builds a simple loop, with no timing management
func Simple(wm *gamekit.WindowManager, ctx context.Context, fn func()) Loop {
	if ctx == nil {
		ctx = context.Background()
	}
	return &simpleLoop{
		wm:      wm,
		running: false,
		render:  fn,
		ctx:     ctx,
	}
}

// run the simple loop
func (sl *simpleLoop) Run() {
	sl.running = true

	go func() {
		sub := sl.wm.WindowCount.Subscribe()
		defer sub.Close()

		for {
			select {
			case <-sl.ctx.Done():
				sl.running = false
				return
			case c := <-sub.C:
				if c == 0 {
					sl.running = false
					return
				}
			}
		}
	}()

	for sl.running {
		sl.wm.DispatchEvents()
		sl.render()

		sdl.Delay(2)

	}
}
