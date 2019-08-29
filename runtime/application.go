package runtime

import (
	"context"
	"time"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/jaztec/ecosystem-simulation/world"
	"golang.org/x/xerrors"
)

type ApplicationConfig struct {
	Layout world.TerrainLayout
}

// Application holds all important variables and handles the application flow
type Application struct {
	win    *pixelgl.Window
	world  *world.World
	lastDt time.Time
}

func (a *Application) update(frame int) {
	// TODO improve this quick hack
	frame++
	if frame > 10 {
		frame = 0
	}
	a.world.Update(frame)
	a.draw()
}

func (a *Application) draw() {
	a.win.Clear(colornames.Skyblue)
	a.world.Draw(a.win)
}

// Run will run the application
func (a *Application) Run(_ context.Context) {
	a.lastDt = time.Now()
	// TODO create some derived context
	frame := 0
	for !a.win.Closed() {
		// TODO fix delta time
		_ = time.Since(a.lastDt).Seconds()
		a.lastDt = time.Now()
		a.update(frame)
		a.win.Update()
	}
}

func createWindow() (*pixelgl.Window, error) {
	cfg := pixelgl.WindowConfig{
		Title:       "Ecosystem simulation",
		Icon:        nil,
		Bounds:      pixel.R(0, 0, 1024, 768),
		Monitor:     nil,
		Resizable:   true,
		Undecorated: false,
		VSync:       true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		return nil, xerrors.Errorf("fatal error creating a new window: %w", err)
	}
	win.SetSmooth(true)
	return win, nil
}

func createWorld(cfg ApplicationConfig) (*world.World, error) {
	tiles, err := LoadPicture("./assets/sprites/tiles.png")
	if err != nil {
		return nil, xerrors.Errorf("fatal error loading tiles: %w", err)
	}

	worldMap, err := world.NewWorld(world.Config{TilePicture: tiles, Layout: cfg.Layout})
	if err != nil {
		return nil, xerrors.Errorf("fatal error creating the world: %w", err)
	}
	return worldMap, nil
}

func NewApplication(cfg ApplicationConfig) (*Application, error) {
	win, err := createWindow()
	if err != nil {
		return nil, err
	}

	worldMap, err := createWorld(cfg)
	if err != nil {
		return nil, err
	}

	return &Application{
		win:   win,
		world: worldMap,
	}, nil
}
