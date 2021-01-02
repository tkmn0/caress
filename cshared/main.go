package main

import "C"
import (
	"unsafe"

	"github.com/tkmn0/caress"
)

var noiseReducers []*caress.NoiseReducer
var encoders []*caress.Encoder
var decorders []*caress.Decoder

//export Initialize
func Initialize() {
	noiseReducers = []*caress.NoiseReducer{}
	encoders = []*caress.Encoder{}
	decorders = []*caress.Decoder{}
}

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
	encoders = append(encoders, e)
	return unsafe.Pointer(e)
}

//export CreateDecoder
func CreateDecoder(sampleRate uint32, channels uint16) unsafe.Pointer {
	d, err := caress.NewDecoder(sampleRate, channels)
	if err != nil {
		return nil
	}
	decorders = append(decorders, d)
	return unsafe.Pointer(d)
}

func main() {}
