package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	container *element
	speed     float64

	sr *spriteRenderer
}

func newKeyboardMover(container *element, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
		sr:        container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (mover *keyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()

	cont := mover.container

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if cont.position.x-(mover.sr.width/2.0) > 0 {
			cont.position.x -= mover.speed
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if cont.position.x+(mover.sr.height/2.0) < screenWidth {
			cont.position.x += mover.speed
		}
	}

	return nil
}

func (mover *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

type keyboardShooter struct {
	container *element
	cooldown  time.Duration
	lastShot  time.Time
}

func newKeyboardShooter(container *element, cooldown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		container: container,
		cooldown:  cooldown,
	}
}

func (mover *keyboardShooter) onUpdate() error {
	keys := sdl.GetKeyboardState()

	pos := mover.container.position

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(mover.lastShot) >= mover.cooldown {
			mover.shoot(pos.x+25, pos.y-20)
			mover.shoot(pos.x-25, pos.y-20)

			mover.lastShot = time.Now()
		}
	}

	return nil
}

func (mover *keyboardShooter) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *keyboardShooter) shoot(x, y float64) {
	if bul, ok := bulletFromPool(); ok {
		bul.active = true
		bul.position.x = x
		bul.position.y = y
		bul.rotation = 270 * (math.Pi / 180)
	}
}
