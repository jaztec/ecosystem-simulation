package fauna

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
)

type Sheep struct {
	object
	movable
	dieable
}

func NewSheep() Sheep {
	return Sheep{
		object{
			age:            0.0,
			gender:         Gender(rand.Intn(2)),
			food:           60.0,
			water:          30.0,
			reproduce:      0.0,
			vision:         float64(100 + rand.Intn(20)),
			speed:          float64(1 + rand.Intn(2)),
			attractiveness: Attractive(rand.Intn(3)),
		},
		movable{
			direction:    Direction(rand.Intn(4)),
			movingSpeedX: 0,
			movingSpeedY: 0,
			position:     pixel.ZV,
			step:         0,
		},
		dieable{
			alive:  true,
			reason: 0,
			time:   time.Time{},
		},
	}
}
