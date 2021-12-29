package screens

import (
	"fmt"
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

	// hp
	img.TextRight("45", 40, 80)
	img.TextRight("/60", 40, 120)

	// spell slots ¤×
	img.TextRight("ss ¤×", 16, 150)
	img.TextRight("misty ¤¤×××", 16, 175)

	// bottom menu
	menuSize := float64(16)
	menuSpacing := 15
	for i, text := range []string{"Actions", "Saves", "Checks", "Spells", "HP/Rest", "Dice", "", "", "Reference", "Chars"} {
		if text == "" {
			continue
		}
		row := i / 2
		txt := fmt.Sprintf("%d: %s", i+1, text)
		y := 330 + menuSpacing*row
		if i%2 == 0 {
			img.Text(txt, menuSize, 0, y)
		} else {
			img.TextRight(txt+"  ", menuSize, y)
		}
	}
}

func (h HomeScreen) OnDial(d int) Screen {
	log.Printf("Home Dial %d", d)
	switch d {
	case 1:
		//actions
	case 2:
		return NewRollResults(func(rrs *RollResultsScreen) {
			rrs.Title = "Booming Blade"
			rrs.Subtitle = "(Sneak, Hex, Adv)"
			a := rrs.Roll(20, "Atk")
			b := rrs.Roll(20, "Adv")
			// apply advantage
			if b > a {
				a = b
			}
			rrs.AddMod(9, "Atk")
			rrs.AddMain(a+9, "To Hit", a == 20)
			slash := rrs.Roll(8, "Atk") + rrs.AddMod(6, "Dmg")
			bludgeoning := rrs.AddMod(3, "Blg")
			boom := rrs.Roll(8, "Boom")
			hex := rrs.Roll(6, "Hex")
			sneak := rrs.Roll(6, "Snk") + rrs.Roll(6, "Snk")
			rrs.AddMain(slash+bludgeoning+boom+hex+sneak, "Damage", false)
			rrs.AddSecondary(slash, "Slsh")
			rrs.AddSecondary(bludgeoning, "Blg")
			rrs.AddSecondary(boom, "Boom")
			rrs.AddSecondary(hex, "Nec")
			rrs.AddSecondary(sneak, "snk")
		})

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
