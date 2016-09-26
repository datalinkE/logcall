package main

import (
	"github.com/datalinkE/logcall/debug"
)

func A() {
	debug.LogCall()
}

func B() {
	debug.LogCall()
	debug.Logf("Formated string from B (with a function name and some arguments %v)", 12345)
	A()
}

func main() {
	debug.Log("Simple string from main, just as with usual log.Print")
	debug.LogCallEnabled = true
	B()
}
