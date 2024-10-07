package main

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var samColors []rl.Color

func GenRandomColors(n int) (colors []rl.Color) {
	for i := 0; i < n; i++ {
		r := rand.IntN(250)
		g := rand.IntN(250)
		b := rand.IntN(250)
		colors = append(colors, rl.NewColor(uint8(r), uint8(g), uint8(b), 250))
	}
	return
}

var CityScapeColors = []rl.Color{
	rl.DarkGray,   // road
	rl.Blue,       // sidewalk
	rl.DarkPurple, // building
	rl.DarkBrown,  // wall
	rl.Brown,      // fence
	rl.DarkGreen,  // pole
	rl.Yellow,     // traffic light
	rl.Red,        // traffic sign
	rl.Green,      // vegetation
	rl.DarkGreen,  // terrain
	rl.SkyBlue,    // sky
	rl.Yellow,     // person
	rl.Yellow,     // rider
	rl.Red,        // car
	rl.Red,        // truck
	rl.Red,        // bus
	rl.Red,        // train
	rl.Red,        // motorcycle
	rl.Red,        // bicycle
}

var CityScapeLabels = []string{
	"road",
	"sidewalk",
	"building",
	"wall",
	"fence",
	"pole",
	"traffic light",
	"traffic sign",
	"vegetation",
	"terrain",
	"sky",
	"person",
	"rider",
	"car",
	"truck",
	"bus",
	"train",
	"motorcycle",
	"bicycle",
}

var Keys = []string{
	"r", // road
	"w", // sidewalk
	"b", // building
	"m", // wall
	"f", // fence
	"y", // pole
	"a", // traffic light
	"i", // traffic sign
	"v", // vegetation
	"n", // terrain
	"s", // sky
	"p", // person
	"t", // rider
	"c", // car
	"g", // truck
	"h", // bus
	"j", // train
	"k", // motorcycle
	"l", // bicycle
}

var keyMap = map[string]int32{
	"q": rl.KeyQ,
	"w": rl.KeyW,
	"e": rl.KeyE,
	"r": rl.KeyR,
	"t": rl.KeyT,
	"y": rl.KeyY,
	"u": rl.KeyU,
	"i": rl.KeyI,
	"o": rl.KeyO,
	"p": rl.KeyP,
	"a": rl.KeyA,
	"s": rl.KeyS,
	"d": rl.KeyD,
	"f": rl.KeyF,
	"g": rl.KeyG,
	"h": rl.KeyH,
	"j": rl.KeyJ,
	"k": rl.KeyK,
	"l": rl.KeyL,
	"b": rl.KeyB,
	"c": rl.KeyC,
	"v": rl.KeyV,
}
