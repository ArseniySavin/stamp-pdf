package pkg

import (
	"image"
	"strings"

	"github.com/fogleman/gg"
)

func AddTextOnImg(inFile, text, executor, applicant string) (image.Image, error) {
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

	tx := float64((imgWidth / 2) + 225)
	y := float64((imgHeight))
	maxWidth := float64(imgWidth)

	dc.SetRGB(0.18, 0.33, 0.59)

	if str1, str2, isFound := CutStrByNewline(applicant); isFound {
		dc.DrawStringWrapped(str1, tx, y-100, 0.5, 0.5, maxWidth, 1.5, gg.AlignLeft)
		dc.DrawStringWrapped(str2, tx, y-78, 0.5, 0.5, maxWidth, 1.5, gg.AlignLeft)

	} else {
		dc.DrawStringWrapped(applicant, tx, y-100, 0.5, 0.5, maxWidth, 1.5, gg.AlignLeft)
	}
	dc.DrawStringWrapped(text, tx, y-53, 0.5, 0.5, maxWidth, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(executor, tx, y-27, 0.5, 0.5, maxWidth, 1.5, gg.AlignLeft)

	return dc.Image(), nil
}

func Save(img image.Image, path string) error {
	if err := gg.SavePNG(path, img); err != nil {
		return err
	}
	return nil
}

func CutStrByNewline(text string) (str1 string, str2 string, found bool) {
	return strings.Cut(text, "\\n")
}
