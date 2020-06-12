package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"sync"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten"
)

var (
	boidImage *ebiten.Image
	boids     []*Boid
)

type Game struct{}

type Boid struct {
	posX float64
	posY float64
	velX float64
	velY float64
	accX float64
	accY float64
}

func CreateBoid() *Boid {
	b := new(Boid)
	b.posX = rand.Float64() * 1000.0
	b.posY = rand.Float64() * 1000.0
	b.velX = rand.Float64()*rand.Float64()*200.0 - 100.0
	b.velY = rand.Float64()*rand.Float64()*200.0 - 100.0

	return b
}

func (b *Boid) Update(step float64) {
	// Add acceleration
	b.velX += b.accX * step
	b.velY += b.accY * step

	// Calculate positions
	b.posX += b.velX * step
	b.posY += b.velY * step

	// Calculate out of bounds positions
	if b.posX < 0 {
		b.posX = 0
		b.velX = math.Abs(b.velX)
		b.accX = math.Abs(b.accX)
	} else if b.posX > 1000 {
		b.posX = 1000
		b.velX = -math.Abs(b.velX)
		b.accX = -math.Abs(b.accX)
	}

	if b.posY < 0 {
		b.posY = 0
		b.velY = math.Abs(b.velY)
		b.accY = math.Abs(b.accY)
	} else if b.posY > 1000 {
		b.posY = 1000
		b.velY = -math.Abs(b.velY)
		b.accY = -math.Abs(b.accY)
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	wg := new(sync.WaitGroup)
	wg.Add(25)

	delta := len(boids) / 25

	for i := 0; i < 25; i++ {
		go func(start int) {
			for index := start; index < start+delta; index++ {
				boids[index].Update(1.0 / 60.0)
			}
			wg.Done()
		}(i * delta)
	}
	wg.Wait()

	fmt.Println(ebiten.CurrentTPS())

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	for _, b := range boids {
		offset := ebiten.GeoM{}
		offset.Translate(b.posX, b.posY)

		ps := ebiten.DrawImageOptions{GeoM: offset}

		screen.DrawImage(boidImage, &ps)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1000, 1000
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

	boids = make([]*Boid, 1000)
	for i := range boids {
		boids[i] = CreateBoid()
	}

	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("Go Boids")

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
