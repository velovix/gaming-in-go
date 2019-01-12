package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 32
	bulletSpeed = 0.15
)

func newBullet(renderer *sdl.Renderer) *element {
	bullet := &element{}

	sr := newSpriteRenderer(bullet, renderer, "sprites/player_bullet.bmp")
	bullet.addComponent(sr)

	mover := newBulletMover(bullet)
	bullet.addComponent(mover)

	col := circle{
		center: bullet.position,
		radius: 8,
	}
	bullet.collisions = append(bullet.collisions, col)

	bullet.tag = "bullet"

	return bullet
}

var bulletPool []*element

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bul := newBullet(renderer)
		bulletPool = append(bulletPool, bul)
		elements = append(elements, bul)
	}
}

func bulletFromPool() (*element, bool) {
	for _, bul := range bulletPool {
		if !bul.active {
			return bul, true
		}
	}

	return nil, false
}
