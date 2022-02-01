package screens

import (
	"fmt"

	"github.com/captncraig/dial-a-die/pkg/drawing"
)

type MenuScreen struct {
	items  []*MenuItem
	big    bool
	offset int
}

type MenuItem struct {
	title string
	f     rollfunc
}

func NewMenu(mis ...*MenuItem) *MenuScreen {
	ms := &MenuScreen{
		items:  mis,
		big:    len(mis) > 10,
		offset: 0,
	}
	return ms
}

func (s *MenuScreen) Render(img *drawing.Image) {
	y := 50
	start := s.offset
	idx := 1
	if s.big {
		img.Text("1: Next", 18, 4, y)
		y += 20
		idx = 2
	}
	for i := 0; i < 9; i++ {
		if start+i >= len(s.items) {
			break
		}
		item := s.items[start+i]
		txt := fmt.Sprintf("%d: %s", idx+i, item.title)
		img.Text(txt, 18, 4, y)
		y += 20
	}
}

func (h *MenuScreen) OnDial(d int) Screen {
	if d == 0 {
		return nil
	}
	if h.big && d == 1 {
		h.offset += 9
		if h.offset >= len(h.items) {
			h.offset = 0
		}
		return h
	}
	idx := d - 1
	if h.big {
		idx = d - 2 + h.offset
	}
	if idx >= len(h.items) {
		return h
	}
	item := h.items[idx]
	return NewRollResults(item.f)
}
