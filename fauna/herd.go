package fauna

import (
	"github.com/faiface/pixel"
	"github.com/jaztec/ecosystem-simulation/runtime"
)

type HerdConfig struct {
	SheepPicture pixel.Picture
}

type Herd struct {
	sprites [5][4]*pixel.Sprite
	herd    []Sheep
}

func (h *Herd) Update(ctx runtime.AppContext) {

}

func (h *Herd) Draw(ctx runtime.AppContext) {

}

func NewHerd(cfg HerdConfig) (*Herd, error) {
	h := &Herd{
		sprites: extractSprites(cfg.SheepPicture),
		herd:    make([]Sheep, 0),
	}

	h.herd = append(h.herd, NewSheep())

	return h, nil
}
