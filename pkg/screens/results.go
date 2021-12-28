package screens

import (
	"log"

	"github.com/captncraig/dial-a-die/pkg/drawing"
)

type RollResultsScreen struct{}

func (h RollResultsScreen) Render(img *drawing.Image) {
	log.Println("RR RENDER")
	img.TextCenter(20, 20, "Booming Blade")
	img.TextCenter(20, 40, "(Sneak Hex)")

	img.TextCenter(50, 100, "23", "113")
	img.TextCenter(18, 120, "To Hit", "Damage")

	img.TextCenter(30, 160, "15", "23", "5", "6", "3")
	img.TextCenter(18, 180, "Prc", "Snk", "Zap", "Nex", "Blu")

}

func (h RollResultsScreen) OnDial(d int) Screen {
	log.Println("RR DIAL")
	return nil
}
