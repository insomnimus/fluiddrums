package main

import (
"fmt"
	"flag"
	"github.com/eiannone/keyboard"
	"net"
)

var c net.Conn

type Drum string

func (d Drum) String() string {
	return string(d)
}

func (d Drum) Play() {
	c.Write([]byte("noteoff 9 " + d + " 0\n"))
	c.Write([]byte("noteon 9 " + d + " 120\n"))
}

var PRESETS = [7]string{
	"prog 9 000\n",
	"prog 9 008\n",
	"prog 9 016\n",
	"prog 9 024\n",
	"prog 9 032\n",
	"prog 9 040\n",
	"prog 9 048\n",
}

var cursor = 0

func ProgramChange(step int) {
	cursor += step
	if cursor < 0 {
		cursor = 6
	}
	if cursor > 6 {
		cursor = 0
	}
	c.Write([]byte(PRESETS[cursor]))
}

const (
	kickA  Drum = "36" // bass
	kickB  Drum = "35" // acoustic
	snareA Drum = "38" // acoustic
	snareB Drum = "40" // electric
	tomA   Drum = "41" // low floor
	tomB   Drum = "43" // high floor
	tomC   Drum = "45" // low
	tomD   Drum = "47" // low-mid
	tomE   Drum = "48" // high-mid
	tomF   Drum = "50" // high
	hatA   Drum = "42" // closed hi-hat
	hatB   Drum = "44" // pedal hi-hat
	hatC   Drum = "46" // open hi-hat
	crashA Drum = "49"
	crashB Drum = "57"
	rideA  Drum = "51" // ride cymbol 1
	rideB  Drum = "53" // ride bell
	rideC  Drum = "56" // cowbell
	rideD  Drum = "59" // ride cymbol 2
	miscA  Drum = "52" // chinese cymbol
	miscB  Drum = "37" // side stick
)

func main() {
	fmt.Println("press 1 and 2 to switch presets")
	var port string
	flag.StringVar(&port, "p", "localhost:9800", "fluidsynth port")
	flag.Parse()
	conn, err := net.Dial("tcp", port)
	if err != nil {
		fmt.Println("run fluidsynth in server mode first with -s flag")
		panic(err)
	}
	c = conn
	var start func()
	start = func() {
		keys, err := keyboard.GetKeys(10)
		if err != nil {
			panic(err)
		}
		defer keyboard.Close()
		var e keyboard.KeyEvent
		for {
			e = <-keys
			switch e.Rune {
			case 'h':
				kickA.Play()
				hatB.Play()
			case 'j', 'k':
				kickA.Play()
			case 'f':
				hatC.Play()
			case 'q':
				crashA.Play()
			case 'a':
				crashB.Play()
			case 'e':
				hatA.Play()
			case 'r':
				hatC.Play()
			case 'w':
				hatA.Play()
			case 'd':
				rideD.Play()
			case 's':
				miscB.Play()
			case 'l':
				snareB.Play()
			case 'u':
				tomE.Play()
			case 'ı', 'i':
				tomD.Play()
			case 'o':
				tomC.Play()
			case 'p':
				tomB.Play()
			case 'ğ':
				tomA.Play()
			case 'g':
				rideA.Play()
			case '1':
				ProgramChange(-1)
			case '2':
				ProgramChange(1)
			}
		}
	}
	start()
}
