package ui

import (
	"github.com/sheenobu/go-gamekit"
	"github.com/sheenobu/rxgen/rx"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/net/context"
)

// ToggleGroup is a group of buttons, only one of which may be selected
type ToggleGroup struct {
	buttons  []*ToggleButton
	Selected *rx.String

	initial string
}

// NewToggleGroup create a new toggle group, marking the the given one as the default selected
func NewToggleGroup(defaultSelected string) *ToggleGroup {

	var tg ToggleGroup
	tg.Selected = rx.NewString(defaultSelected)
	tg.initial = defaultSelected

	return &tg
}

// Add a new toggle button, must be done before Run is called
func (tg *ToggleGroup) Add(tb *ToggleButton) {
	tg.buttons = append(tg.buttons, tb)
	if tg.initial == tb.name {
		tb.isSelected = true
	}
}

// Run the backend event processes
func (tg *ToggleGroup) Run(ctx context.Context, m *gamekit.Mouse) {

	for _, b := range tg.buttons {
		go b.Run(ctx, m, tg.Selected)
	}

	<-ctx.Done()
}

// Render the toggle group buttons
func (tg *ToggleGroup) Render(r *sdl.Renderer) {
	for _, b := range tg.buttons {
		b.Render(r)
	}
}
