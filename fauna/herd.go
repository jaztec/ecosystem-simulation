package fauna

import (
	"math/rand"
	"time"

	"github.com/jaztec/ecosystem-simulation/runtime"

	"github.com/jaztec/ecosystem-simulation/world"

	"github.com/faiface/pixel"
)

const (
	maxWater = 10.0
	maxFood  = 10.0
)

type HerdConfig struct {
	SheepPicture  pixel.Picture
	Bounds        pixel.Rect
	NumberOfSheep int
}

type Herd struct {
	sprites [5][4]*pixel.Sprite
	herd    []*Sheep
	batch   *pixel.Batch
	bounds  pixel.Rect
}

func (h *Herd) Stats() []Dieable {
	r := make([]Dieable, 0, len(h.herd))
	for _, s := range h.herd {
		r = append(r, s)
	}
	return r
}

func (h *Herd) Update(ctx runtime.AppContext) {
	dt := ctx.GetValue("deltaTime").(float64)
	wo := ctx.GetValue("world").(*world.World)
	frame := ctx.GetValue("frame").(uint8)
	for i := 0; i < len(h.herd); i++ {
		sheep := h.herd[i]
		if !sheep.Alive() {
			continue
		}
		sheep.food -= dt
		sheep.water -= dt
		sheep.reproduce += dt
		sheep.age += dt

		if sheep.water <= 0 {
			sheep.alive = false
			sheep.reason = Thirst
			sheep.time = time.Time{}
			return
		}
		if sheep.food <= 0 {
			sheep.alive = false
			sheep.reason = Starvation
			sheep.time = time.Time{}
			return
		}
		if sheep.age >= float64(120)+float64(rand.Intn(50)) {
			sheep.alive = false
			sheep.reason = Age
			sheep.time = time.Time{}
			return
		}

		tip := wo.TilesInProximity(sheep.Position(), sheep.Vision())
		switch tip.On.Tile.Terrain() {
		case world.Food:
			sheep.food += tip.On.Tile.DeductQuantity(maxFood - sheep.Food())
		case world.Water:
			sheep.water += maxWater - sheep.Water()
		}

		if sheep.water < (maxWater/100.0)*30.0 {
			if ok, _, onSpot := tip.HasTileType(world.Water); ok || onSpot {
				// head for the water
			}
		} else if sheep.food < (maxFood/100.0)*30.0 {
			if ok, _, onSpot := tip.HasTileType(world.Food); ok || onSpot {
				// head for the food
			}
		} else {
			randomMovement(sheep, h.bounds)
		}

		if frame%8 == 0 {
			sheep.step++
			if sheep.step > 3 {
				sheep.step = 0
			}
		}
	}
	return
}

func (h *Herd) Draw(ctx runtime.AppContext) {
	h.batch.Clear()
	win := ctx.Win()
	mat := pixel.IM
	for _, sheep := range h.herd {
		if !sheep.Alive() {
			continue
		}

		var s *pixel.Sprite

		if sheep.MovingSpeedX() == 0 {
			s = h.sprites[StoppedIndex][int(sheep.Direction())]
		} else {
			switch sheep.Direction() {
			case Up:
				s = h.sprites[MovingUp][int(sheep.Step())]
			case Down:
				s = h.sprites[MovingDown][int(sheep.Step())]
			case Left:
				s = h.sprites[MovingLeft][int(sheep.Step())]
			case Right:
				fallthrough
			default:
				s = h.sprites[MovingRight][int(sheep.Step())]
			}
		}

		mat2 := mat.Moved(sheep.position)
		s.Draw(h.batch, mat2)
	}

	h.batch.Draw(win)
}

func NewHerd(cfg HerdConfig) (*Herd, error) {
	h := &Herd{
		sprites: extractSprites(cfg.SheepPicture),
		herd:    make([]*Sheep, 0),
		batch:   pixel.NewBatch(&pixel.TrianglesData{}, cfg.SheepPicture),
		bounds:  cfg.Bounds,
	}

	var i int
	for i < cfg.NumberOfSheep {
		s := NewSheep()
		s.position.X = float64(rand.Intn(int(cfg.Bounds.Max.X)))
		s.position.Y = float64(rand.Intn(int(cfg.Bounds.Max.Y)))
		s.direction = Direction(rand.Intn(4))
		s.gender = Gender(rand.Intn(2))
		h.herd = append(h.herd, &s)
		i++
	}

	return h, nil
}
