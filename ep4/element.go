package main

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type vector struct {
	x, y float64
}

type component interface {
	onUpdate() error
	onDraw(renderer *sdl.Renderer) error
}

type element struct {
	position   vector
	rotation   float64
	active     bool
	components []component
}

func (elem *element) draw(renderer *sdl.Renderer) error {
	for _, comp := range elem.components {
		err := comp.onDraw(renderer)
		if err != nil {
			return err
		}
	}

	return nil
}

func (elem *element) update() error {
	for _, comp := range elem.components {
		err := comp.onUpdate()
		if err != nil {
			return err
		}
	}

	return nil
}

func (elem *element) addComponent(new component) {
	for _, existing := range elem.components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf(
				"attempt to add new component with existing type %v",
				reflect.TypeOf(new)))
		}
	}
	elem.components = append(elem.components, new)
}

func (elem *element) getComponent(withType component) component {
	typ := reflect.TypeOf(withType)
	for _, comp := range elem.components {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}

	panic(fmt.Sprintf("no component with type %v", reflect.TypeOf(withType)))
}

var elements []*element
