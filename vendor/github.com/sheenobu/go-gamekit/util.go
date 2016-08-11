package gamekit

import (
	"sync"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"golang.org/x/net/context"
)

// CountHistogram renders a time series of data points
type CountHistogram struct {
	sync.RWMutex

	points []int

	startTime time.Time
	stopTime  time.Time
}

// Run runs the count histogram
func (c *CountHistogram) Run(ctx context.Context, interval time.Duration, fn func() int) {

	t := time.NewTicker(interval)
	defer t.Stop()
	c.startTime = time.Now()
	for {
		select {
		case <-t.C:
			c.Lock()
			c.points = append(c.points, fn())
			if len(c.points) >= 30 {
				c.points = c.points[1:30]
			}
			c.Unlock()
			c.stopTime = time.Now()
		case <-ctx.Done():
			return
		}
	}
}

// Render draws the histogram to the SDL renderer
func (c *CountHistogram) Render(r *sdl.Renderer) {
	c.RLock()
	defer c.RUnlock()
	//TODO: make height logorithmic
	//TODO: make bottom of histogram the bottom of the window
	for idx, val := range c.points {
		r.SetDrawColor(255, 255, 255, 255)

		r.FillRect(&sdl.Rect{
			X: int32(20 + idx*20),
			Y: int32(500 - val),
			W: 20,
			H: int32(val),
		})

		if idx%2 == 0 {
			r.SetDrawColor(0, 255, 0, 255)
		} else {
			r.SetDrawColor(0, 255, 255, 255)
		}
		r.DrawRect(&sdl.Rect{
			X: int32(20 + idx*20),
			Y: int32(500 - val),
			W: 20,
			H: int32(val),
		})

	}
}
