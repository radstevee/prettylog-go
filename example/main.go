package main

import (
	e "errors"

	"github.com/go-errors/errors"
	. "github.com/radstevee/prettylog"
)

var Death = errors.Errorf("Death")
var StdlibError = e.New("Death from the stdlib")

func main() {
	LoggerSettings.LoggerStyle = FULL

	Log("Running main() in Demo.go..", Debug)
	Log("Very informative information", Information)
	Log("I am running on time!", Runtime)
	Log("Downloading maxwell.mp3", Network)
	Log("maxwell.mp3 has been downloaded!", Success)
	Log("Warning.. file maxwell.mp3 may be corrupted!", Warning)
	Log("maxwell.mp3 cannot be played using IDrawableTrack", Error)
	Log("Critical issue detected in the payment system!", Critical)
	Log("User SkibidyToilet727 accessed the admin panel", Audit)
	Log("Entering detailed trace mode for debugging", Trace)
	Log("Security breach attempt detected!", Security)
	Log("NeuroSama updated her profile picture to bread.png", UserAction)
	Log("Response time is 250ms", Performance)
	Log("MaxConnections set to 1000", Config)
	Log("Your life will be terminated", Fatal)
	LogException(errors.New(Death))
	LogException(errors.New(StdlibError))
}
