package visualisation

import (
	"1msnakes/game"
	"1msnakes/vectors"
	"image"
	"image/color"
	"image/gif"
	"os"
)

func Visualize() {
	palette := []color.Color{color.White, color.Black}

	numFrames := 40
	size := 100
	anim := gif.GIF{LoopCount: numFrames}

	snake := game.CreateSnake([]*vectors.Vector{
		{X: 40, Y: 50},
		{X: 30, Y: 50},
	})

	dirs := []vectors.Directions{
		vectors.VecS,
		vectors.VecW,
		vectors.VecN,
		vectors.VecE,
		vectors.VecW,
		vectors.VecN,
		vectors.VecE,
		vectors.VecN,
		vectors.VecE,
		vectors.VecW,
		vectors.VecN,
	}

	for i := 0; i < len(dirs)*len(snake.GetPixels()); i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		snakePixels := snake.GetPixels()
		for _, px := range snakePixels {
			img.SetColorIndex(int(px.X), int(px.Y), 1)
		}

		dir := (i / len(snakePixels)) % len(dirs)
		snake.Move(dirs[dir])

		anim.Delay = append(anim.Delay, 8)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(os.Stdout, &anim)
}
