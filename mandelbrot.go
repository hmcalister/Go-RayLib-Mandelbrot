package main

import rl "github.com/gen2brain/raylib-go/raylib"

func drawMandelbrot(params drawParameters) *rl.Image {
	maxIterations := 100

	var pixelX int32
	var pixelY int32
	for pixelX = 0; pixelX < WINDOW_WIDTH; pixelX += 1 {
		for pixelY = 0; pixelY < WINDOW_HEIGHT; pixelY += 1 {
			pixelComplex := ComplexNumber{
				(centerX + float64(pixelX) - float64(WINDOW_WIDTH)/2) / zoom,
				(centerY + float64(pixelY) - float64(WINDOW_HEIGHT)/2) / zoom,
			}

			valueComplex := ComplexNumber{0, 0}
			iterationIndex := 0

			// As soon as the magnitude is above 4, the product will diverge
			for ComplexMagnitude(ComplexProduct(valueComplex, valueComplex)) < 4.0 && iterationIndex < maxIterations {
				valueComplex = ComplexAddition(ComplexProduct(valueComplex, valueComplex), pixelComplex)
				iterationIndex += 1
			}

			colorLERP := uint8(255 * iterationIndex / maxIterations)
			color := rl.NewColor(colorLERP, colorLERP, colorLERP, 255)
			rl.DrawPixel(pixelX, pixelY, color)

		}
	}

}
