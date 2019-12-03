package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type bulletMover struct {
	container *element
	speed     float64
}

func newBulletMover(container *element) *bulletMover {
	return &bulletMover{container: container}
}

func (mover *bulletMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *bulletMover) onUpdate() error {
	c := mover.container

	c.position.x += bulletSpeed * math.Cos(c.rotation) * delta
	c.position.y += bulletSpeed * math.Sin(c.rotation) * delta

	if c.position.x > screenWidth || c.position.x < 0 ||
		c.position.y > screenHeight || c.position.y < 0 {
		c.active = false
	}

	c.collisions[0].center = c.position

	return nil
}

func (mover *bulletMover) onCollision(other *element) error {
	mover.container.active = false
	return nil
}
