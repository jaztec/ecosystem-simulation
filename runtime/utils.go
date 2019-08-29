package runtime

import (
	"image"
	"os"

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
