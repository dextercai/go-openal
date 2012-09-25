package main

import (
	"github.com/timshannon/go-openal/al"
	"github.com/timshannon/go-openal/alc"
	"io/ioutil"
)

func main() {
	device := alc.OpenDevice("")
	context := device.CreateContext()
	context.Activate()

	//listener := new(al.Listener)

	source := al.NewSource()
	source.SetPitch(1)
	source.SetGain(1)
	source.SetPosition(al.Vector{0, 0, 0})
	source.SetVelocity(al.Vector{0, 0, 0})
	source.SetLooping(false)

	buffer := al.NewBuffer()

	data, err := ioutil.ReadFile("welcome.wav")
	if err != nil {
		panic(err)
	}

	buffer.SetData(al.FormatMono16, data, 44100)

	source.SetBuffer(buffer)
	source.Play()
	for source.State() == al.Playing {

		//loop long enough to let the wave file finish

	}
	source.Pause()
	context.Destroy()
}
