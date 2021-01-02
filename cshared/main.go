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
func CreateNoiseReducer(config unsafe.Pointer) {
	c := (*NoiseReducerConfig)(config)
	nr := caress.NewNoiseReducer(int(c.NumChannels), uint32(c.SampleRate), c.Attenuation, caress.RnnoiseModel(c.Model.StringValue()))
	noiseReducers = append(noiseReducers, nr)
}

//export CreateEncoder
func CreateEncoder(config unsafe.Pointer, result unsafe.Pointer) {
	c := (*EncoderConfig)(config)
	r := (*PointerResult)(result)
	e, err := caress.NewEncoder(c.SampleRate, c.Channels, c.Application)
	r.ApiError = *CreateApiError(err)
	if e != nil {
		encoders = append(encoders, e)
		r.Ptr = unsafe.Pointer(e)
	}
}

//export CreateDecoder
func CreateDecoder(config unsafe.Pointer, result unsafe.Pointer) {
	c := (*DecoderConfig)(config)
	r := (*PointerResult)(result)
	d, err := caress.NewDecoder(c.SampleRate, c.Channels)
	r.ApiError = *CreateApiError(err)
	if d != nil {
		decorders = append(decorders, d)
		r.Ptr = unsafe.Pointer(d)
	}
}

//export Encode
func Encode(ep unsafe.Pointer, input unsafe.Pointer, result unsafe.Pointer) {
	i := (*Data)(input)
	r := (*DataResult)(result)
	e := (*caress.Encoder)(ep)
	var pcm []int16
	var buffer []byte
	DataToSlice(*i, unsafe.Pointer(&pcm))
	DataToSlice(r.ResultData, unsafe.Pointer(&buffer))
	l, err := e.Encode(pcm, buffer)
	r.ApiError = *CreateApiError(err)
	r.ResultData.Length = uint32(l)
}

//export EncodeFloat
func EncodeFloat(ep unsafe.Pointer, input unsafe.Pointer, result unsafe.Pointer) {
	i := (*Data)(input)
	r := (*DataResult)(result)
	e := (*caress.Encoder)(ep)
	var pcm []float32
	var buffer []byte
	DataToSlice(*i, unsafe.Pointer(&pcm))
	DataToSlice(r.ResultData, unsafe.Pointer(&buffer))
	l, err := e.EncodeFloat(pcm, buffer)
	r.ApiError = *CreateApiError(err)
	r.ResultData.Length = uint32(l)
}

//export Decode
func Decode(dp unsafe.Pointer, fec bool, input unsafe.Pointer, result unsafe.Pointer) {
	i := (*Data)(input)
	r := (*DataResult)(result)
	d := (*caress.Decoder)(dp)
	var buffer []byte
	var pcm []int16
	DataToSlice(*i, unsafe.Pointer(&buffer))
	DataToSlice(r.ResultData, unsafe.Pointer(&pcm))
	l, err := d.Decode(buffer, pcm, fec)
	r.ApiError = *CreateApiError(err)
	r.ResultData.Length = uint32(l)
}

//export DecodeFloat
func DecodeFloat(dp unsafe.Pointer, fec bool, input unsafe.Pointer, result unsafe.Pointer) {
	i := (*Data)(input)
	r := (*DataResult)(result)
	d := (*caress.Decoder)(dp)
	var buffer []byte
	var pcm []float32
	DataToSlice(*i, unsafe.Pointer(&buffer))
	DataToSlice(r.ResultData, unsafe.Pointer(&pcm))
	l, err := d.DecodeFloat(buffer, pcm, fec)
	r.ApiError = *CreateApiError(err)
	r.ResultData.Length = uint32(l)
}

func DataToSlice(source Data, dist unsafe.Pointer) {
	slice := (*reflect.SliceHeader)(dist)
	slice.Len = int(source.Length)
	slice.Cap = int(source.Length)
	slice.Data = uintptr(source.Ptr)
}

func main() {}
