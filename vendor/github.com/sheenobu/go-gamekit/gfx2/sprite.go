package gfx2

import "github.com/veandco/go-sdl2/sdl"

// A Sprite is an object represented by a texture and a position
type Sprite struct {
	Position  sdl.Rect
	textureID int
	sheet     *Sheet
	scale     int32
}

// NewSprite creates a new sprite at the given position
func NewSprite(pos sdl.Rect, sheet *Sheet, textureID int, scale int32) *Sprite {
	return &Sprite{
		Position:  pos,
		sheet:     sheet,
		textureID: textureID,
		scale:     scale,
	}
}

// Render renders the sprite
func (sp *Sprite) Render(r *sdl.Renderer) error {
	return sp.sheet.Copy(r, &sp.Position, sp.textureID, sp.scale)
}
