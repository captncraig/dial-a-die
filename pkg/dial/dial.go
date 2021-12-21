package dial

import (
	"fmt"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
)

var BounceDelay = 5 * time.Millisecond
var LatchPinNum = 12
var PulsePinNum = 16

func Init() (<-chan int, error) {
	latchPin, err := newDebouncer(LatchPinNum)
	if err != nil {
		return nil, err
	}
	pulsePin, err := newDebouncer(PulsePinNum)
	if err != nil {
		return nil, err
	}
	ch := make(chan int)
	go readDial(latchPin, pulsePin, ch)
	return ch, nil
}

func readDial(latchPin, pulsePin *debouncer, ch chan<- int) {
	ticks := time.NewTicker(time.Millisecond)

	dialing := false
	pulsing := false
	pulses := 0

	for range ticks.C {
		latch := bool(latchPin.Read())
		pulse := bool(pulsePin.Read())

		//latch goes low to signify dialing
		if !dialing && !latch {
			dialing = true
			pulses = 0
			pulsing = false
		} else if dialing && latch {
			dialing = false
			ch <- pulses
		}

		// pulse goes high to signify pulse
		if dialing && pulse && !pulsing {
			pulsing = true
			pulses++
		} else if dialing && !pulse && pulsing {
			pulsing = false
		}
	}
}

type debouncer struct {
	state      gpio.Level
	lastState  gpio.Level
	lastChange time.Time
	pin        gpio.PinIO
}

func newDebouncer(pinNum int) (*debouncer, error) {
	pin := gpioreg.ByName(fmt.Sprint(pinNum))
	if pin == nil {
		return nil, fmt.Errorf("Failed to find pin %d", pinNum)
	}
	pin.In(gpio.PullUp, gpio.NoEdge)
	return &debouncer{
		pin:        pin,
		state:      pin.Read(),
		lastState:  pin.Read(),
		lastChange: time.Now(),
	}, nil
}

func (d *debouncer) Read() gpio.Level {
	now := time.Now()
	reading := d.pin.Read()
	if reading != d.lastState {
		d.lastChange = now
	}
	if (now.Sub(d.lastChange)) > BounceDelay && reading != d.state {
		d.state = reading
	}
	d.lastState = reading
	return d.state
}
