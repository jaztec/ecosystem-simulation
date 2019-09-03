package runtime

import (
	"context"
	"sync"
	"time"

	"github.com/faiface/pixel/pixelgl"
)

type AppContext interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}

	Win() *pixelgl.Window
	GetValue(string) interface{}
	SetValue(string, interface{})
}

type appContext struct {
	ctx       context.Context
	frame     uint8
	deltaTime float64
	win       *pixelgl.Window
	internals map[string]interface{}
	mutex     *sync.RWMutex
}

func (ac *appContext) Deadline() (deadline time.Time, ok bool) {
	return ac.Deadline()
}

func (ac *appContext) Done() <-chan struct{} {
	return ac.ctx.Done()
}

func (ac *appContext) Err() error {
	return ac.ctx.Err()
}

func (ac *appContext) Value(key interface{}) interface{} {
	return ac.ctx.Value(key)
}

func (ac *appContext) GetValue(k string) (r interface{}) {
	ac.mutex.RLock()
	defer ac.mutex.RUnlock()

	if v, ok := ac.internals[k]; ok {
		return v
	}
	return
}

func (ac *appContext) SetValue(k string, v interface{}) {
	ac.mutex.Lock()
	defer ac.mutex.Unlock()

	ac.internals[k] = v
}

//func (ac *appContext) Frame() uint8 {
//	return ac.frame
//}
//
//func (ac *appContext) SetFrame(f uint8) {
//	ac.frame = f
//}
//
//func (ac *appContext) DeltaTime() float64 {
//	return ac.deltaTime
//}
//
//func (ac *appContext) SetDeltaTime(dt float64) {
//	ac.deltaTime = dt
//}

func (ac *appContext) Win() *pixelgl.Window {
	return ac.win
}

func FromContext(ctx context.Context, win *pixelgl.Window, params map[string]interface{}) AppContext {
	ap := &appContext{
		ctx:       ctx,
		win:       win,
		frame:     0,
		deltaTime: 0,
		internals: make(map[string]interface{}, 0),
		mutex:     &sync.RWMutex{},
	}
	return ap
}
