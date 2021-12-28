package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/captncraig/dial-a-die/pkg/dial"
	"github.com/captncraig/dial-a-die/pkg/drawing"
	"github.com/captncraig/dial-a-die/pkg/screens"
	"periph.io/x/host/v3"
)

func main() {
	host.Init()
	dialChan, err := dial.Init()
	if err != nil {
		log.Fatal(err)
	}

	go qWorker()
	go exeWorker()

	stack := ScreenStack{}
	stack.Push(screens.HomeScreen{})

	render(stack.Top())

	for d := range dialChan {
		log.Printf("DIALED %d", d)
		nextScreen := stack.Top().OnDial(d)
		if nextScreen == nil {
			stack.Pop()
		} else if nextScreen != stack.Top() {
			stack.Push(nextScreen)
		}
		render(stack.Top())
	}
}

func render(s screens.Screen) {
	img := drawing.New()
	s.Render(img)
	renderQueue <- img
}

type ScreenStack struct {
	stack []screens.Screen
}

func (s *ScreenStack) Top() screens.Screen {
	return s.stack[len(s.stack)-1]
}

func (s *ScreenStack) Push(sc screens.Screen) {
	log.Println("PUSH!")
	s.stack = append(s.stack, sc)
}

func (s *ScreenStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
}

var renderQueue = make(chan *drawing.Image)
var execQueue = make(chan *drawing.Image)
var doneQueue = make(chan bool)

func qWorker() {
	running := false
	var nextUp *drawing.Image
	for {
		select {
		case r := <-renderQueue:
			if !running {
				running = true
				execQueue <- r
			} else {
				nextUp = r
			}
		case <-doneQueue:
			if nextUp == nil {
				running = false
			} else {
				execQueue <- nextUp
				nextUp = nil
			}
		}
	}
}

func exeWorker() {
	cmd := exec.Command("./epd")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(cmd.Run())
	for i := range execQueue {
		fname, _ := i.Save()
		log.Println(fname)
		cmd := exec.Command("./epd", fname)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		log.Println(cmd.Run())
		doneQueue <- true
	}
}
