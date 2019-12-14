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

func TileAt(w *World, pos pixel.Vec) *Tile {
	x, y := pos.X, pos.Y

	r := int(y / tileEdge)

	if r < 0 {
		r = 0
	}
	if r >= len(w.tiles) {
		r = len(w.tiles) - 1
	}
	row := w.tiles[r]

	c := int(x / tileEdge)
	if c < 0 {
		c = 0
	}
	if c >= len(row) {
		c = len(row) - 1
	}
	tile := row[c]

	return tile
}

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

func (w *World) Camera() Camera {
	return w.camera
}

func (w *World) Bounds() pixel.Rect {
	if len(w.tiles) == 0 {
		return pixel.Rect{}
	}
	y := float64(len(w.tiles)) * tileEdge
	x := float64(len(w.tiles[0])) * tileEdge
	// using tileEdge/2 because edges are placed centered
	return pixel.R(-tileEdge/2, -tileEdge/2, x-tileEdge/2, y-tileEdge/2)
}

func (w *World) TilesInProximity(pos pixel.Vec, radius float64) TilesInProximity {
	// get local copies of coords, make sure to take the world offset into account
	x, y := pos.X+(tileEdge/2), pos.Y+(tileEdge/2)
	tile := TileAt(w, pixel.Vec{X: x, Y: y})

	// auto-set max foreseeable range
	proximity := make([]TileType, 0, 8)

	//// do a simple perimeter check
	minX := x - radius
	maxX := x + radius
	minY := y - radius
	maxY := y + radius
	points := make([]pixel.Vec, 8)
	points[0] = pixel.Vec{X: minX, Y: y}
	points[1] = pixel.Vec{X: maxX, Y: y}
	points[2] = pixel.Vec{X: x, Y: minY}
	points[3] = pixel.Vec{X: x, Y: maxY}
	points[4] = pixel.Vec{X: minX, Y: minY}
	points[5] = pixel.Vec{X: maxX, Y: maxY}
	points[6] = pixel.Vec{X: minX, Y: maxY}
	points[7] = pixel.Vec{X: maxX, Y: minY}

	for _, p := range points {
		proximity = append(proximity, TileType{
			Pos:  p,
			Tile: TileAt(w, p),
		})
	}

	return TilesInProximity{
		On: TileType{
			Pos:  pos,
			Tile: tile,
		},
		Proximity: proximity,
	}
}

func (w *World) Update(ctx runtime.AppContext) {
	dt := ctx.GetValue("deltaTime").(float64)

	if ctx.Win().Pressed(pixelgl.KeyLeft) {
		w.camera.SetPosX(w.camera.Pos().X - (w.camera.Speed() * dt))
	}
	if ctx.Win().Pressed(pixelgl.KeyRight) {
		w.camera.SetPosX(w.camera.Pos().X + (w.camera.Speed() * dt))
	}
	if ctx.Win().Pressed(pixelgl.KeyDown) {
		w.camera.SetPosY(w.camera.Pos().Y - (w.camera.Speed() * dt))
	}
	if ctx.Win().Pressed(pixelgl.KeyUp) {
		w.camera.SetPosY(w.camera.Pos().Y + (w.camera.Speed() * dt))
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
					if ctx.GetValue("frame").(uint8)%8 == 0 {
						tile.quantity++
					}
				}
			}
		}
	}
}

func (w *World) Draw(ctx runtime.AppContext) {
	win := ctx.Win()
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

				mat2 := mat.Moved(win.Bounds().Min.Add(pixel.V(tileEdge*float64(c), tileEdge*float64(r))))
				s.Draw(w.batch, mat2)
			}
		}
		w.needsRedraw = false
	}
	w.batch.Draw(win)
	win.SetMatrix(pixel.IM.Moved(win.Bounds().Center().Sub(w.camera.Pos())))
}

func createTiles(tl TerrainLayout) [][]*Tile {
	tiles := make([][]*Tile, 0, len(tl))
	for i, r := range tl {
		tiles = append(tiles, make([]*Tile, 0, len(r)))
		for x, tt := range r {
			tiles[i] = append(tiles[i], &Tile{
				terrainTile: tt,
				sinceChange: rand.Intn(foodTime),
				center:      pixel.V((float64(i)*tileEdge)-(tileEdge/2), (float64(x)*tileEdge)-(tileEdge/2)),
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
	CamPosition pixel.Vec
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
		camera:      NewCamera(cfg.CamPosition),
	}
	return w, nil
}
