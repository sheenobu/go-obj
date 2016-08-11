package gfx2

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

// A Sheet is a collection of sprites
type Sheet struct {
	texture   *sdl.Texture
	positions []sdl.Rect
}

// NewSheet creates a new sheet given the image file
func NewSheet(r *sdl.Renderer, imgFile string) *Sheet {

	t, err := img.LoadTexture(r, imgFile)
	if err != nil {
		panic(err)
	}

	return &Sheet{
		texture:   t,
		positions: make([]sdl.Rect, 0),
	}
}

// Add adds specific region as a referenceable area and returns the reference ID
func (s *Sheet) Add(r *sdl.Rect) (idx int) {
	s.positions = append(s.positions, *r)
	return len(s.positions) - 1
}

// Copy copies the given reference index from the sprite sheet onto the renderer
func (s *Sheet) Copy(r *sdl.Renderer, dest *sdl.Rect, idx int, scale int32) error {
	dest.W = s.positions[idx].W * scale
	dest.H = s.positions[idx].H * scale
	return r.Copy(s.texture, &s.positions[idx], dest)
}
