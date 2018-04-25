package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

var deltaTime = 1.0 / 60.0

func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	// Initialize the sprite as a texture
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Sprintf("loading %v: %v", filename, err))
	}
	defer img.Free()

	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Sprintf("creating texture for %v: %v", filename, err))
	}

	return tex
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		"Gaming in Go Episode 2",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("creating window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("creating renderer:", err)
		return
	}
	defer renderer.Destroy()

	player := newPlayer(renderer)
	defer player.destroy()

	bg := newBackground(renderer)
	defer bg.destroy()

	const (
		waveWidth  = 5
		waveHeight = 3
	)
	var enemyWave []basicEnemy
	for i := 0; i < waveWidth; i++ {
		for j := 0; j < waveHeight; j++ {
			x := (float64(i)/waveWidth)*screenWidth + enemySize/2
			y := float64(enemySize*j) + enemySize/2

			e := newBasicEnemy(renderer, x, y)
			defer e.destroy()
			enemyWave = append(enemyWave, e)
		}
	}

	initBulletRegistry(renderer)

	quitting := false
	for !quitting {
		startTime := time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				quitting = true
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		bg.draw(renderer)

		player.update()
		player.draw(renderer)

		for _, enemy := range enemyWave {
			enemy.draw(renderer)
		}

		for _, bul := range bulletRegistry {
			bul.draw(renderer)
			bul.update()
		}

		renderer.Present()

		deltaTime = time.Since(startTime).Seconds()
	}
}
