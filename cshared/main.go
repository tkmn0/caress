package main

import "C"
import (
	"reflect"
	"unsafe"

	"github.com/tkmn0/caress"
)

var noiseReducers map[unsafe.Pointer]*caress.NoiseReducer
var encoders map[unsafe.Pointer]*caress.Encoder
var decoders map[unsafe.Pointer]*caress.Decoder

//export CreateNoiseReducer
func CreateNoiseReducer(config unsafe.Pointer, result unsafe.Pointer) {
	c := (*NoiseReducerConfig)(config)
	r := (*PointerResult)(result)
	nr := caress.NewNoiseReducer(int(c.NumChannels), uint32(c.SampleRate), c.Attenuation, caress.RnnoiseModel(c.Model.StringValue()))
	ptr := unsafe.Pointer(nr)
	noiseReducers[ptr] = nr
	r.Ptr = ptr
	r.ApiError = *CreateApiError(nil)
}

//export CreateEncoder
func CreateEncoder(config unsafe.Pointer, result unsafe.Pointer) {
	c := (*EncoderConfig)(config)
	r := (*PointerResult)(result)
	e, err := caress.NewEncoder(c.SampleRate, c.Channels, c.Application)
	r.ApiError = *CreateApiError(err)
	if e != nil {
		ptr := unsafe.Pointer(e)
		encoders[ptr] = e
		r.Ptr = ptr
	}
}

//export CreateDecoder
func CreateDecoder(config unsafe.Pointer, result unsafe.Pointer) {
	c := (*DecoderConfig)(config)
	r := (*PointerResult)(result)
	d, err := caress.NewDecoder(c.SampleRate, c.Channels)
	r.ApiError = *CreateApiError(err)
	if d != nil {
		ptr := unsafe.Pointer(d)
		decoders[ptr] = d
		r.Ptr = ptr
	}
}

//export Encode
func Encode(ep unsafe.Pointer, input unsafe.Pointer, result unsafe.Pointer) {
	i := (*Data)(input)
	r := (*DataResult)(result)
	e := (*caress.Encoder)(ep)
	var pcm []int16
	var buffer []byte
	dataToSlice(*i, unsafe.Pointer(&pcm))
	dataToSlice(r.ResultData, unsafe.Pointer(&buffer))
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
	dataToSlice(*i, unsafe.Pointer(&pcm))
	dataToSlice(r.ResultData, unsafe.Pointer(&buffer))
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
	dataToSlice(*i, unsafe.Pointer(&buffer))
	dataToSlice(r.ResultData, unsafe.Pointer(&pcm))
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
	dataToSlice(*i, unsafe.Pointer(&buffer))
	dataToSlice(r.ResultData, unsafe.Pointer(&pcm))
	l, err := d.DecodeFloat(buffer, pcm, fec)
	r.ApiError = *CreateApiError(err)
	r.ResultData.Length = uint32(l)
}

func dataToSlice(source Data, dist unsafe.Pointer) {
	slice := (*reflect.SliceHeader)(dist)
	slice.Len = int(source.Length)
	slice.Cap = int(source.Length)
	slice.Data = uintptr(source.Ptr)
}

//export DestroyNoiseReducer
func DestroyNoiseReducer(rp unsafe.Pointer) {
	delete(noiseReducers, rp)
}

//export DestroyEncoder
func DestroyEncoder(ep unsafe.Pointer) {
	delete(encoders, ep)
}

//export DestroyDecoder
func DestroyDecoder(dp unsafe.Pointer) {
	delete(decoders, dp)
}

func init() {
	noiseReducers = map[unsafe.Pointer]*caress.NoiseReducer{}
	encoders = map[unsafe.Pointer]*caress.Encoder{}
	decoders = map[unsafe.Pointer]*caress.Decoder{}
}

func main() {}
