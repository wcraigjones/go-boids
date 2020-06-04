# Go Boids

A toy implementation of boids algorithm in Go.

## Stack
* https://github.com/hajimehoshi/ebiten - 2D Rendering
* https://github.com/fogleman/gg - Sprite Builder
* https://github.com/SolarLune/resolv - Simple 2D Physics

## Functionality
* Create and delete flocks of Boids ("bird-oid object")
    * Flocks should avoid each other
    * Flocks should try to stay together
    * Flocks should avoid objects
* Create and delete squares in the world
* Control a Boid of Prey that the flocks will avoid