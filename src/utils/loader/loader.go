package loader

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/tiff"
)
func LoadImage(path string, width, height *uint32 ) ([]byte, error) {
	file, err := OpenFile(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, s, err := image.Decode(file)
	fmt.Println(s)
	if err != nil {
		return nil, err
	}
	*width = uint32(img.Bounds().Dx())
	*height = uint32(img.Bounds().Dy())
	rgbaImg := image.NewRGBA(img.Bounds())
	draw.Draw(rgbaImg,rgbaImg.Bounds(),img, image.Pt(0,0), draw.Src)
	return rgbaImg.Pix, nil
}