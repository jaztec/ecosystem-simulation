package fauna

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/jaztec/ecosystem-simulation/runtime"
)

type HerdConfig struct {
	SheepPicture pixel.Picture
	Bounds       pixel.Rect
}

type Herd struct {
	sprites [5][4]*pixel.Sprite
	herd    []*Sheep
	batch   *pixel.Batch
}

func (h *Herd) Update(ctx *runtime.AppContext) {
	for i := 0; i < len(h.herd); i++ {
		sheep := h.herd[i]
		if !sheep.Alive() {
			continue
		}
		sheep.food -= ctx.DeltaTime()
		sheep.water -= ctx.DeltaTime()
		sheep.reproduce += ctx.DeltaTime()
		sheep.age += ctx.DeltaTime()

		if sheep.water <= 0 {
			sheep.alive = false
			sheep.reason = Thirst
			sheep.time = time.Time{}
		}
		if sheep.food <= 0 {
			sheep.alive = false
			sheep.reason = Starvation
			sheep.time = time.Time{}
		}
		if sheep.age >= float64(120)+float64(rand.Intn(50)) {
			sheep.alive = false
			sheep.reason = Age
			sheep.time = time.Time{}
		}
	}
}

func (h *Herd) Draw(ctx *runtime.AppContext) {
	h.batch.Clear()
	win := ctx.Win()
	mat := pixel.IM
	for _, sheep := range h.herd {
		if !sheep.Alive() {
			continue
		}

		var s *pixel.Sprite

		if sheep.MovingSpeed() == 0 {
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
	}

	var i int
	for i < 50 {
		s := NewSheep()
		s.position.X = (cfg.Bounds.Max.X - cfg.Bounds.Min.X) * (rand.Float64() + 0.5)
		s.position.Y = (cfg.Bounds.Max.Y - cfg.Bounds.Min.Y) * (rand.Float64() + 0.5)
		s.direction = Direction(rand.Intn(4))
		s.gender = Gender(rand.Intn(2))
		h.herd = append(h.herd, &s)
		i++
	}

	return h, nil
}
