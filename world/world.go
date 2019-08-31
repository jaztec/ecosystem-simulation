package world

import (
	"math/rand"

	"github.com/jaztec/ecosystem-simulation/runtime"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	foodTime int = 10000
)

type TerrainLayout [][]TerrainTile

// World holds all data and functions for the simulation surroundings
type World struct {
	grassSprite *pixel.Sprite
	waterSprite *pixel.Sprite
	foodSprite  *pixel.Sprite
	tiles       [][]*Tile
	batch       *pixel.Batch
	camera      Camera
	needsRedraw bool
}

func (w *World) Update(ctx *runtime.AppContext) {
	if ctx.Win().Pressed(pixelgl.KeyLeft) {
		w.camera.Pos().X -= w.camera.Speed() * ctx.DeltaTime()
	}
	if ctx.Win().Pressed(pixelgl.KeyRight) {
		w.camera.Pos().X += w.camera.Speed() * ctx.DeltaTime()
	}
	if ctx.Win().Pressed(pixelgl.KeyDown) {
		w.camera.Pos().Y -= w.camera.Speed() * ctx.DeltaTime()
	}
	if ctx.Win().Pressed(pixelgl.KeyUp) {
		w.camera.Pos().Y += w.camera.Speed() * ctx.DeltaTime()
	}
	for _, row := range w.tiles {
		for _, tile := range row {
			if tile.terrainTile.CanChange() {
				tile.sinceChange++
				switch tile.terrainTile {
				case Grass:
					if tile.sinceChange > foodTime {
						tile.terrainTile = Food
						tile.sinceChange = 0
						tile.quantity = 1
						w.needsRedraw = true
					}
				case Food:
					if tile.sinceChange > foodTime*3 || tile.quantity == 0 {
						tile.terrainTile = Grass
						tile.sinceChange = 0
						tile.quantity = 0
						w.needsRedraw = true
					}
					if ctx.Frame()%8 == 0 {
						tile.quantity++
					}
				}
			}
		}
	}
}

func (w *World) Draw(win *pixelgl.Window) {
	if w.needsRedraw {
		w.batch.Clear()
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
				s.Draw(w.batch, mat2)
			}
		}
		w.needsRedraw = false
	}
	w.batch.Draw(win)
	win.SetMatrix(pixel.IM.Moved(win.Bounds().Center().Sub(*w.camera.Pos())))
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

// GenerateWorldLayout will return some randomly generated terrain layout
func GenerateWorldLayout(w, h int) TerrainLayout {
	tl := make([][]TerrainTile, 0, h)
	for x := 0; x < h; x++ {
		r := make([]TerrainTile, 0, w)
		for y := 0; y < w; y++ {
			r = append(r, TerrainTile(rand.Intn(3)))
		}
		tl = append(tl, r)
	}
	return tl
}

// Config holds some properties to create a new world from
type Config struct {
	TilePicture pixel.Picture
	Layout      TerrainLayout
}

// NewWorld will return a new initialized World object
func NewWorld(cfg Config) (*World, error) {
	w := &World{
		grassSprite: pixel.NewSprite(cfg.TilePicture, pixel.R(0, tileEdge*2, tileEdge, tileEdge*3)),
		waterSprite: pixel.NewSprite(cfg.TilePicture, pixel.R(tileEdge, tileEdge*2, tileEdge*2, tileEdge*3)),
		foodSprite:  pixel.NewSprite(cfg.TilePicture, pixel.R(tileEdge*2, tileEdge*2, tileEdge*3, tileEdge*3)),
		tiles:       createTiles(cfg.Layout),
		needsRedraw: true,
		batch:       pixel.NewBatch(&pixel.TrianglesData{}, cfg.TilePicture),
		camera:      NewCamera(),
	}
	return w, nil
}
