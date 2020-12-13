FluidDrums
========

A simple command line application to play drums with [fluidsynth] (http://www.fluidsynth.org/).

Installation
====
You will need to install fluidsynth and a soundfont.

	go get -v -u github.com/insomnimus/fluiddrums

Usage
====

First start fluidsynth as a server:
	fluidsynth -s '/path/to/your/soundfont.sf2'

By default, fluidsynth runs on tcp port 9800 with -s flag.
Now you can just do
	go run main.go

Have fun.