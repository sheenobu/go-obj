package ui

import (
	"github.com/sheenobu/go-gamekit"
	"github.com/sheenobu/go-gamekit/gfx2"
	"github.com/sheenobu/rxgen/rx"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/net/context"
)

// A ToggleButton is a named button that is part of a toggle group
type ToggleButton struct {
	name     string
	button   *Button
	checkbox *gfx2.Sprite

	isSelected bool
}

// NewToggleButton creates a new toggle button
func NewToggleButton(name string, r *sdl.Rect, checkboxOffset *sdl.Rect, sheet *gfx2.Sheet, textureID int, checkboxID int) *ToggleButton {
	return &ToggleButton{
		name:     name,
		button:   NewButton(r, sheet, textureID),
		checkbox: gfx2.NewSprite(sdl.Rect{X: checkboxOffset.X + r.X, Y: checkboxOffset.Y + r.Y, W: checkboxOffset.W, H: checkboxOffset.H}, sheet, checkboxID, 2),
	}
}

// Run runs the toggle button event processes
func (tb *ToggleButton) Run(ctx context.Context, m *gamekit.Mouse, selected *rx.String) {
	clickSub := tb.button.Clicked.Subscribe()
	defer clickSub.Close()

	selectedSub := selected.Subscribe()
	defer selectedSub.Close()

	go tb.button.Run(ctx, m)

	for {
		select {
		case <-ctx.Done():
			return
		case name := <-selectedSub.C:
			tb.isSelected = name == tb.name
		case clicked := <-clickSub.C:
			if clicked {
				go selected.Set(tb.name)
			}
		}
	}
}

// Render draws the toggle button to the renderer
func (tb *ToggleButton) Render(r *sdl.Renderer) {

	tb.button.Render(r)

	if tb.isSelected {
		tb.checkbox.Render(r)
	}
}
