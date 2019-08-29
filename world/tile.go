package world

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

type Tile struct {
	terrainTile TerrainTile
	sinceChange int
	quantity    int
}
