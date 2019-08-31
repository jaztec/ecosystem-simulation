package runtime

import (
	"context"
	"time"

	"github.com/faiface/pixel/pixelgl"
)

type AppContext struct {
	ctx       context.Context
	frame     uint8
	deltaTime float64
	win       *pixelgl.Window
}

func (ac *AppContext) Deadline() (deadline time.Time, ok bool) {
	return ac.Deadline()
}

func (ac *AppContext) Done() <-chan struct{} {
	return ac.ctx.Done()
}

func (ac *AppContext) Err() error {
	return ac.ctx.Err()
}

func (ac *AppContext) Value(key interface{}) interface{} {
	return ac.ctx.Value(key)
}

func (ac *AppContext) Frame() uint8 {
	return ac.frame
}

func (ac *AppContext) SetFrame(f uint8) {
	ac.frame = f
}

func (ac *AppContext) DeltaTime() float64 {
	return ac.deltaTime
}

func (ac *AppContext) SetDeltaTime(dt float64) {
	ac.deltaTime = dt
}

func (ac *AppContext) Win() *pixelgl.Window {
	return ac.win
}

func FromContext(ctx context.Context, win *pixelgl.Window) *AppContext {
	return &AppContext{ctx: ctx, win: win}
}
