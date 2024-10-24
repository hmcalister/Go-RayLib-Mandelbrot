package main

import rl "github.com/gen2brain/raylib-go/raylib"

func createMandelbrotTexture(params drawParameters) rl.Texture2D {
	maxIterations := 100

	image := rl.GenImageColor(int(WINDOW_WIDTH), int(WINDOW_HEIGHT), rl.Black)

	var pixelX int32
	var pixelY int32
	for pixelX = 0; pixelX < WINDOW_WIDTH; pixelX += 1 {
		for pixelY = 0; pixelY < WINDOW_HEIGHT; pixelY += 1 {
			pixelComplex := ComplexNumber{
				params.centerX + (float64(pixelX)-float64(WINDOW_WIDTH)/2)/params.zoom,
				params.centerY + (float64(pixelY)-float64(WINDOW_HEIGHT)/2)/params.zoom,
			}

			valueComplex := ComplexNumber{0, 0}
			iterationIndex := 0

			// As soon as the magnitude is above 4, the product will diverge
			for valueComplex.Product(valueComplex).Magnitude() < 4.0 && iterationIndex < maxIterations {
				valueComplex = valueComplex.Product(valueComplex).Add(pixelComplex)
				iterationIndex += 1
			}

			colorLERP := uint8(255 * iterationIndex / maxIterations)
			color := rl.NewColor(colorLERP, colorLERP, colorLERP, 255)
			rl.ImageDrawPixel(image, pixelX, pixelY, color)
		}
	}

	texture := rl.LoadTextureFromImage(image)
	return texture
}
