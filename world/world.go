package world

import (
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	foodTime int = 100
)

type TerrainLayout [][]TerrainTile

// World holds all data and functions for the simulation surroundings
type World struct {
	baseSheet   pixel.Picture
	grassSprite *pixel.Sprite
	waterSprite *pixel.Sprite
	foodSprite  *pixel.Sprite
	tiles       [][]*Tile
}

func (w *World) Update(frame int) {
	for _, row := range w.tiles {
		for _, tile := range row {
			if tile.terrainTile.CanChange() {
				switch tile.terrainTile {
				case Grass:
					if tile.sinceChange > foodTime {
						tile.terrainTile = Food
						tile.sinceChange = 0
						tile.quantity = 1
					}
				case Food:
					if tile.sinceChange > foodTime*3 || tile.quantity == 0 {
						tile.terrainTile = Grass
						tile.sinceChange = 0
						tile.quantity = 0
					}
					if frame%8 == 0 {
						tile.quantity++
					}
				}
			}
		}
	}
}

func (w *World) Draw(win *pixelgl.Window) {
	batch := pixel.NewBatch(&pixel.TrianglesData{}, w.baseSheet)
	mat := pixel.IM
	for r, row := range w.tiles {
		for c, tile := range row {
			var s *pixel.Sprite
			switch tile.terrainTile {
			case Grass:
				s = w.grassSprite
			case Food:
				s = w.foodSprite
			case Water:
				s = w.waterSprite
			default:
				s = w.grassSprite
			}

			mat2 := mat.Moved(win.Bounds().Min.Add(pixel.V(tileEdge*float64(r), tileEdge*float64(c))))
			s.Draw(batch, mat2)
		}
	}
	batch.Draw(win)
}

func createTiles(tl TerrainLayout) [][]*Tile {
	tiles := make([][]*Tile, 0, len(tl))
	for i, r := range tl {
		tiles = append(tiles, make([]*Tile, 0, len(r)))
		for _, tt := range r {
			tiles[i] = append(tiles[i], &Tile{
				terrainTile: tt,
				sinceChange: rand.Intn(foodTime),
			})
		}
	}
	return tiles
}

// Config holds some properties to create a new world from
type Config struct {
	TilePicture pixel.Picture
	Layout      TerrainLayout
}

// NewWorld will return a new initialized World object
func NewWorld(cfg Config) (*World, error) {
	w := &World{
		baseSheet:   cfg.TilePicture,
		grassSprite: pixel.NewSprite(cfg.TilePicture, pixel.R(0, tileEdge*2, tileEdge, tileEdge*3)),
		waterSprite: pixel.NewSprite(cfg.TilePicture, pixel.R(tileEdge, tileEdge*2, tileEdge*2, tileEdge*3)),
		foodSprite:  pixel.NewSprite(cfg.TilePicture, pixel.R(tileEdge*2, tileEdge*2, tileEdge*3, tileEdge*3)),
		tiles:       createTiles(cfg.Layout),
	}
	return w, nil
}
