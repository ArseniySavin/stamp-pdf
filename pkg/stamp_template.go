package pkg

import (
	"image"

	"github.com/fogleman/gg"
)

func AddTextOnImg(inFile, text, executor string) (image.Image, error) {
	bgImage, err := gg.LoadImage(inFile)
	if err != nil {
		return nil, err
	}
	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)

	if err := dc.LoadFontFace("./stamp.ttf", 14); err != nil {
		return nil, err
	}

	tx := float64((imgWidth / 2) + 100)
	y := float64((imgHeight))
	maxWidth := float64(imgWidth)

	dc.SetRGB(0.18, 0.33, 0.59)

	dc.DrawStringWrapped(text, tx, y-53, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)
	dc.DrawStringWrapped(executor, tx, y-27, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)

	return dc.Image(), nil
}

func Save(img image.Image, path string) error {
	if err := gg.SavePNG(path, img); err != nil {
		return err
	}
	return nil
}
