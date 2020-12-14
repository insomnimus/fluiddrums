package main

import (
	"flag"
	"fmt"
	"github.com/nsf/termbox-go"
	"net"
)

var c net.Conn

type Key struct {
	Note     int
	Octave   int
	Original int
	Text     string
}

func NewKey(note, octave int) *Key {
	return &Key{
		Note:     note,
		Original: note - 11,
		Octave:   octave,
		Text:     fmt.Sprintf("%d", octave*12+note),
	}
}

func (k *Key) Play() {
	c.Write([]byte("noteoff 1 " + k.Text + " 0\n"))
	c.Write([]byte("noteon 1 " + k.Text + " 127\n"))
}

func (k *Key) ChangeOctave() {
	k.Octave = oct
	k.Text = fmt.Sprintf("%d", 12*oct+k.Note)
}

var (
	keyAb  = NewKey(11, 2)
	keyA   = NewKey(12, 2)
	keyBb  = NewKey(13, 2)
	keyB   = NewKey(14, 2)
	keyC   = NewKey(15, 2)
	keyCs  = NewKey(16, 2)
	keyD   = NewKey(17, 2)
	keyEb  = NewKey(18, 2)
	keyE   = NewKey(19, 2)
	keyF   = NewKey(20, 2)
	keyFs  = NewKey(21, 2)
	keyG   = NewKey(22, 2)
	keyGs  = NewKey(23, 2)
	keyA2  = NewKey(24, 2)
	keyBb2 = NewKey(25, 2)
	keyB2  = NewKey(26, 2)
	keyC2  = NewKey(27, 2)
	keyCs2 = NewKey(28, 2)
)

var notes = map[rune]*Key{
	'q': keyAb,
	'a': keyA,
	'w': keyBb,
	's': keyB,
	'd': keyC,
	'r': keyCs,
	'f': keyD,
	't': keyEb,
	'g': keyE,
	'h': keyF,
	'y': keyFs,
	'j': keyG,
	'i': keyGs,
	'ı': keyGs,
	'k': keyA2,
	'o': keyBb2,
	'l': keyB2,
	'ş': keyC2,
	'ğ': keyC2,
	'p': keyCs2,
}

var oct int

func OctaveChange(step int) {
	oct += step
	if oct > 7 {
		oct = 0
	}
	if oct < 0 {
		oct = 7
	}
	for _, val := range notes {
		val.ChangeOctave()
	}
}

var preset = 0

func ProgramChange(step int) {
	preset += step
	if preset < 0 {
		preset = 20
	}
	if preset > 20 {
		preset = 0
	}
	c.Write([]byte(fmt.Sprintf("prog 1 %d\n", preset)))
}

func Release() {
	for i := 0; i < 128; i++ {
		c.Write([]byte(fmt.Sprintf("noteoff 1 %d 0\n", i)))
	}
}

func main() {
	fmt.Println(`
	1-2: change octave
	3-4: change preset
	spacebar: stop all playing notes`)
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
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
		var tempKey *Key
		var ch rune
		for {
			e = termbox.PollEvent()
			if e.Type != termbox.EventKey {
				continue
			}
			if e.Key == termbox.KeySpace {
				Release()
				continue
			}
			ch = e.Ch
			switch ch {
			case '4':
				ProgramChange(1)
			case '3':
				ProgramChange(-1)
			case '2':
				OctaveChange(1)
			case '1':
				OctaveChange(-1)
			case '.', '-', ',':
				return
			default:
				tempKey = notes[ch]
				if tempKey != nil {
					tempKey.Play()
				}
			}
		}
	}
	start()
}
