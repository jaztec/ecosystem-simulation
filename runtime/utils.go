package runtime

import (
	"image"
	"math/rand"
	"os"

	"github.com/jaztec/ecosystem-simulation/world"

	_ "image/png" // make sure we can decode PNG images

	"github.com/faiface/pixel"
	"golang.org/x/xerrors"
)

// LoadPicture reads a file from disk and returns a pixel wrapped image
func LoadPicture(p string) (pixel.Picture, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, xerrors.Errorf("error when opening image: %w", err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, xerrors.Errorf("error when decoding image: %w", err)
	}

	return pixel.PictureDataFromImage(img), nil
}

// GenerateWorldLayout will return some randomly generated terrain layout
func GenerateWorldLayout(w, h int) world.TerrainLayout {
	tl := make([][]world.TerrainTile, 0, h)
	for x := 0; x < h; x++ {
		r := make([]world.TerrainTile, 0, w)
		for y := 0; y < w; y++ {
			r = append(r, world.TerrainTile(rand.Intn(3)))
		}
		tl = append(tl, r)
	}
	return tl
}
