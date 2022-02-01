package screens

import (
	"fmt"
	"log"
	"sort"

	"github.com/captncraig/dial-a-die/pkg/character"
	"github.com/captncraig/dial-a-die/pkg/drawing"
)

type HomeScreen struct {
	PC *character.PC
}

func (h HomeScreen) Render(img *drawing.Image) {
	log.Println("Home Screen Render!")

	img.TextCenter(20, 20, h.PC.FirstName)
	img.TextCenter(20, 40, h.PC.LastName)

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
	img.TextRight(fmt.Sprint(h.PC.HP), 40, 80)
	img.TextRight(fmt.Sprintf("/%d", h.PC.HPMax), 40, 120)

	// spell slots ¤×
	img.TextRight(fmt.Sprintf("ss: %d/%d", h.PC.SpellSlots, h.PC.SpellSlotsMax), 18, 150)
	img.TextRight("misty: 4/5", 18, 175)

	// bottom menu
	menuSize := float64(16)
	menuSpacing := 15
	for i, text := range []string{"Actions", "Saves", "Checks", "HP/Rest", "Dice"} {
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
	case 2:
		//saves
		attrs := []string{
			"Strength", "Dexterity", "Constitution", "Intelligence", "Wisdom", "Charisma",
		}
		mis := []*MenuItem{}
		for _, a := range attrs {
			attr := a
			mis = append(mis, &MenuItem{title: attr, f: func(rrs *RollResultsScreen) {
				rrs.Title = fmt.Sprintf("%s Save", attr)
				d := rrs.Roll(20, "Save")
				mod := h.PC.Mod(attr)
				rrs.AddMod(mod, "Mod")
				p := h.PC.SaveProficient(attr)
				if p != 0 {
					rrs.AddMod(p, "Pro")
				}
				rrs.AddMain(d+mod+p, "Save", d == 20)
			}})
		}
		return NewMenu(mis...)
	case 3:
		// checks
		mis := []*MenuItem{}
		for k, v := range character.Skills {
			sk := k
			attr := v
			mis = append(mis, &MenuItem{title: sk, f: func(rrs *RollResultsScreen) {
				rrs.Title = fmt.Sprintf("%s(%s)", sk, attr)
				rrs.Subtitle = "Check"
				d := rrs.Roll(20, "Check")
				mod := h.PC.Mod(attr)
				p := h.PC.SkillProficient(sk)
				if p != 0 {
					rrs.AddMod(p, "Pro")
				}
				e := h.PC.ExpertiseMod(sk)
				if e != 0 {
					rrs.AddMod(p, "Exp")
				}
				rrs.AddMod(mod, "mod")
				rrs.AddMain(d+mod+p+e, "Check", d == 20)
			}})
		}
		sort.Slice(mis, func(i, j int) bool {
			return mis[i].title < mis[j].title
		})
		return NewMenu(mis...)
	case 4:
		// status(hp,money,rests)
	case 5:
		// arbitrary dice
		mis := []*MenuItem{}
		for _, d := range []int{4, 6, 8, 10, 20, 100} {
			dc := d
			mis = append(mis, &MenuItem{title: fmt.Sprintf("d%d", dc), f: func(rrs *RollResultsScreen) {
				d := rrs.Roll(dc, fmt.Sprintf("1d%d", dc))
				rrs.Title = fmt.Sprintf("Roll 1d%d", dc)
				rrs.AddMain(d, "Result", false)
			}})
		}
		return NewMenu(mis...)
	case 0:
		// reference
	}
	return h
}
