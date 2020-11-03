package main

import "github.com/tkmn0/caress"

var noiseReducers []*caress.NoiseReducer

//export CreateNoiseReducer
func CreateNoiseReducer(numChannels int32, sampleRate int32, attenuation float64, model string) {
	nr := caress.NewNoiseReducer(int(numChannels), uint32(sampleRate), attenuation, caress.RnnoiseModel(model))
	noiseReducers = append(noiseReducers, nr)
}

func main() {}
