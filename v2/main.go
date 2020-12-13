package main

import (
	"flag"
	"fmt"
	"github.com/nsf/termbox-go"
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
	floorA Drum = "41" // low floor
	floorB Drum = "43" // high floor
	tomA   Drum = "45" // low
	tomB   Drum = "47" // low-mid
	tomC   Drum = "48" // high-mid
	//tomF   Drum = "50" // high
	hatC   Drum = "42" // closed hi-hat
	hatB   Drum = "44" // pedal hi-hat
	hatA   Drum = "46" // open hi-hat
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
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	fmt.Println(`
	hit some keys to play, (j, i and l mix well)
	hit 1 and 2 to change presets
	hit dot (.) comma(,) or dash(-) to exit.
	`)
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
		var e termbox.Event
		for {
			e = termbox.PollEvent()
			if e.Type != termbox.EventKey {
				continue
			}
			switch e.Ch {
			case '.', ',', '-':
				return
			case 'j', 'h':
				kickA.Play()
			case 'l':
				snareA.Play()
			case 'ı', 'i':
				hatA.Play()
			case 'f', 'k':
				rideA.Play()
			case 'a':
				crashA.Play()
			case 'q':
				floorA.Play()
			case 'g':
				floorB.Play()
			case 'e':
				tomA.Play()
			case 's':
				hatB.Play()
			case 'r':
				tomB.Play()
			case 'w':
				snareA.Play()
			case 't':
				tomC.Play()
				//hatA.Play()
			case 'b':
				hatB.Play()
			case 'n', 'm':
				kickB.Play()
			case 'd':
				miscA.Play()
			case '1':
				ProgramChange(-1)
			case '2':
				ProgramChange(1)
			}
		}
	}
	start()
}
