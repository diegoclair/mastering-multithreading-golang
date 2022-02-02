package main

import (
	"fmt"
	"image/color"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()

}

const (
	screenWidth, screenHeight = 640, 360
	boidCount                 = 500
	viewRadius                = 13
	adjRate                   = 0.015
)

var (
	green = color.RGBA{10, 255, 50, 255}
	boids [boidCount]*Boid
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

// func (g *Game) Draw(screen *ebiten.Image) {
// 	for _, boid := range boids {
// 		screen.Set(int(boid.Position.X+1), int(boid.Position.Y), green)
// 		screen.Set(int(boid.Position.X-1), int(boid.Position.Y), green)
// 		screen.Set(int(boid.Position.X), int(boid.Position.Y-1), green)
// 		screen.Set(int(boid.Position.X), int(boid.Position.Y+1), green)
// 	}

// }

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func main() {

	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()
	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	for !window.ShouldClose() {
		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}
	//	glfw.WindowHint(glfw.Visible, glfw.False)

	fmt.Println("OPA START")

	// for i := 0; i < boidCount; i++ {
	// 	createBoid(i)
	// }
	// ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	// ebiten.SetWindowTitle("Boids in a box")
	// if err := ebiten.RunGame(&Game{}); err != nil {
	// 	log.Fatal(err)
	// }
}
