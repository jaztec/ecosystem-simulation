package world

import (
	"fmt"

	"github.com/faiface/pixel"
)

const (
	tileEdge float64 = 150
)

var (
	// Grass tile
	Grass TerrainTile = 0x00
	// Water tile
	Water TerrainTile = 0x01
	// Food tile
	Food TerrainTile = 0x02
)

// TerrainTile is an holder to indicate different terrains on different places
type TerrainTile uint8

func (tt TerrainTile) CanChange() (canIt bool) {
	switch tt {
	case Grass:
		fallthrough
	case Food:
		canIt = true
	}
	return
}

type TileType struct {
	Pos  pixel.Vec
	Tile *Tile
}

type TilesInProximity struct {
	On        TileType
	Proximity []TileType
}

func (tip TilesInProximity) HasTileType(terrain TerrainTile) (bool, TileType, bool) {
	// check if on-spot
	if tip.On.Tile.terrainTile == terrain {
		return true, tip.On, true
	}
	for _, tt := range tip.Proximity {
		if tt.Tile.terrainTile == terrain {
			return true, tt, false
		}
	}

	return false, TileType{}, false
}

type Tile struct {
	terrainTile TerrainTile
	sinceChange int
	quantity    float64
	center      pixel.Vec
}

func (t *Tile) String() string {
	var tt string
	switch t.terrainTile {
	case Water:
		tt = "water"
	case Grass:
		tt = "grass"
	case Food:
		tt = "food"
	}
	return fmt.Sprintf("Tile\n========\ntype: %s\nqty: %f\npos: %s\n========\n", tt, t.quantity, t.center)
}

func (t *Tile) Terrain() TerrainTile {
	return t.terrainTile
}

func (t *Tile) Center() pixel.Vec {
	return t.center
}

func (t *Tile) DeductQuantity(v float64) float64 {
	if r := t.quantity; r < v {
		t.quantity = 0.0
		return r
	}
	t.quantity -= v
	return v
}
