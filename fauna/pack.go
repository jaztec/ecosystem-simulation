package fauna

import "github.com/faiface/pixel"

type PackConfig struct {
	WolfPicture pixel.Picture
	Bounds      pixel.Rect
}

type Pack struct {
	sprites [5][4]*pixel.Sprite
	pack    []*Wolf
	batch   *pixel.Batch
}
