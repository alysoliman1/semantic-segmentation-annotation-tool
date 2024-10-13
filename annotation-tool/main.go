package main

import (
	"fmt"
	"os"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	textureW int32 = 6080
	textureH int32 = 3040
	maskW    int32 = 608
	maskH    int32 = 304

	brushOn              bool
	brushSize            = 10
	lastSelectedCategory int

	srcRect = rl.NewRectangle(
		0,
		0,
		float32(textureW),
		float32(textureH),
	)
	dstRect = rl.NewRectangle(
		float32(maskW),
		0,
		float32(maskW),
		float32(maskH),
	)
	dstRect2 = rl.NewRectangle(
		float32(maskW),
		float32(maskH),
		float32(maskW),
		float32(maskH),
	)

	currentImageIndex int
	images            = []string{
		"ATTA1",
		"ATTA2",
		"ATTA4",
		"ATTA5",
		"AhmedOrabi1",
		"Banat1",
		"Banat2",
		"Banat3",
		"FO1",
		"FO2",
		"FO3",
		"FO5",
		"FO6",
		"FO9",
		"FR0",
		"FR1",
		"FR2",
		"FR5",
		"FR6",
		"FR7",
		"IshaakKadeem1",
		"KaedGohar1",
		"KaedGohar2",
		"KaedGohar4",
		"MANSH1",
		"MANSH3",
		"MANSH4",
		"MANSH5",
		"MANSH6",
		"MANSH7",
		"MANSH8",
		"MANSH9",
		"MamdouhAzmy1",
		"MamdouhAzmy2",
		"MamdouhAzmy3",
		"MamdouhAzmy4",
		"ND1",
		"ND2",
		"ND3",
		"ND4",
		"NS1",
		"NS2",
		"NS3",
		"NS4",
		"OL1",
		"OL2",
		"OL3",
		"OL5",
		"OL6",
		"OL7",
		"OL8",
		"SA1",
		"SA2",
		"SA3",
		"SA5",
		"SA6",
		"SA7",
		"SA8",
		"SA9-2",
		"SA9-5",
		"SA9-6",
		"SA9",
		"SDin1",
		"SDin2",
		"SDin3",
		"SF1",
		"SF2",
		"SF3",
		"SF4",
		"SF5",
		"SF7",
		"SF8",
		"SF9",
		"SM2-2",
		"SM2",
		"SM3",
		"SM4",
		"SM5",
		"SM6",
		"SM7",
		"SM8",
		"SM9-1",
		"SM9-2",
		"SM9-3",
		"SM9-4",
		"SM9",
		"SalahSal1",
		"SalahSal4",
		"SalahSal5",
		"Sero1",
		"Sero2",
		"SidiM2",
		"SidiM3",
		"SidiM4",
		"Sm8-2",
	}

	textureLoader TextureLoader
	samLoader     = NewMasksLoader("sam-labels")
	semLoader     = NewMasksLoader("sem-labels")
	selector      = NewSelector("selectors")

	keySheet []string

	highlightedSAMCategory int
)

func init() {
	rl.InitWindow(maskW*2+400, maskH*2, "annotate")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	// ls images is piped into program args.
	if len(os.Args) > 1 && len(images) == 0 {
		for _, arg := range os.Args[1:] {
			image := strings.TrimSuffix(arg, ".jpg")
			images = append(images, image)
		}
	}

	samColors = GenRandomColors(500)
	var currentImage = images[currentImageIndex]
	textureLoader.SetCurrentImage(currentImage)
	samLoader.SetCurrentImage(currentImage)
	semLoader.SetCurrentImage(currentImage)
	selector.Update(currentImage, samLoader, semLoader)

	for i := 0; i < 19; i++ {
		keySheet = append(keySheet, fmt.Sprintf("%s: %s (%d)", Keys[i], CityScapeLabels[i], i))
	}
}

func main() {
	defer rl.CloseWindow()
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		var imageChanged bool
		if rl.IsKeyPressed(rl.KeyRight) {
			imageChanged = true
			currentImageIndex = (currentImageIndex + 1) %
				len(images)
		}
		if rl.IsKeyPressed(rl.KeyLeft) {
			imageChanged = true
			currentImageIndex = (currentImageIndex + len(images) - 1) %
				len(images)
		}

		if rl.IsKeyPressed(rl.KeyX) {
			brushOn = !brushOn
		}

		if imageChanged {
			var currentImage = images[currentImageIndex]

			textureLoader.SetCurrentImage(currentImage)
			samLoader.SetCurrentImage(currentImage)
			semLoader.SetCurrentImage(currentImage)
			selector.Update(currentImage, samLoader, semLoader)
		}

		// Render texture.
		if textureLoader.currentName != "" {
			rl.DrawTexturePro(
				textureLoader.texture,
				srcRect,
				dstRect,
				rl.Vector2{},
				0,
				rl.White,
			)
			rl.DrawTexturePro(
				textureLoader.texture,
				srcRect,
				dstRect2,
				rl.Vector2{},
				0,
				rl.White,
			)
		}

		// Render SAM masks.
		for i, row := range samLoader.LabelMatrix {
			for j, v := range row {

				var (
					x = rl.GetMouseX()
					y = rl.GetMouseY()
				)

				if x == int32(j) && y == int32(i) {
					highlightedSAMCategory = v
					if !brushOn {
						lastSelectedCategory = selector.Select(j, i, v)
					}
				}

				color := samColors[v]
				if highlightedSAMCategory == v {
					color = rl.Black
				}

				rl.DrawPixel(
					int32(j),
					int32(i),
					color,
				)
				color = CityScapeColors[selector.Select(j, i, v)]
				color.A = 200
				rl.DrawPixel(
					maskW+int32(j),
					maskH+int32(i),
					color,
				)
			}
		}

		// Render semantic segmentation masks.
		for i, row := range semLoader.LabelMatrix {
			for j, v := range row {
				rl.DrawPixel(
					int32(j),
					maskH+int32(i),
					CityScapeColors[v],
				)
			}
		}

		for i, key := range Keys {
			if rl.IsKeyDown(keyMap[key]) {
				if brushOn {
					lastSelectedCategory = i
				} else {
					selector.SetOverride(highlightedSAMCategory, i)
				}
			}
		}

		rl.DrawText(fmt.Sprintf("%d SAM masks", len(samLoader.Labels)),
			2*maskW+10, 10, 25, rl.Black)
		rl.DrawText(fmt.Sprintf("category: %s (%d)",
			CityScapeLabels[lastSelectedCategory],
			lastSelectedCategory,
		),
			2*maskW+10, 35, 25, rl.Black)

		for i := 0; i < 19; i++ {
			rl.DrawText(keySheet[i], 2*maskW+40, 65+25*int32(i), 22, rl.Black)
			rl.DrawRectangle(2*maskW+10, 65+25*int32(i), 20, 20, CityScapeColors[i])
		}

		if brushOn {
			var (
				x = rl.GetMouseX() - maskW
				y = rl.GetMouseY() - maskH
			)
			if x >= 0 && x <= maskW && y >= 0 && y <= maskH {
				var (
					rx = x + maskW - int32(brushSize)/2
					ry = y + maskH - int32(brushSize)/2
				)
				rl.DrawRectangle(rx, ry, int32(brushSize), int32(brushSize),
					CityScapeColors[lastSelectedCategory])
				if rl.IsKeyPressed(rl.KeyUp) && brushSize < 100 {
					brushSize += 10
				}
				if rl.IsKeyPressed(rl.KeyDown) && brushSize > 10 {
					brushSize -= 10
				}
				if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
					selector.AddBox(lastSelectedCategory, Box{
						X: int(x),
						Y: int(y),
						S: brushSize,
					})
				}
			}
		}
		rl.EndDrawing()
	}
}
