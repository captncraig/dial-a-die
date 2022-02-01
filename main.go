package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"sync"

	"github.com/captncraig/dial-a-die/pkg/character"
	"github.com/captncraig/dial-a-die/pkg/dial"
	"github.com/captncraig/dial-a-die/pkg/drawing"
	"github.com/captncraig/dial-a-die/pkg/screens"
	"github.com/gorilla/websocket"
	"periph.io/x/host/v3"
)

var web = flag.Bool("web", false, "Disable all hardware, websocket server instead")
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var dabbs = &character.PC{
	FirstName:     "Dabbledopp",
	LastName:      "Fabblestabble",
	HP:            65,
	HPMax:         65,
	SpellSlots:    2,
	SpellSlotsMax: 2,
	Strength:      0,
	Dexterity:     5,
	Constitution:  1,
	Intelligence:  1,
	Wisdom:        3,
	Charisma:      3,

	Proficiency: 4,
	Saves:       []string{"Dexterity", "Intelligence"},
	Skills:      []string{"Deception", "History", "Investigation", "Perception", "Sleight of Hand", "Survival"},
	Expertise:   []string{"Investigation", "Perception"},
}

var dialChan chan int

func main() {
	flag.Parse()
	dialChan = make(chan int)
	if !*web {
		host.Init()
		err := dial.Init(dialChan)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		go func() {
			http.HandleFunc("/", handler)
			http.HandleFunc("/img", imghandler)
			http.HandleFunc("/poke", dialhandler)
			log.Fatal(http.ListenAndServe(":8080", nil))
		}()
	}
	go qWorker()
	go epdWorker()

	stack := ScreenStack{}
	stack.Push(screens.HomeScreen{PC: dabbs})

	render(stack.Top())

	for d := range dialChan {
		log.Printf("DIALED %d", d)
		nextScreen := stack.Top().OnDial(d)
		if nextScreen == nil {
			for len(stack.stack) > 1 {
				stack.Pop()
			}
		} else if nextScreen != stack.Top() {
			stack.Push(nextScreen)
		}
		render(stack.Top())
	}
}

var mostRecent *drawing.Image
var lock sync.Mutex

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "test.html")
}

func imghandler(w http.ResponseWriter, r *http.Request) {
	if mostRecent == nil {
		return
	}
	lock.Lock()
	defer lock.Unlock()
	if err := mostRecent.RenderPng(w); err != nil {
		log.Println(err)
	}
}

func dialhandler(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.URL.Query().Get("d"))
	if err != nil {
		log.Println(err)
	} else {
		dialChan <- n
	}
	w.WriteHeader(200)
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
			lock.Lock()
			mostRecent = r
			lock.Unlock()
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

func epdWorker() {
	if !*web {
		cmd := exec.Command("./epd")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		log.Println(cmd.Run())
	}
	for i := range execQueue {
		if !*web {
			fname, _ := i.Save()
			log.Println(fname)
			cmd := exec.Command("./epd", fname)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			log.Println(cmd.Run())
		}
		doneQueue <- true
	}
}
