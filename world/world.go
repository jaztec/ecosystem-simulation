package world

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// TerrainTile is an holder to indicate different terrains on different places
type TerrainTile uint8

var (
	// Grass tile
	Grass TerrainTile = 0x00
	// Water tile
	Water TerrainTile = 0x01
	// Food tile
	Food TerrainTile = 0x02
)

const (
	foodTime int = 100
)

type TerrainLayout [][]uint8

type Tile struct {
	T TerrainTile
}

// World holds all data and functions for the simulation surroundings
type World struct {
	grassSprite *pixel.Sprite
	waterSprite *pixel.Sprite
	foodSprite  *pixel.Sprite
	tiles       pixel.Picture
}

func (w *World) Draw(win *pixelgl.Window) {
	w.grassSprite.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(pixel.V(0, 0))))
	w.waterSprite.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(pixel.V(300, 300))))
	w.foodSprite.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(pixel.V(150, 150))))
}

// Config holds some properties to create a new world from
type Config struct {
	TilePicture pixel.Picture
	Layout      TerrainLayout
}

// NewWorld will return a new initialized World object
func NewWorld(cfg Config) (*World, error) {
	w := &World{
		grassSprite: pixel.NewSprite(cfg.TilePicture, pixel.R(0, 300, 150, 450)),
		waterSprite: pixel.NewSprite(cfg.TilePicture, pixel.R(150, 300, 300, 450)),
		foodSprite:  pixel.NewSprite(cfg.TilePicture, pixel.R(300, 300, 450, 450)),
	}
	return w, nil
}
