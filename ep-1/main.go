package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		"Gaming in Go Episode 2",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		500, 800,
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

	quitting := false
	for !quitting {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				quitting = true
			}
		}

		renderer.Clear()
		renderer.Present()
	}
}
