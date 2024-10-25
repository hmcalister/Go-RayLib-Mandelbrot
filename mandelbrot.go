package main

import (
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// A worker function to generate one row of the mandelbrot image.
// This goroutine is partitioned on the *rows* of the image,
// so no two goroutines touch the same row at the same time.
func mandelbrotWorker(image *rl.Image, params drawParameters, rowChannel chan int32) {
	var pixelX int32
	var pixelY int32

	for pixelX = range rowChannel {
		for pixelY = 0; pixelY < WINDOW_HEIGHT; pixelY += 1 {
			pixelComplex := params.convertPixelToComplex(pixelX, pixelY)

			valueComplex := ComplexNumber{0, 0}
			iterationIndex := 0

			// As soon as the magnitude is above 4, the product will diverge
			for valueComplex.Product(valueComplex).Magnitude() < 4.0 && iterationIndex < params.maxIterations {
				valueComplex = valueComplex.Product(valueComplex).Add(pixelComplex)
				iterationIndex += 1
			}

			colorLERP := uint8(255 * iterationIndex / params.maxIterations)
			color := rl.NewColor(colorLERP, colorLERP, colorLERP, 255)
			rl.ImageDrawPixel(image, pixelX, pixelY, color)
		}
	}
}

func createMandelbrotTexture(params drawParameters, numWorkerGoroutines int) rl.Texture2D {
	image := rl.GenImageColor(int(WINDOW_WIDTH), int(WINDOW_HEIGHT), rl.Black)

	var pixelX int32
	var workersFinishedWaitGroup sync.WaitGroup
	rowChannel := make(chan int32)

	for range numWorkerGoroutines {
		workersFinishedWaitGroup.Add(1)
		go func() {
			defer workersFinishedWaitGroup.Done()
			mandelbrotWorker(image, params, rowChannel)
		}()
	}

	for pixelX = 0; pixelX < WINDOW_WIDTH; pixelX += 1 {
		rowChannel <- pixelX
	}
	close(rowChannel)
	workersFinishedWaitGroup.Wait()

	texture := rl.LoadTextureFromImage(image)
	return texture
}
