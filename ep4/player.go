package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed = 0.05
	playerSize  = 105

	playerShotCooldown = time.Millisecond * 250
)

func newPlayer(renderer *sdl.Renderer) *element {
	player := &element{}

	player.position = vector{
		x: screenWidth / 2.0,
		y: screenHeight - playerSize/2.0,
	}

	player.active = true

	sr := newSpriteRenderer(player, renderer, "sprites/player.bmp")
	player.addComponent(sr)

	mover := newKeyboardMover(player, playerSpeed)
	player.addComponent(mover)

	shooter := newKeyboardShooter(player, playerShotCooldown)
	player.addComponent(shooter)

	return player
}
