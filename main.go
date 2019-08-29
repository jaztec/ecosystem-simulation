package main

import (
	"context"

	"github.com/faiface/pixel/pixelgl"
	"github.com/jaztec/ecosystem-simulation/runtime"
)

func run() {
	app, err := runtime.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Run(context.Background())
}

func main() {
	pixelgl.Run(run)
}