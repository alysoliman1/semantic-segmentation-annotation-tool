package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextureLoader struct {
	currentName string
	texture     rl.Texture2D
}

func (s *TextureLoader) SetCurrentImage(name string) {
	if name == s.currentName {
		return
	}
	if s.currentName != "" {
		rl.UnloadTexture(s.texture)
	}
	s.currentName = name
	s.texture = rl.LoadTexture(fmt.Sprintf("images/%s.png", name))
}
