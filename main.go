package main

import (
	"image/color"
	"log"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten"
)

var (
	boidImage *ebiten.Image
)

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	offset := ebiten.GeoM{}
	offset.Translate(156, 116)

	ps := ebiten.DrawImageOptions{GeoM: offset}

	screen.DrawImage(boidImage, &ps)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func generateImages() {
	dc := gg.NewContext(16, 16)
	dc.DrawCircle(8, 8, 8)
	dc.SetColor(color.White)
	dc.Fill()
	boidImage, _ = ebiten.NewImageFromImage(dc.Image(), ebiten.FilterDefault)
}

func main() {
	// Generate our images
	generateImages()

	game := &Game{}
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Go Boids")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
