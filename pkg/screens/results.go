package screens

import (
	"fmt"
	"log"
	"math/rand"

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

func (s *RollResultsScreen) Render(img *drawing.Image) {
	img.TextCenter(20, 20, s.Title)
	img.TextCenter(20, 40, s.Subtitle)

	img.TextCenterRows(50, 100, 18, 120, s.MainNumbers, s.MainTitles)

	img.TextCenterRows(30, 160, 12, 180, s.SecondaryNumbers, s.SecondaryTitles)

	for len(s.DetailText) < len(s.Details) {
		s.DetailText = append(s.DetailText, "")
	}
	inRow := 5
	y1 := 240
	y2 := 260
	offset := 50
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
		img.TextCenterRows(25, y1, 17, y2, row1, row2)
		y1 += offset
		y2 += offset
	}

	for i := 0; i < len(s.Details); i += inRow {

	}
}

func (h *RollResultsScreen) OnDial(d int) Screen {
	log.Println("RR DIAL")
	return nil
}

func (h *RollResultsScreen) Roll(n int, reason string) int {
	v := rand.Intn(n) + 1
	g := ""
	switch n {
	case 20:
		g = "¤"
	case 8:
		g = "#"
	case 6:
		g = "*"
	case 4:
		g = "^"
	}
	h.Details = append(h.Details, fmt.Sprintf("%s%d", g, v))
	h.DetailText = append(h.DetailText, reason)
	return v
}

func (h *RollResultsScreen) AddMod(n int, reason string) int {
	h.Details = append(h.Details, fmt.Sprintf("+%d", n))
	h.DetailText = append(h.DetailText, reason)
	return n
}

func (h *RollResultsScreen) AddMain(n int, reason string, crit bool) {
	text := fmt.Sprint(n)
	if crit {
		text += "!"
	}
	h.MainNumbers = append(h.MainNumbers, text)
	h.MainTitles = append(h.MainTitles, reason)
}

func (h *RollResultsScreen) AddSecondary(n int, reason string) {
	h.SecondaryNumbers = append(h.SecondaryNumbers, fmt.Sprint(n))
	h.SecondaryTitles = append(h.SecondaryTitles, reason)
}
