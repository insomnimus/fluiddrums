// +build windows
package main

import (
	"fmt"
	"github.com/kbinani/win"
	"github.com/nsf/termbox-go"
)

type Note struct {
	on, off win.DWORD
}

var (
	crashB = Note{0x007F3999, 0x00003999}
	rideB  = Note{0x007F3599, 0x00003599}
	rideD  = Note{0x007F3B99, 0x00003B99}
	miscB  = Note{0x007F2599, 0x00002599}
	kickA  = Note{0x007F2499, 0x00002499}
	tomC   = Note{0x007F3099, 0x00003099}
	hatA   = Note{0x007F2E99, 0x00002E99}
	kickB  = Note{0x007F2399, 0x00002399}
	snareB = Note{0x007F2899, 0x00002899}
	tomB   = Note{0x007F2F99, 0x00002F99}
	hatB   = Note{0x007F2C99, 0x00002C99}
	rideC  = Note{0x007F3899, 0x00003899}
	miscA  = Note{0x007F3499, 0x00003499}
	floorB = Note{0x007F2B99, 0x00002B99}
	tomA   = Note{0x007F2D99, 0x00002D99}
	floorA = Note{0x007F2999, 0x00002999}
	hatC   = Note{0x007F2A99, 0x00002A99}
	crashA = Note{0x007F3199, 0x00003199}
	rideA  = Note{0x007F3399, 0x00003399}
	snareA = Note{0x007F2699, 0x00002699}
	snareC = Note{0x007F2899, 0x00002899}
)

func (n Note) Play() {
	win.MidiOutShortMsg(mout, n.off)
	win.MidiOutShortMsg(mout, n.on)
}

var mout win.HMIDIOUT

func main() {
	win.MidiOutOpen(&mout, 0, nil, nil, 0)
	defer win.MidiOutClose(mout)
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
			case 'j', 'k':
				hatB.Play()
			case 'l':
				hatC.Play()
			case 'ı', 'i':
				hatA.Play()
			case 'd':
				kickB.Play()
			//hatB.Play()
			case 'f':
				kickB.Play()
			case 'a':
				snareA.Play()
			case 'q':
				floorA.Play()
			case 'g':
				floorB.Play()
			case 'e':
				tomA.Play()
			case 'p':
				crashA.Play()
			case 'r':
				tomB.Play()
			case 'o', 's':
				snareC.Play()
			case 'w':
				snareB.Play()
			case 't':
				tomC.Play()
				//hatA.Play()
			case 'b':
				hatB.Play()
			case 'n', 'm':
				kickA.Play()
			case 'ş':
				rideD.Play()
			}

		}
	}
	start()
}
