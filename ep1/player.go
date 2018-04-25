package main

import (
	"fmt"
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSize  = 105
	playerSpeed = 240

	playerShotDelay = time.Millisecond * 300
)

// player represents the player character.
type player struct {
	tex  *sdl.Texture
	x, y float64

	lastShot time.Time
}

// newPlayer creates a new player object, loading all sprites.
func newPlayer(renderer *sdl.Renderer) (p player) {
	p.tex = textureFromBMP(renderer, "sprites/player.bmp")

	// Position the player at the bottom middle of the screen
	p.x = screenWidth / 2.0
	p.y = screenHeight - (playerSize / 2.0)

	return p
}

func (p *player) draw(renderer *sdl.Renderer) {
	// Convert points to refer to top left of sprite
	x := p.x - (playerSize / 2.0)
	y := p.y - (playerSize / 2.0)

	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: playerSize, H: playerSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: playerSize, H: playerSize})
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if p.x-playerSize/2.0 > 0 {
			p.x -= playerSpeed * deltaTime
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if p.x+playerSize/2.0 < screenWidth {
			p.x += playerSpeed * deltaTime
		}
	}

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(p.lastShot) >= playerShotDelay {
			p.shoot(p.x+25, p.y-20)
			p.shoot(p.x-25, p.y-20)

			p.lastShot = time.Now()
		}
	}
}

func (p *player) shoot(x, y float64) {
	if bul, ok := fromBulletRegistry(); ok {
		bul.x = x
		bul.y = y
		bul.angle = 270 * (math.Pi / 180.0)
		bul.active = true
	} else {
		fmt.Println("Warning: Ran out of bullets in registry")
	}
}

func (p *player) destroy() {
	p.tex.Destroy()
}
