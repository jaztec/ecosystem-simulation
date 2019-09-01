package fauna

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/jaztec/ecosystem-simulation/runtime"
)

type HerdConfig struct {
	SheepPicture pixel.Picture
}

type Herd struct {
	sprites [5][4]*pixel.Sprite
	herd    []Sheep
	batch   *pixel.Batch
}

func (h *Herd) Update(ctx *runtime.AppContext) {

}

func (h *Herd) Draw(ctx *runtime.AppContext) {
	h.batch.Clear()
	win := ctx.Win()
	mat := pixel.IM
	for _, sheep := range h.herd {
		var s *pixel.Sprite

		if sheep.MovingSpeed() == 0 {
			s = h.sprites[StoppedIndex][int(sheep.Direction())]
		} else {
			switch sheep.Direction() {
			case Up:
				s = h.sprites[MovingUp][int(sheep.State())]
			case Down:
				s = h.sprites[MovingDown][int(sheep.State())]
			case Left:
				s = h.sprites[MovingLeft][int(sheep.State())]
			case Right:
				fallthrough
			default:
				s = h.sprites[MovingRight][int(sheep.State())]
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
		herd:    make([]Sheep, 0),
		batch:   pixel.NewBatch(&pixel.TrianglesData{}, cfg.SheepPicture),
	}

	var i int
	for i < 1 {
		h.herd = append(h.herd, NewSheep())
		fmt.Println(h.herd[0].Position())
		i++
	}

	return h, nil
}
