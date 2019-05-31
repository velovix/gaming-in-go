package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct {
	container     *element
	tex           *sdl.Texture
	width, height int
}

func newSpriteRenderer(container *element, renderer *sdl.Renderer, filename string) *spriteRenderer {
	sr := &spriteRenderer{}

	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	defer img.Free()
	sr.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("creating texture from %v: %v", filename, err))
	}

	_, _, width, height, err := sr.tex.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture: %v", err))
	}
	sr.width = int(width)
	sr.height = int(height)

	sr.container = container

	return sr
}

func (sr *spriteRenderer) start() {
}

func (sr *spriteRenderer) onUpdate() error {
	return nil
}

func (sr *spriteRenderer) onDraw(renderer *sdl.Renderer) error {
	_, _, width, height, err := sr.tex.Query()
	if err != nil {
		return fmt.Errorf("querying texture: %v", err)
	}

	// Convert coordinates to the top left of the sprite
	pos := sr.container.position
	pos.x -= float64(width) / 2.0
	pos.y -= float64(height) / 2.0

	renderer.CopyEx(
		sr.tex,
		&sdl.Rect{X: 0, Y: 0, W: width, H: height},
		&sdl.Rect{X: int32(pos.x), Y: int32(pos.y), W: width, H: height},
		sr.container.rotation,
		&sdl.Point{X: width / 2, Y: height / 2},
		sdl.FLIP_NONE)

	return nil
}

func (sr *spriteRenderer) onCollision(other *element) error {
	return nil
}
