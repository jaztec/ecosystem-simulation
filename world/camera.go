package world

import "github.com/faiface/pixel"

type Camera interface {
	Pos() *pixel.Vec
	Speed() float64
}

type camera struct {
	pos   pixel.Vec
	speed float64
}

func (c *camera) Pos() *pixel.Vec {
	return &c.pos
}

func (c *camera) Speed() float64 {
	return c.speed
}

func NewCamera() Camera {
	return &camera{pos: pixel.ZV}
}
