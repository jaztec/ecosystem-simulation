package application

import (
	"context"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/jaztec/ecosystem-simulation/fauna"

	"github.com/jaztec/ecosystem-simulation/runtime"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/jaztec/ecosystem-simulation/world"
	"golang.org/x/xerrors"
)

type Config struct {
	Layout world.TerrainLayout
}

// Application holds all important variables and handles the application flow
type Application struct {
	win    *pixelgl.Window
	world  *world.World
	herd   *fauna.Herd
	lastDt time.Time
}

func (a *Application) update(ctx runtime.AppContext) {
	// TODO This is a test log line
	if ctx.Win().JustPressed(pixelgl.MouseButtonLeft) {
		camPos := pixel.IM.Scaled(a.world.Camera().Pos(), 1).Moved(ctx.Win().Bounds().Center().Sub(a.world.Camera().Pos()))
		mouse := camPos.Unproject(ctx.Win().MousePosition())
		log.Printf("Clicked at pos %v", mouse)
	}
	a.world.Update(ctx)
	if ia := a.herd.Update(ctx); ia == false {
		a.win.SetClosed(true)
		return
	}
	a.draw(ctx)
}

func (a *Application) draw(ctx runtime.AppContext) {
	a.win.Clear(colornames.Skyblue)
	a.world.Draw(ctx)
	a.herd.Draw(ctx)
}

func (a *Application) params() map[string]interface{} {
	params := make(map[string]interface{}, 0)
	return params
}

// Run will run the application
func (a *Application) Run(c context.Context) {
	a.lastDt = time.Now()
	ctx := runtime.FromContext(c, a.win, a.params())
	var frame uint8 = 0
	defer fauna.PrintStats(os.Stdout, a.herd.Stats())

	for !a.win.Closed() {
		dt := time.Since(a.lastDt).Seconds()
		frame++
		a.lastDt = time.Now()
		ctx.SetValue("deltaTime", dt)
		ctx.SetValue("frame", frame)
		ctx.SetValue("world", a.world)
		a.update(ctx)
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

func createWorld(cfg Config, win *pixelgl.Window) (*world.World, error) {
	tiles, err := runtime.LoadPicture("./assets/sprites/tiles.png")
	if err != nil {
		return nil, xerrors.Errorf("fatal error loading tiles: %w", err)
	}

	worldMap, err := world.NewWorld(world.Config{
		TilePicture: tiles,
		Layout:      cfg.Layout,
		CamPosition: win.Bounds().Center(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fatal error creating the world: %w", err)
	}
	return worldMap, nil
}

func createHerd(bounds pixel.Rect) (*fauna.Herd, error) {
	sprite, err := runtime.LoadPicture("./assets/sprites/sheep.png")
	if err != nil {
		return nil, xerrors.Errorf("fatal error loading sheeps: %w", err)
	}

	herd, err := fauna.NewHerd(fauna.HerdConfig{
		SheepPicture:  sprite,
		Bounds:        bounds,
		NumberOfSheep: 15,
	})
	if err != nil {
		return nil, xerrors.Errorf("fatal error creating the herd: %w", err)
	}
	return herd, nil
}

func New(cfg Config) (*Application, error) {
	rand.Seed(int64(time.Now().Second()))
	win, err := createWindow()
	if err != nil {
		return nil, err
	}
	log.Printf("Bounds of win are %v", win.Bounds())

	worldMap, err := createWorld(cfg, win)
	if err != nil {
		return nil, err
	}
	log.Printf("Bounds of world are %v", worldMap.Bounds())

	herd, err := createHerd(worldMap.Bounds())
	if err != nil {
		return nil, err
	}

	return &Application{
		win:   win,
		world: worldMap,
		herd:  herd,
	}, nil
}
