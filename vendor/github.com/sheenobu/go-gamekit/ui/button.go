package ui

import (
	"github.com/sheenobu/go-gamekit"
	"github.com/sheenobu/go-gamekit/gfx2"
	"github.com/sheenobu/rxgen/rx"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/net/context"
)

// NewButton creates a new button
func NewButton(r *sdl.Rect, sheet *gfx2.Sheet, textureID int) *Button {
	return &Button{
		sprite:        gfx2.NewSprite(*r, sheet, textureID, 2),
		Clicked:       rx.NewBool(false),
		clickedState:  false,
		hoveringState: false,
	}
}

// A Button is a grahical item that can be clicked
type Button struct {
	sprite *gfx2.Sprite

	Clicked *rx.Bool

	clickedState  bool
	hoveringState bool
}

// Run the event processes
func (b *Button) Run(ctx context.Context, m *gamekit.Mouse) {

	posS := m.Position.Subscribe()
	clickS := m.LeftButtonState.Subscribe()
	defer posS.Close()
	defer clickS.Close()

	for {
		select {
		case <-ctx.Done():
			return
		case pos := <-posS.C:
			x := pos.L
			y := pos.R
			p := b.sprite.Position
			b.hoveringState = x > p.X && y > p.Y && x < p.X+p.W && y < p.Y+p.H
		case leftClick := <-clickS.C:
			if b.hoveringState && leftClick {
				b.clickedState = true
				b.Clicked.Set(true)
			} else {
				b.clickedState = false
				b.Clicked.Set(false)
			}
		}
	}
}

// Render draws the button on the screen
func (b *Button) Render(r *sdl.Renderer) {
	b.sprite.Render(r)
}
