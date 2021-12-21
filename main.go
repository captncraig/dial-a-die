package main

import (
	"log"

	"github.com/captncraig/dial-a-die/pkg/dial"
	"github.com/captncraig/dial-a-die/pkg/screens"
	"periph.io/x/host/v3"
)

func main() {
	host.Init()

	dialChan, err := dial.Init()
	if err != nil {
		log.Fatal(err)
	}

	stack := ScreenStack{}
	stack.Push(screens.HomeScreen{})
	stack.Top().Render()

	for d := range dialChan {
		log.Printf("DIALED %d", d)
		nextScreen := stack.Top().OnDial(d)
		if nextScreen == nil {
			stack.Pop()
		} else if nextScreen != stack.Top() {
			stack.Push(nextScreen)
		}
		stack.Top().Render()
	}
}

type ScreenStack struct {
	stack []screens.Screen
}

func (s *ScreenStack) Top() screens.Screen {
	return s.stack[len(s.stack)-1]
}

func (s *ScreenStack) Push(sc screens.Screen) {
	s.stack = append(s.stack, sc)
}

func (s *ScreenStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
}
