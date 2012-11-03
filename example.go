package main

import (
	"github.com/timshannon/go-openal/openal"
	"io/ioutil"
)

func main() {
	device := openal.OpenDevice("")
	context := device.CreateContext()
	context.Activate()

	//listener := new(openal.Listener)

	source := openal.NewSource()
	source.SetPitch(1)
	source.SetGain(1)
	position := &openal.Vector{0, 0, 0}
	velocity := &openal.Vector{0, 0, 0}
	source.SetPosition(position)
	source.SetVelocity(velocity)
	source.SetLooping(false)

	buffer := openal.NewBuffer()

	data, err := ioutil.ReadFile("welcome.wav")
	if err != nil {
		panic(err)
	}

	buffer.SetData(openal.FormatMono16, data, 44100)

	source.SetBuffer(buffer)
	source.Play()
	for source.State() == openal.Playing {

		//loop long enough to let the wave file finish

	}
	source.Pause()
	context.Destroy()
}
