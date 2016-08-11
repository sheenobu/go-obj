package gamekit

import (
	"github.com/sheenobu/go-gamekit/pair"
	"github.com/sheenobu/rxgen/rx"
)

// Mouse is the representation of our mouse input on the screen
type Mouse struct {
	Position         *pair.RxInt32Pair
	LeftButtonState  *rx.Bool
	RightButtonState *rx.Bool
}

// NewMouse creates a new object object
func NewMouse() *Mouse {
	return &Mouse{
		Position:         pair.NewRxInt32Pair(pair.Int32Pair{0, 0}),
		LeftButtonState:  rx.NewBool(false),
		RightButtonState: rx.NewBool(false),
	}
}
