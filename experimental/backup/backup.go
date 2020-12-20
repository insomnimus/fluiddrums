package main

import (
	"flag"
	"fmt"
	"github.com/nsf/termbox-go"
	"net"
)

var c net.Conn

type Drum struct{
	Note, V string
}

func (d Drum) Play() {
	c.Write([]byte("noteoff 9 " + d.Note + " 0\n"))
	c.Write([]byte("noteon 9 " + d.Note + " " + d.V + "\n"))
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

var(
kickA= Drum{Note: "36", V: "127"}
kickB= Drum{Note: "35", V: "127"}
snareA= Drum{Note: "38", V: "96"}
snareB= Drum{Note: "40", V: "127"}
snareC= Drum{Note: "40", V: "80"}
floorA= Drum{Note: "41", V: "127"}
floorB= Drum{Note: "43", V: "127"}
tomA= Drum{Note: "45", V: "127"}
tomB= Drum{Note: "47", V: "127"}
tomC= Drum{Note: "48", V: "127"}
hatC= Drum{Note: "42", V: "127"}
hatB= Drum{Note: "44", V: "127"}
hatA= Drum{Note: "46", V: "127"}
crashA= Drum{Note: "49", V: "127"}
crashB= Drum{Note: "57", V: "127"}
rideA= Drum{Note: "51", V: "127"}
rideB= Drum{Note: "53", V: "127"}
rideC= Drum{Note: "56", V: "127"}
rideD= Drum{Note: "59", V: "127"}
miscA= Drum{Note: "52", V: "127"}
miscB= Drum{Note: "37", V: "127"}
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
			case 's', 'k':
				rideA.Play()
			case 'a':
				crashA.Play()
			case 'q':
				floorA.Play()
			case 'g':
				floorB.Play()
			case 'e':
				tomA.Play()
			case 'f', 'p':
				hatB.Play()
			case 'r':
				tomB.Play()
			case 'o': 
			snareC.Play()
			case 'w':
				snareB.Play()
			case 't':
				tomC.Play()
				//hatA.Play()
			case 'b':
				hatB.Play()
			case 'n', 'm':
				kickB.Play()
			case 'd':
				rideD.Play()
			case '1':
				ProgramChange(-1)
			case '2':
				ProgramChange(1)
			}
		}
	}
	start()
}