package screens

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/captncraig/dial-a-die/pkg/character"
	"github.com/captncraig/dial-a-die/pkg/drawing"
)

var PC *character.PC

type HomeScreen struct {
}

func (h HomeScreen) Render(img *drawing.Image) {
	log.Println("Home Screen Render!")

	img.TextCenter(20, 20, PC.FirstName)
	img.TextCenter(20, 40, PC.LastName)

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
	img.TextRight(fmt.Sprint(PC.HP), 40, 80)
	img.TextRight(fmt.Sprintf("/%d", PC.HPMax), 40, 120)

	// spell slots ¤×
	img.TextRight(fmt.Sprintf("ss: %d/%d", PC.SpellSlots, PC.SpellSlotsMax), 18, 150)
	img.TextRight(fmt.Sprintf("misty: %d/5", PC.Misty), 18, 175)

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
		mis := []*MenuItem{
			{title: "ATTACK!", f2: func() Screen {
				return NewRollResults(func(rrs *RollResultsScreen) {
					rrs.Title = "Booming Blade"
					a := rrs.Roll(20, "Atk")

					rrs.AddMod(9, "Atk")
					rrs.AddMain(a+9, "To Hit", a == 20)
					slash := rrs.Roll(8, "Atk") + rrs.AddMod(6, "Dmg")
					bludgeoning := rrs.AddMod(3, "Blg")
					boom := rrs.Roll(8, "Boom")
					rrs.AddMain(slash+bludgeoning+boom, "Damage", false)
					rrs.AddSecondary(slash, "Slsh")
					rrs.AddSecondary(bludgeoning, "Blg")
					rrs.AddSecondary(boom, "Boom")

				}, []rollfunc{
					func(rrs *RollResultsScreen) {
						d := rrs.Roll(20, "Adv")
						if d > rrs.Rolls[0] {
							diff := d - rrs.Rolls[0]
							topnum, _ := strconv.Atoi(strings.TrimSuffix(rrs.MainNumbers[0], "!"))
							newVal := fmt.Sprint(topnum + diff)
							if d == 20 {
								newVal += "!"
							}
							rrs.MainNumbers[0] = newVal
						}
						rrs.ops = rrs.ops[2:]
						rrs.opnames = rrs.opnames[2:]
					},
					func(rrs *RollResultsScreen) {
						d := rrs.Roll(20, "Dis")
						if d < rrs.Rolls[0] {
							diff := rrs.Rolls[0] - d
							topnum, _ := strconv.Atoi(strings.TrimSuffix(rrs.MainNumbers[0], "!"))
							newVal := fmt.Sprint(topnum - diff)
							rrs.MainNumbers[0] = newVal
						}
						rrs.ops = rrs.ops[2:]
						rrs.opnames = rrs.opnames[2:]
					},
					func(rrs *RollResultsScreen) {
						sneak := rrs.Roll(6, "Snk") + rrs.Roll(6, "Snk")
						rrs.AddSecondary(sneak, "snk")
						topnum, _ := strconv.Atoi(rrs.MainNumbers[1])
						rrs.MainNumbers[1] = fmt.Sprint(topnum + sneak)
					},
					func(rrs *RollResultsScreen) {
						hex := rrs.Roll(6, "Hex")
						rrs.AddSecondary(hex, "hex")
						topnum, _ := strconv.Atoi(rrs.MainNumbers[1])
						rrs.MainNumbers[1] = fmt.Sprint(topnum + hex)
					},
				},
					[]string{
						"Advantage",
						"Disadvantage",
						"Sneak",
						"Hex",
					})
			}},
			{title: "Use Spell", f2: func() Screen {
				PC.SpellSlots--
				return nil
			}},
			{title: "Short Rest", f2: func() Screen {
				PC.SpellSlots = PC.SpellSlotsMax
				return nil
			}},
			{title: "Long Rest", f2: func() Screen {
				PC.SpellSlots = PC.SpellSlotsMax
				PC.HP = PC.HPMax
				PC.Misty = 5
				return nil
			}},
			{title: "Misty Step", f2: func() Screen {
				PC.Misty--
				return nil
			}},
			{title: "Initiative", f: func(rrs *RollResultsScreen) {
				rrs.Title = "Initiative"
				d := rrs.Roll(20, "1d20")
				m := rrs.AddMod(PC.Dexterity, "Dex")
				rrs.AddMain(d+m, "Initiative", d == 20)
			}},
		}
		return NewMenu(mis...)
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
				mod := PC.Mod(attr)
				rrs.AddMod(mod, "Mod")
				p := PC.SaveProficient(attr)
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
				mod := PC.Mod(attr)
				p := PC.SkillProficient(sk)
				if p != 0 {
					rrs.AddMod(p, "Pro")
				}
				e := PC.ExpertiseMod(sk)
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
		mis := []*MenuItem{
			{title: "Heal", f2: func() Screen { return &HPScreen{true} }},
			{title: "Hurt", f2: func() Screen { return &HPScreen{false} }},
		}
		return NewMenu(mis...)
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
