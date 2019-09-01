package fauna

import "github.com/faiface/pixel"

type Gender uint8
type Attractive uint8
type Direction uint8

const (
	spriteEdgeH = 60
	spriteEdgeV = 75

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
	Food() int
	Water() int
	Reproduce() int
	Vision() int
	Speed() int
	Attractiveness() Attractive
}

type Movable interface {
	Position() pixel.Vec
	Direction() Direction
	MovingSpeed() int
}

type object struct {
	gender         Gender
	food           int
	water          int
	reproduce      int
	vision         int
	speed          int
	attractiveness Attractive
}

func (o *object) Gender() Gender {
	return o.gender
}

func (o *object) Food() int {
	return o.food
}

func (o *object) Water() int {
	return o.water
}

func (o *object) Reproduce() int {
	return o.reproduce
}

func (o *object) Vision() int {
	return o.vision
}

func (o *object) Speed() int {
	return o.speed
}

func (o *object) Attractiveness() Attractive {
	return o.attractiveness
}

type movable struct {
	direction   Direction
	movingSpeed int
	position    pixel.Vec
}

func (m *movable) Direction() Direction {
	return m.direction
}

func (m *movable) MovingSpeed() int {
	return m.movingSpeed
}

func (m *movable) Position() pixel.Vec {
	return m.position
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
