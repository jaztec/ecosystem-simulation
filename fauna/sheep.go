package fauna

import (
	"math/rand"

	"github.com/faiface/pixel"
)

type Sheep struct {
	object
	movable
}

func NewSheep() Sheep {
	return Sheep{
		object{
			gender:         Gender(rand.Intn(2)),
			food:           10,
			water:          10,
			reproduce:      20,
			vision:         100 + rand.Intn(20),
			speed:          5 + rand.Intn(5),
			attractiveness: Attractive(rand.Intn(3)),
		},
		movable{
			direction:   Direction(rand.Intn(4)),
			movingSpeed: 0,
			position:    pixel.ZV,
		},
	}
}
