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
func CreateEncoder(sampleRate uint32, channels uint16, application int32, result unsafe.Pointer) {
	r := (*PointerResult)(result)
	e, err := caress.NewEncoder(sampleRate, channels, application)
	if err != nil {
		apiError := CreateApiError(FromErrorToErrorCode(err), err)
		r.Ptr = nil
		r.ApiError = *apiError
	}
	encoders = append(encoders, e)
	r.Ptr = unsafe.Pointer(e)
	r.ApiError.Code = byte(caressOk)
}

//export CreateDecoder
func CreateDecoder(sampleRate uint32, channels uint16, result unsafe.Pointer) {
	r := (*PointerResult)(result)
	d, err := caress.NewDecoder(sampleRate, channels)
	if err != nil {
		apiError := CreateApiError(FromErrorToErrorCode(err), err)
		r.Ptr = nil
		r.ApiError = *apiError
	}
	decorders = append(decorders, d)
	r.Ptr = unsafe.Pointer(d)
	r.ApiError.Code = byte(caressOk)
}

func main() {}
