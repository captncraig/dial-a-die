package drawing

import (
	"image"
	"image/color"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

const fontfile = "YacimientoExtraBoldEx.ttf"

var textFont *truetype.Font

type Image struct {
	img *image.Gray
}

func init() {
	fontBytes, err := os.ReadFile(fontfile)
	if err != nil {
		log.Fatal(err)
	}
	textFont, err = truetype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}
}

func New() *Image {
	im := image.NewGray(image.Rect(0, 0, 300, 400))
	for i := range im.Pix {
		im.Pix[i] = 255
	}
	return &Image{
		img: im,
	}
}

var Totwidth = fixed.Int26_6(300 * 64)

func (i *Image) Text(s string, size float64, x, y int) {
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}
	i.TextPoint(s, size, point)
}

func (i *Image) TextRight(s string, size float64, y int) {
	width := i.Measure(s, size)
	point := fixed.Point26_6{Totwidth - width, fixed.Int26_6(y * 64)}
	i.TextPoint(s, size, point)
}

func (i *Image) TextCenter(size float64, y int, ss ...string) {
	widths := make([]fixed.Int26_6, len(ss))
	total := fixed.Int26_6(0)
	for ix, s := range ss {
		w := i.Measure(s, size)
		widths[ix] = w
		total += w
	}
	spacer := (Totwidth - total) / (fixed.Int26_6(len(ss)) + 1)
	xpos := spacer
	for ix, s := range ss {
		point := fixed.Point26_6{xpos, fixed.Int26_6(y * 64)}
		i.TextPoint(s, size, point)
		xpos += widths[ix] + spacer
	}
}

func (i *Image) TextPoint(s string, size float64, point fixed.Point26_6) {
	col := color.Gray{0}

	d := &font.Drawer{
		Dst: i.img,
		Src: image.NewUniform(col),
		Face: truetype.NewFace(
			textFont, &truetype.Options{
				Size:    size,
				DPI:     72,
				Hinting: font.HintingFull,
			},
		),
		Dot: point,
	}
	d.DrawString(s)
}

func (i *Image) Measure(s string, size float64) fixed.Int26_6 {
	col := color.Gray{0}

	d := &font.Drawer{
		Dst: i.img,
		Src: image.NewUniform(col),
		Face: truetype.NewFace(
			textFont, &truetype.Options{
				Size:    size,
				DPI:     72,
				Hinting: font.HintingFull,
			},
		),
	}
	return d.MeasureString(s)
}

func (i *Image) Circle(x, y, d int) {

}

func (i *Image) Save() (string, error) {
	f, err := os.CreateTemp(os.TempDir(), "img")
	if err != nil {
		return "", err
	}

	defer f.Close()
	f.Write([]byte{
		// bitmap header
		// magic id
		0x42, 0x4d,
		// size
		0xd8, 0xea, 0, 0,
		// reserved
		0, 0, 0, 0,
		// start of data
		0x76, 0, 0, 0,
		//BITMAPINFOHEADER
		// size
		0x28, 0, 0, 0,
		// bmp width
		0x90, 1, 0, 0,
		// bmp height
		0x2c, 1, 0, 0,
		//color planes
		1, 0,
		// bpp
		4, 0,
		//compression
		0, 0, 0, 0,
		// img size
		0x62, 0xea, 0, 0,
		// hor_rez
		0x12, 0x0b, 0, 0,
		// ver_rez
		0x12, 0x0b, 0, 0,
		// ncolors
		0, 0, 0, 0,
		// important_colors
		0, 0, 0, 0,
		// now colors
		0, 0, 0, 0,
		0x11, 0x11, 0x11, 0,
		0x22, 0x22, 0x22, 0,
		0x33, 0x33, 0x33, 0,
		0x44, 0x44, 0x44, 0,
		0x55, 0x55, 0x55, 0,
		0x66, 0x66, 0x66, 0,
		0x77, 0x77, 0x77, 0,
		0x88, 0x88, 0x88, 0,
		0x99, 0x99, 0x99, 0,
		0xaa, 0xaa, 0xaa, 0,
		0xbb, 0xbb, 0xbb, 0,
		0xcc, 0xcc, 0xcc, 0,
		0xdd, 0xdd, 0xdd, 0,
		0xee, 0xee, 0xee, 0,
		0xff, 0xff, 0xff, 0,
	})
	for x := 0; x < 300; x++ {
		for y := 0; y < 400; y += 2 {
			p := i.img.At(x, y).(color.Gray)
			p2 := i.img.At(x, y+1).(color.Gray)

			b := byte((p.Y << 4) | (p2.Y & 0x0f))
			f.Write([]byte{b})
		}
	}
	return f.Name(), nil
}