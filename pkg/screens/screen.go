package screens

import "github.com/captncraig/dial-a-die/pkg/drawing"

type Screen interface {
	Render(*drawing.Image)
	OnDial(int) Screen
}
