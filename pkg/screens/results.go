package screens

import (
	"log"

	"github.com/captncraig/dial-a-die/pkg/drawing"
)

type RollResultsScreen struct {
	Title            string
	Subtitle         string
	MainNumbers      []string
	MainTitles       []string
	SecondaryNumbers []string
	SecondaryTitles  []string
	Details          []string
	DetailText       []string
}

func (s RollResultsScreen) Render(img *drawing.Image) {
	img.TextCenter(20, 20, s.Title)
	img.TextCenter(20, 40, s.Subtitle)

	img.TextCenterRows(50, 100, 18, 120, s.MainNumbers, s.MainTitles)

	img.TextCenterRows(30, 160, 12, 180, s.SecondaryNumbers, s.SecondaryTitles)

	inRow := 5
	y1 := 240
	y2 := 260
	offset := 80
	d := s.Details
	t := s.DetailText
	for len(d) > 0 {
		row1 := d
		row2 := t
		if len(d) > inRow {
			row1 = d[:inRow]
			d = d[inRow:]
			row2 = t[:inRow]
			t = t[inRow:]
		} else {
			d = nil
			t = nil
		}
		img.TextCenterRows(24, y1, 14, y2, row1, row2)
		y1 += offset
		y2 += offset
	}

	for i := 0; i < len(s.Details); i += inRow {

	}
}

func (h RollResultsScreen) OnDial(d int) Screen {
	log.Println("RR DIAL")
	return nil
}
