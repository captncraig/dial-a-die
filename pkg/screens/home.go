package screens

import "log"

type HomeScreen struct{}

func (h HomeScreen) Render() {
	log.Println("Home Screen Render!")
}

func (h HomeScreen) OnDial(d int) Screen {
	log.Printf("Home Dial %d", d)
	switch d {
	case 1:
		//actions
	case 2:
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
