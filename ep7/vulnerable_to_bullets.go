package main

import "github.com/veandco/go-sdl2/sdl"

type vulnerableToBullets struct {
	container *element
	animator  *animator
}

func newVulnerableToBullets(container *element) *vulnerableToBullets {
	return &vulnerableToBullets{
		container: container,
		animator:  container.getComponent(&animator{}).(*animator)}
}

func (vtb *vulnerableToBullets) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (vtb *vulnerableToBullets) onUpdate() error {
	if vtb.animator.finished && vtb.animator.current == "destroy" {
		vtb.container.active = false
	}

	return nil
}

func (vtb *vulnerableToBullets) onCollision(other *element) error {
	if other.tag == "bullet" {
		vtb.animator.setSequence("destroy")
	}
	return nil
}
