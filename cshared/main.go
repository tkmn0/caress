package main

import "C"
import (
	"reflect"
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
	r.ApiError = *CreateApiError(err)
	if e != nil {
		encoders = append(encoders, e)
		r.Ptr = unsafe.Pointer(e)
	}
}

//export CreateDecoder
func CreateDecoder(sampleRate uint32, channels uint16, result unsafe.Pointer) {
	r := (*PointerResult)(result)
	d, err := caress.NewDecoder(sampleRate, channels)
	r.ApiError = *CreateApiError(err)
	if d != nil {
		decorders = append(decorders, d)
		r.Ptr = unsafe.Pointer(d)
	}
}

//export Encode
func Encode(ep unsafe.Pointer, input unsafe.Pointer, result unsafe.Pointer) {
	i := (*Array)(input)
	r := (*ArrayResult)(result)
	e := (*caress.Encoder)(ep)
	var pcm []int16
	var buffer []byte
	arrayToSlice(*i, unsafe.Pointer(&pcm))
	arrayToSlice(r.ResultArray, unsafe.Pointer(&buffer))
	l, err := e.Encode(pcm, buffer)
	r.ApiError = *CreateApiError(err)
	r.ResultArray.Length = uint32(l)
}

//export Decode
func Decode(dp unsafe.Pointer, fec bool, input unsafe.Pointer, result unsafe.Pointer) {
	i := (*Array)(input)
	r := (*ArrayResult)(result)
	d := (*caress.Decoder)(dp)
	var buffer []byte
	var pcm []int16
	arrayToSlice(*i, unsafe.Pointer(&buffer))
	arrayToSlice(r.ResultArray, unsafe.Pointer(&pcm))
	l, err := d.Decode(buffer, pcm, fec)
	r.ApiError = *CreateApiError(err)
	r.ResultArray.Length = uint32(l)
}

func arrayToSlice(source Array, dist unsafe.Pointer) {
	slice := (*reflect.SliceHeader)(dist)
	slice.Len = int(source.Length)
	slice.Cap = int(source.Length)
	slice.Data = uintptr(source.Ptr)
}

func main() {}
