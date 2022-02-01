package screens

import (
	"fmt"
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
	Rolls            []int

	f       func(rrs *RollResultsScreen)
	ops     []rollfunc
	opnames []string
}

type rollfunc func(rrs *RollResultsScreen)

func NewRollResults(f rollfunc, ops []rollfunc, names []string) *RollResultsScreen {
	rrs := &RollResultsScreen{
		f:       f,
		ops:     ops,
		opnames: names,
	}
	f(rrs)
	return rrs
}

func (s *RollResultsScreen) clear() {
	s.Title = ""
	s.Subtitle = ""
	s.MainNumbers = nil
	s.MainTitles = nil
	s.SecondaryNumbers = nil
	s.SecondaryTitles = nil
	s.Details = nil
	s.DetailText = nil
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

	if len(s.ops) > 3 {
		img.TextRight(fmt.Sprintf("6: %s", s.opnames[3]), 15, 315)
	}
	if len(s.ops) > 2 {
		img.TextRight(fmt.Sprintf("5: %s", s.opnames[2]), 15, 330)
	}
	if len(s.ops) > 1 {
		img.TextRight(fmt.Sprintf("4: %s", s.opnames[1]), 15, 345)
	}
	if len(s.ops) > 0 {
		img.TextRight(fmt.Sprintf("3: %s", s.opnames[0]), 15, 360)
	}
	img.TextRight("2: reroll", 15, 375)
	img.TextRight("*: home", 15, 390)

}

func (h *RollResultsScreen) OnDial(d int) Screen {
	if d == 2 {
		h.clear()
		h.f(h)
		return h
	}
	if d == 3 && len(h.ops) > 0 {
		h.ops[0](h)
		return h
	}
	if d == 4 && len(h.ops) > 1 {
		h.ops[1](h)
		return h
	}
	if d == 5 && len(h.ops) > 2 {
		h.ops[2](h)
		return h
	}
	if d == 6 && len(h.ops) > 3 {
		h.ops[3](h)
		return h
	}
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
	h.Rolls = append(h.Rolls, v)
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
