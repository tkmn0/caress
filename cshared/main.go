package main

import (
	"unsafe"

	"github.com/tkmn0/caress"
)

var noiseReducers []*caress.NoiseReducer

//export CreateNoiseReducer
func CreateNoiseReducer(numChannels int32, sampleRate int32, attenuation float64, model string) {
	nr := caress.NewNoiseReducer(int(numChannels), uint32(sampleRate), attenuation, caress.RnnoiseModel(model))
	noiseReducers = append(noiseReducers, nr)
}

//export CreateEncoder
func CreateEncoder(sampleRate uint32, channels uint16, application int32) unsafe.Pointer {
	e, err := caress.NewEncoder(sampleRate, channels, application)
	if err != nil {
		return nil
	}
	return unsafe.Pointer(e)
}

func main() {}
