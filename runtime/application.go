package runtime

import (
	"context"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/jaztec/ecosystem-simulation/world"
	"golang.org/x/xerrors"
)

// Application holds all important variables and handles the application flow
type Application struct {
	win   *pixelgl.Window
	world *world.World
}

func (a *Application) update() {
	a.win.Clear(colornames.Skyblue)
	a.draw()
}

func (a *Application) draw() {
	go a.world.Draw(a.win)
}

// Run will run the application
func (a *Application) Run(_ context.Context) {
	for !a.win.Closed() {
		a.update()
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

func createWorld() (*world.World, error) {
	tiles, err := LoadPicture("./assets/sprites/tiles.png")
	if err != nil {
		return nil, xerrors.Errorf("fatal error loading tiles: %w", err)
	}

	worldMap, err := world.NewWorld(world.Config{TilePicture: tiles})
	if err != nil {
		return nil, xerrors.Errorf("fatal error creating the world: %w", err)
	}
	return worldMap, nil
}

func NewApplication() (*Application, error) {
	win, err := createWindow()
	if err != nil {
		return nil, err
	}

	worldMap, err := createWorld()
	if err != nil {
		return nil, err
	}

	return &Application{
		win:   win,
		world: worldMap,
	}, nil
}
