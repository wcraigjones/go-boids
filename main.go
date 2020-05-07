package main

import (
	"boids/boundary"
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/engine/window"
	"math/rand"
	"time"
)

func main() {

	// Create application and scene
	a := app.App()
	scene := core.NewNode()

	// Set the scene to be managed by the gui manager
	gui.Manager().Set(scene)

	// Create perspective camera
	cam := camera.New(1)
	cam.SetPosition(0,0, 3)
	scene.Add(cam)

	// Set up orbit control for the camera
	orbitCamera := camera.NewOrbitControl(cam)
	orbitCamera.MaxDistance = 5
	orbitCamera.MinDistance = 2.5

	// Set up callback to update viewport and camera aspect ratio when the window is resized
	onResize := func(evname string, ev interface{}) {
		// Get framebuffer size and update viewport accordingly
		width, height := a.GetSize()
		a.Gls().Viewport(0, 0, int32(width), int32(height))
		// Update the camera's aspect ratio
		cam.SetAspect(float32(width) / float32(height))
	}
	a.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	// Create and add an axis helper to the scene
	scene.Add(helper.NewAxes(0.5))

	// Add the boundaries
	boundary.AddBoundary(scene)

	// Create and add lights to the scene
	scene.Add(light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.8))
	pointLight := light.NewPoint(&math32.Color{1, 1, 1}, 5.0)
	pointLight.SetPosition(1, 0, 2)
	scene.Add(pointLight)

	// Add a ball

	ballMat := material.NewStandard(math32.NewColor("red"))
	ballGeom := geometry.NewSphere(0.01, 10, 10)
	for i := 0; i < 10; i++ {
		ballMesh := graphic.NewMesh(ballGeom, ballMat)
		go func() {
			position := ballMesh.Position()
			lastTime := a.RunTime()
			for scene.Visible() {
				// Update Delta Time
				currentTime := a.RunTime()
				deltaTime := float32((currentTime - lastTime).Seconds())
				lastTime = currentTime

				// Get Drift
				deltaVector := math32.Vector3{
					X: (rand.Float32() - 0.5) * 10 * deltaTime,
					Y: (rand.Float32() - 0.5) * 10 * deltaTime,
					Z: (rand.Float32() - 0.5) * 10 * deltaTime,
				}
				newPosition := position.Add(&deltaVector).ClampScalar(-1, 1)
				ballMesh.SetPositionVec(newPosition)

				// Pause for 10ms
				time.Sleep(time.Millisecond * 10 - time.Second * time.Duration(deltaTime))
			}
		}()
		scene.Add(ballMesh)
	}

	// Set background color to black
	a.Gls().ClearColor(0.0,0.0,0.0, 1.0)

	// Run the application
	a.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		a.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		renderer.Render(scene, cam)
	})
}