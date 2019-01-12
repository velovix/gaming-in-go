package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const basicEnemySize = 105

func newBasicEnemy(renderer *sdl.Renderer, position vector) *element {
	basicEnemy := &element{}

	basicEnemy.position = position
	basicEnemy.rotation = 180

	sr := newSpriteRenderer(basicEnemy, renderer, "sprites/basic_enemy.bmp")
	basicEnemy.addComponent(sr)

	basicEnemy.active = true

	return basicEnemy
}
