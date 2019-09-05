package fauna

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
)

type Gender uint8
type Attractive uint8
type Direction uint8
type ReasonOfDeath uint8

// TODO Fix this function, just use []byte
func PrintStats(w io.Writer, animals []Dieable) {
	lines := make(map[ReasonOfDeath]int, 5)

	_, _ = fmt.Fprintf(w, "|%24s|\n", "Printing list of deaths")
	_, _ = fmt.Fprintf(w, "|%12s|%12d|\n", "no. animals", len(animals))
	_, _ = fmt.Fprintf(w, "|%12s|%12s|\n", "============", "============")
	_, _ = fmt.Fprintf(w, "|%12s|%12s|\n", "reason", "quantity")
	_, _ = fmt.Fprintf(w, "|%12s|%12s|\n", "============", "============")

	for _, a := range animals {
		lines[a.Reason()]++
	}

	for k, l := range lines {
		_, _ = fmt.Fprintf(w, "|%12s|%12d|\n", k, l)
	}
}

func (rof ReasonOfDeath) String() string {
	var s string
	switch rof {
	case Starvation:
		s = "starvation"
	case Thirst:
		s = "thirst"
	case Eaten:
		s = "eaten"
	case Age:
		s = "age"
	case Not:
		s = "not"
	}
	return s
}

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
	Animal
	Position() pixel.Vec
	SetPosition(pixel.Vec)
	Direction() Direction
	SetDirection(Direction)
	CalcDirection() Direction
	MovingSpeedX() float64
	MovingSpeedY() float64
	SetMovingSpeedX(float64)
	SetMovingSpeedY(float64)
	Step() uint8
	TimeIdle() float64
	SetTimeIdle(float64)
}

type Dieable interface {
	Animal
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
	direction    Direction
	movingSpeedX float64
	movingSpeedY float64
	position     pixel.Vec
	step         uint8
	timeIdle     float64
}

func (m *movable) Direction() Direction {
	return m.direction
}

func (m *movable) SetDirection(d Direction) {
	m.direction = d
}

func (m *movable) CalcDirection() Direction {
	var d Direction
	if m.movingSpeedX < 0 && math.Abs(m.movingSpeedY) <= math.Abs(m.movingSpeedX) {
		d = Left
	} else if m.movingSpeedX > 0 && math.Abs(m.movingSpeedY) <= math.Abs(m.movingSpeedX) {
		d = Right
	} else if m.movingSpeedY < 0 && math.Abs(m.movingSpeedX) <= math.Abs(m.movingSpeedY) {
		d = Up
	} else if m.movingSpeedY > 0 && math.Abs(m.movingSpeedX) <= math.Abs(m.movingSpeedY) {
		d = Down
	}
	return d
}

func (m *movable) MovingSpeedX() float64 {
	return m.movingSpeedX
}

func (m *movable) SetMovingSpeedX(ms float64) {
	m.movingSpeedX = ms
}

func (m *movable) MovingSpeedY() float64 {
	return m.movingSpeedY
}

func (m *movable) SetMovingSpeedY(ms float64) {
	m.movingSpeedY = ms
}

func (m *movable) Position() pixel.Vec {
	return m.position
}

func (m *movable) SetPosition(p pixel.Vec) {
	m.position = p
}

func (m *movable) Step() uint8 {
	return m.step
}

func (m *movable) TimeIdle() float64 {
	return m.timeIdle
}

func (m *movable) SetTimeIdle(ti float64) {
	m.timeIdle = ti
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

func randomMovement(m Movable, bounds pixel.Rect) Movable {
	m.SetTimeIdle(m.TimeIdle() + 1)
	changer := float64(75 + rand.Intn(50))
	if m.MovingSpeedX() == 0.0 && m.MovingSpeedY() == 0.0 {
		// optionally start moving again
		if m.TimeIdle() > changer {
			m.SetMovingSpeedX(m.Speed())
			m.SetMovingSpeedY(m.Speed())
			m.SetTimeIdle(0)
		}
	} else {
		// calculate and apply delta
		delta := pixel.ZV
		if m.TimeIdle() > changer {
			o := rand.Intn(3)
			switch o {
			case 0:
				// stop moving
				m.SetMovingSpeedX(0.0)
				m.SetMovingSpeedY(0.0)
				m.SetTimeIdle(0)
			case 1:
				// change direction
				var speeds [2]float64
				speeds[0] = m.Speed() * rand.Float64()
				speeds[1] = m.Speed() * rand.Float64()
				// possibly one of them is inverted
				if rand.Intn(1) == 0 {
					inv := rand.Intn(1)
					speeds[inv] = -speeds[inv]
				}

				m.SetMovingSpeedX(speeds[0])
				m.SetMovingSpeedY(speeds[1])
				m.SetTimeIdle(0)
			case 2:
				// do nothing
			}
		}

		delta.X = m.MovingSpeedX()
		delta.Y = m.MovingSpeedY()

		m.SetPosition(m.Position().Add(delta))
	}

	// check if we are still in bounds
	checkOutBounds(m, bounds)
	m.SetDirection(m.CalcDirection())

	return m
}

func checkOutBounds(m Movable, bounds pixel.Rect) bool {
	if m.Position().X < 0 {
		m.SetPosition(pixel.V(0.0, m.Position().Y))
		m.SetMovingSpeedX(-m.MovingSpeedX())
		return true
	}
	if m.Position().Y < 0 {
		m.SetPosition(pixel.V(m.Position().X, 0.0))
		m.SetMovingSpeedY(-m.MovingSpeedY())
		return true
	}
	if m.Position().X > bounds.Max.X {
		m.SetPosition(pixel.V(bounds.Max.X, m.Position().Y))
		m.SetMovingSpeedX(-m.MovingSpeedX())
		return true
	}
	if m.Position().Y > bounds.Max.Y {
		m.SetPosition(pixel.V(m.Position().X, bounds.Max.Y))
		m.SetMovingSpeedY(-m.MovingSpeedY())
		return true
	}
	return false
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
