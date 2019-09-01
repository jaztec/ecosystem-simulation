package world

import (
	"github.com/faiface/pixel"
)

type Camera interface {
	Pos() pixel.Vec
	Speed() float64
	SetPosX(float64)
	SetPosY(float64)
}

type camera struct {
	pos   pixel.Vec
	speed float64
}

func (c *camera) Pos() pixel.Vec {
	return c.pos
}

func (c *camera) SetPosX(x float64) {
	c.pos.X = x
}

func (c *camera) SetPosY(y float64) {
	c.pos.Y = y
}

func (c *camera) Speed() float64 {
	return c.speed
}

func NewCamera() Camera {
	return &camera{pos: pixel.ZV, speed: 1000.0}
}
