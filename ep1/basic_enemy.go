package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const enemySize = 105

type basicEnemy struct {
	tex  *sdl.Texture
	x, y float64
}

var basicEnemyTexture *sdl.Texture

func newBasicEnemy(renderer *sdl.Renderer, x, y float64) (e basicEnemy) {
	if basicEnemyTexture == nil {
		basicEnemyTexture = textureFromBMP(renderer, "sprites/basic_enemy.bmp")
	}
	e.tex = basicEnemyTexture

	e.x = x
	e.y = y

	return e
}

func (e *basicEnemy) draw(renderer *sdl.Renderer) {
	// Convert points to refer to top left of sprite
	x := e.x - (enemySize / 2.0)
	y := e.y - (enemySize / 2.0)

	renderer.CopyEx(
		e.tex,
		&sdl.Rect{X: 0, Y: 0, W: enemySize, H: enemySize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: enemySize, H: enemySize},
		180,
		&sdl.Point{X: enemySize / 2, Y: enemySize / 2},
		sdl.FLIP_NONE)
}

func (e *basicEnemy) destroy() {
	e.tex.Destroy()
}
