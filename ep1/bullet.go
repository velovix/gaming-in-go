package main

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 32
	bulletSpeed = 600
)

type bullet struct {
	tex         *sdl.Texture
	x, y, angle float64
	active      bool
}

var bulletTexture *sdl.Texture

func newBullet(renderer *sdl.Renderer) (b bullet, err error) {
	if bulletTexture == nil {
		bulletTexture = textureFromBMP(renderer, "sprites/player_bullet.bmp")
	}
	b.tex = bulletTexture

	b.active = false

	return b, nil
}

func (bul *bullet) draw(renderer *sdl.Renderer) {
	if !bul.active {
		return
	}

	// Convert points to refer to top left of sprite
	x := bul.x - (bulletSize / 2.0)
	y := bul.y - (bulletSize / 2.0)

	renderer.Copy(bul.tex,
		&sdl.Rect{X: 0, Y: 0, W: bulletSize, H: bulletSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: bulletSize, H: bulletSize})
}

func (bul *bullet) update() {
	bul.x += bulletSpeed * math.Cos(bul.angle) * deltaTime
	bul.y += bulletSpeed * math.Sin(bul.angle) * deltaTime

	// Check if the bullet fell out of screen
	if bul.x > screenWidth || bul.x < 0 || bul.y > screenHeight || bul.y < 0 {
		bul.active = false
	}
}

var bulletRegistry []*bullet

func initBulletRegistry(renderer *sdl.Renderer) {
	bulletRegistry = make([]*bullet, 30)

	for i := range bulletRegistry {
		bul, err := newBullet(renderer)
		if err != nil {
			panic(fmt.Sprintf("initializing a bullet for registry: %v", err))
		}
		bulletRegistry[i] = &bul
	}
}

func fromBulletRegistry() (*bullet, bool) {
	for _, bul := range bulletRegistry {
		if !bul.active {
			return bul, true
		}
	}
	return nil, false
}
