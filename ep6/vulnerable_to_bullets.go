package main

import "github.com/veandco/go-sdl2/sdl"

type vulnerableToBullets struct {
	container *element
}

func newVulnerableToBullets(container *element) *vulnerableToBullets {
	return &vulnerableToBullets{container: container}
}

func (vtb *vulnerableToBullets) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (vtb *vulnerableToBullets) onUpdate() error {
	return nil
}

func (vtb *vulnerableToBullets) onCollision(other *element) error {
	if other.tag == "bullet" {
		vtb.container.active = false
	}
	return nil
}
