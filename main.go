package main

import (
	"context"

	"github.com/jaztec/ecosystem-simulation/application"
	"github.com/jaztec/ecosystem-simulation/world"

	"github.com/faiface/pixel/pixelgl"
)

func run() {
	app, err := application.NewApplication(application.Config{Layout: world.GenerateWorldLayout(20, 15)})
	if err != nil {
		panic(err)
	}

	app.Run(context.Background())
}

func main() {
	pixelgl.Run(run)
}
