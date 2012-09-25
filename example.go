package main

import (
	"github.com/timshannon/go-openal/al"
	"github.com/timshannon/go-openal/alc"
)

func main() {
	device := alc.OpenDevice("")
	context := device.CreateContext()
	context.Activate()
}
