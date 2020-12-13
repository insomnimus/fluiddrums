/*
Before running this program you first need;
1- fluidsynth
2- a soundfont

start off by running fluidsynth as a service with
`fluidsynth -s '<path/to/the/soundfont.sf2>'`

fluidsynth will start as a server on port 9800 by default, check docs for more info.

in the terminal with fluidsynth running, try typing `noteon 9 40 120`
if you get a snare sound, it works, now run this code.
*/

package main

import (
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

const (
	drumA Drum = "36" // base drum
	drumB Drum = "40" // electric snare
	drumC Drum = "59" // ride symbol2
	drumD Drum = "41" // low floor tom
	drumE Drum = "42" // closed hi-hat
	drumF Drum = "46" // open hi-hat
	drumG Drum = "49" // crash symbol1
	drumH Drum = "44" // pedal hi-hat
	drumI Drum = "48" // high mid tom
	drumJ Drum = "43" // ride bell
	drumK Drum = "48" // high mid tom
)

func main() {
	var port string
	flag.StringVar(&port, "p", "localhost:9800", "fluidsynth port")
	flag.Parse()
	conn, err := net.Dial("tcp", port)
	if err != nil {
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
			drumA.Play()
			drumH.Play()
			case 'j', 'k':
				drumA.Play()
			case 'f':
				drumC.Play()
			case 'e':
				drumB.Play()
			case 'a':
				drumG.Play()
			case 'g':
				drumD.Play()
			case 's':
				drumE.Play()
			case 'l':
				drumF.Play()
			case 'w':
				drumH.Play()
			case 'r':
				drumI.Play()
			case 'd':
				drumK.Play()
			case 'q':
				drumJ.Play()
			}
		}
	}
	start()
}