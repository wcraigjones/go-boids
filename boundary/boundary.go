package boundary

import (
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

func AddBoundary(scene *core.Node) {
	// Create the boundary box, 2x2x2 centered around origin
	planeMat := material.NewStandard(math32.NewColor("blue"))
	planeGeom := geometry.NewPlane(2, 2)

	// X Offsets
	x1PlaneMesh := graphic.NewMesh(planeGeom, planeMat)
	x1PlaneMesh.SetPosition(0,1,0)
	x1PlaneMesh.SetRotation(math32.Pi / 2, 0, 0)
	scene.Add(x1PlaneMesh)

	x2PlaneMesh := graphic.NewMesh(planeGeom, planeMat)
	x2PlaneMesh.SetPosition(0,-1,0)
	x2PlaneMesh.SetRotation(math32.Pi / 2, 0, math32.Pi)
	scene.Add(x2PlaneMesh)

	// Y Offsets
	y1PlaneMesh := graphic.NewMesh(planeGeom, planeMat)
	y1PlaneMesh.SetPosition(-1,0,0)
	y1PlaneMesh.SetRotation(0, math32.Pi / 2, 0)
	scene.Add(y1PlaneMesh)

	y2PlaneMesh := graphic.NewMesh(planeGeom, planeMat)
	y2PlaneMesh.SetPosition(1,0,0)
	y2PlaneMesh.SetRotation(math32.Pi, math32.Pi / 2, 0)
	scene.Add(y2PlaneMesh)

	// Z Offsets
	z1PlaneMesh := graphic.NewMesh(planeGeom, planeMat)
	z1PlaneMesh.SetPosition(0,0,-1)
	scene.Add(z1PlaneMesh)

	z2PlaneMesh := graphic.NewMesh(planeGeom, planeMat)
	z2PlaneMesh.SetPosition(0,0,1)
	z2PlaneMesh.SetRotation(math32.Pi, 0, 0)
	scene.Add(z2PlaneMesh)

}
