package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type background struct {
	tex *sdl.Texture
}

func newBackground(renderer *sdl.Renderer) (b background) {
	b.tex = textureFromBMP(renderer, "sprites/background.bmp")

	return b
}

func (b *background) draw(renderer *sdl.Renderer) {
	renderer.Copy(b.tex,
		&sdl.Rect{X: 0, Y: 0, W: screenWidth, H: screenHeight},
		&sdl.Rect{X: 0, Y: 0, W: screenWidth, H: screenHeight})
}

func (b *background) destroy() {
	b.tex.Destroy()
}
