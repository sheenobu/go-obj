package gamekit

import (
	"sync"

	"github.com/sheenobu/go-gamekit/pair"
	"github.com/sheenobu/rxgen/rx"

	"github.com/veandco/go-sdl2/sdl"
)

// The WindowManager is responsible for handling a group of windows with different event handlers
type WindowManager struct {
	WindowCount     *rx.Uint32
	CurrentWindowID *rx.Uint32

	lock     sync.RWMutex
	windows  map[uint32]*Window
	handlers map[uint32]func(e sdl.Event)
}

// NewWindowManager builds a new window manager
func NewWindowManager() *WindowManager {
	return &WindowManager{
		CurrentWindowID: rx.NewUint32(0),
		WindowCount:     rx.NewUint32(0),

		windows:  make(map[uint32]*Window),
		handlers: make(map[uint32]func(e sdl.Event)),
	}
}

// NewWindow creates a new window, returning the window
func (wm *WindowManager) NewWindow(title string, w int, h int, extraFlags int) (*Window, error) {

	wm.lock.Lock()
	defer wm.lock.Unlock()

	window, renderer, err := sdl.CreateWindowAndRenderer(w, h, 0)

	if err != nil {
		return nil, err
	}

	window.SetTitle(title)

	win := &Window{
		ID:       window.GetID(),
		Size:     pair.NewRxInt32Pair(pair.Int32Pair{L: int32(w), R: int32(h)}),
		Window:   window,
		Renderer: renderer,
		Mouse:    NewMouse(),

		wm: wm,
	}

	wm.windows[window.GetID()] = win

	wm.handlers[window.GetID()] = func(e sdl.Event) {
		switch t := e.(type) {
		case *sdl.WindowEvent:
			if t.Event == sdl.WINDOWEVENT_CLOSE {
				window.Destroy()
			}
		}
	}

	wm.WindowCount.Set(wm.WindowCount.Get() + 1)

	return win, nil
}

// Destroy destroys the window
func (wm *WindowManager) Destroy(id uint32) {
	wm.lock.RLock()
	defer wm.lock.RUnlock()

	wm.windows[id].Window.Destroy()
	delete(wm.windows, id)
	delete(wm.handlers, id)
}

// SetHandler sets the event handler for the given window id
func (wm *WindowManager) SetHandler(id uint32, h func(e sdl.Event)) {
	wm.lock.Lock()
	defer wm.lock.Unlock()
	wm.handlers[id] = h
}

// DispatchEvents handles events and sends them to the required window event handlers.
func (wm *WindowManager) DispatchEvents() {
	if e := sdl.PollEvent(); e != nil {
		wm.lock.RLock()
		defer wm.lock.RUnlock()

		switch t := e.(type) {
		case *sdl.WindowEvent:

			win := wm.windows[t.WindowID]

			switch t.Event {
			case sdl.WINDOWEVENT_RESIZED:
				// if we handle these events we actually get a crash????
				// wm.windows[t.WindowID].Size.Set(pair.Int32Pair{t.Data1, t.Data2})
				return
			case sdl.WINDOWEVENT_SIZE_CHANGED:
				if win != nil {
					win.Size.Set(pair.Int32Pair{L: t.Data1, R: t.Data2})
				}
				return
			case sdl.WINDOWEVENT_FOCUS_GAINED:
				wm.CurrentWindowID.Set(t.WindowID)
				return
			case sdl.WINDOWEVENT_FOCUS_LOST:
				wm.CurrentWindowID.Set(0)
				return
			case sdl.WINDOWEVENT_CLOSE:
				wm.handlers[t.WindowID](e)
				wm.WindowCount.Set(wm.WindowCount.Get() - 1)
				return
			}

		case *sdl.MouseMotionEvent:
			win := wm.windows[t.WindowID]
			if win != nil {
				win.Mouse.Position.Set(pair.Int32Pair{L: t.X, R: t.Y})
			}
		case *sdl.MouseButtonEvent:
			win := wm.windows[t.WindowID]

			if win != nil {
				mouse := win.Mouse
				switch t.Button {
				case sdl.BUTTON_LEFT:
					go mouse.LeftButtonState.Set(t.State == sdl.PRESSED)
				case sdl.BUTTON_RIGHT:
					go mouse.RightButtonState.Set(t.State == sdl.PRESSED)
				}
			}
		case *sdl.QuitEvent:
			return
		}

		cwd := wm.CurrentWindowID.Get()
		if cwd == 0 {
			return
		}

		wm.handlers[cwd](e)
	}
}
