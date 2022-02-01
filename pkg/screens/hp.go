package screens

import (
	"fmt"

	"github.com/captncraig/dial-a-die/pkg/drawing"
)

type HPScreen struct {
	Heal bool
}

func (s *HPScreen) Render(img *drawing.Image) {
	img.TextCenter(50, 100, fmt.Sprintf("%d/%d", PC.HP, PC.HPMax))
}

func (h *HPScreen) OnDial(d int) Screen {
	if d == 0 {
		return nil
	}
	if h.Heal {
		PC.HP += d
	} else {
		PC.HP -= d
	}
	if PC.HP < 0 {
		PC.HP = 0
	}
	if PC.HP > PC.HPMax {
		PC.HP = PC.HPMax
	}
	return h
}
