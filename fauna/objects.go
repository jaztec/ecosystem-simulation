package fauna

import (
	"time"

	"github.com/faiface/pixel"
)

type Gender uint8
type Attractive uint8
type Direction uint8
type ReasonOfDeath uint8

const (
	spriteEdgeH = 60
	spriteEdgeV = 75

	StoppedIndex = 4
	MovingRight  = 3
	MovingLeft   = 2
	MovingDown   = 1
	MovingUp     = 0

	Not        ReasonOfDeath = 0
	Starvation ReasonOfDeath = 1
	Thirst     ReasonOfDeath = 2
	Eaten      ReasonOfDeath = 3
	Age        ReasonOfDeath = 4

	Male   Gender = 0x01
	Female Gender = 0x02

	Ugly        Attractive = 0x01
	Normal      Attractive = 0x02
	GoodLooking Attractive = 0x03

	Up    Direction = 0x01
	Down  Direction = 0x02
	Left  Direction = 0x03
	Right Direction = 0x04
)

type Animal interface {
	Gender() Gender
	Food() float64
	Water() float64
	Reproduce() float64
	Vision() float64
	Speed() float64
	Attractiveness() Attractive
	Age() float64
}

type Movable interface {
	Position() pixel.Vec
	Direction() Direction
	MovingSpeed() float64
	Step() uint8
}

type Dieable interface {
	Alive() bool
	Reason() ReasonOfDeath
	Time() time.Time
}

type object struct {
	gender         Gender
	age            float64
	food           float64
	water          float64
	reproduce      float64
	vision         float64
	speed          float64
	attractiveness Attractive
}

func (o *object) Gender() Gender {
	return o.gender
}

func (o *object) Age() float64 {
	return o.age
}

func (o *object) Food() float64 {
	return o.food
}

func (o *object) Water() float64 {
	return o.water
}

func (o *object) Reproduce() float64 {
	return o.reproduce
}

func (o *object) Vision() float64 {
	return o.vision
}

func (o *object) Speed() float64 {
	return o.speed
}

func (o *object) Attractiveness() Attractive {
	return o.attractiveness
}

type movable struct {
	direction   Direction
	movingSpeed float64
	position    pixel.Vec
	step        uint8
}

func (m *movable) Direction() Direction {
	return m.direction
}

func (m *movable) MovingSpeed() float64 {
	return m.movingSpeed
}

func (m *movable) Position() pixel.Vec {
	return m.position
}

func (m *movable) Step() uint8 {
	return m.step
}

type dieable struct {
	alive  bool
	reason ReasonOfDeath
	time   time.Time
}

func (d *dieable) Alive() bool {
	return d.alive
}

func (d *dieable) Reason() ReasonOfDeath {
	return d.reason
}

func (d *dieable) Time() time.Time {
	return d.time
}

func extractSprites(p pixel.Picture) (set [5][4]*pixel.Sprite) {
	for h := 0; h < 5; h++ {
		for v := 0; v < 4; v++ {
			set[h][v] = pixel.NewSprite(p, pixel.R(
				float64(spriteEdgeV*v), float64(spriteEdgeH*h),
				float64(spriteEdgeV+(spriteEdgeV*v)), float64(spriteEdgeH+(spriteEdgeH*h))))
		}
	}
	return
}
