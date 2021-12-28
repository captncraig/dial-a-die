package screens

import (
	"log"

	"github.com/captncraig/dial-a-die/pkg/drawing"
)

type HomeScreen struct{}

func (h HomeScreen) Render(img *drawing.Image) {
	log.Println("Home Screen Render!")

	img.TextCenter(20, 20, "Dabbledopp")
	img.TextCenter(20, 40, "Fabblestabble")

	img.Text("STR:", 18, 0, 60)
	img.Text("DEX:", 18, 0, 80)
	img.Text("CON:", 18, 0, 100)
	img.Text("INT:", 18, 0, 120)
	img.Text("WIS:", 18, 0, 140)
	img.Text("CHA:", 18, 0, 160)
	img.Text("AC:", 18, 0, 180)
	img.Text("Init:", 18, 0, 200)
	img.Text("Percep:", 18, 0, 220)

	xalign := 95
	img.Text("10(0)", 18, xalign, 60)
	img.Text("20(5)", 18, xalign, 80)
	img.Text("12(1)", 18, xalign, 100)
	img.Text("12(1)", 18, xalign, 120)
	img.Text("16(3)", 18, xalign, 140)
	img.Text("16(3)", 18, xalign, 160)
	img.Text("18", 18, xalign, 180)
	img.Text("+5", 18, xalign, 200)
	img.Text("19", 18, xalign, 220)

	img.TextRight("45", 40, 80)
	img.TextRight("/60", 40, 120)

}

func (h HomeScreen) OnDial(d int) Screen {
	log.Printf("Home Dial %d", d)
	switch d {
	case 1:
		//actions
	case 2:
		return RollResultsScreen{}
		//saves
	case 3:
		// checks
	case 4:
		// spells
	case 5:
		// status(hp,money,rests)
	case 9:
		// arbitrary dice
	case 0:
		// reference
	}
	return h
}
