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
	nr := caress.NewNoiseReducer(int(c.NumChannels), uint32(c.SampleRate), c.Attenuation, GetRnnoiseModelName(RnnoiseModelCode(c.Model)))
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

//export ReduceNoise
func ReduceNoise(
	ptr unsafe.Pointer,
	pcm unsafe.Pointer,
	pcmLen int32,
	channel int32) {
	rn := (*caress.NoiseReducer)(ptr)
	var p []int16
	pointerToSlice(pcm, pcmLen, unsafe.Pointer(&p))
	rn.ProcessFrame(p, int(channel))
}

//export ReduceNoiseFloat
func ReduceNoiseFloat(
	ptr unsafe.Pointer,
	pcm unsafe.Pointer,
	pcmLen int32,
	channel int32) {
	rn := (*caress.NoiseReducer)(ptr)
	var p []float32
	pointerToSlice(pcm, pcmLen, unsafe.Pointer(&p))
	rn.ProcessFrameFloat(p, int(channel))
}

//export Encode
func Encode(
	ptr unsafe.Pointer,
	pcm unsafe.Pointer,
	pcmLen int32,
	buffer unsafe.Pointer,
	bufferLen int32,
	result unsafe.Pointer) {
	r := (*IntResult)(result)
	e := (*caress.Encoder)(ptr)
	var p []int16
	var b []byte
	pointerToSlice(pcm, pcmLen, unsafe.Pointer(&p))
	pointerToSlice(buffer, bufferLen, unsafe.Pointer(&b))
	l, err := e.Encode(p, b)
	r.Value = int32(l)
	r.ApiError = *CreateApiError(err)
}

//export EncodeFloat
func EncodeFloat(
	ptr unsafe.Pointer,
	pcm unsafe.Pointer,
	pcmLen int32,
	buffer unsafe.Pointer,
	bufferLen int32,
	result unsafe.Pointer) {
	r := (*IntResult)(result)
	e := (*caress.Encoder)(ptr)
	var p []float32
	var b []byte
	pointerToSlice(pcm, pcmLen, unsafe.Pointer(&p))
	pointerToSlice(buffer, bufferLen, unsafe.Pointer(&b))
	l, err := e.EncodeFloat(p, b)
	r.Value = int32(l)
	r.ApiError = *CreateApiError(err)
}

//export EncoderSetBitrate
func EncoderSetBitrate(ptr unsafe.Pointer, br int32, result unsafe.Pointer) {
	e := (*caress.Encoder)(ptr)
	r := (*ApiError)(result)
	err := e.Setbitrate(br)
	apiErr := CreateApiError(err)
	r.Code = apiErr.Code
	r.Data = apiErr.Data
}

//export EncoderGetBitrate
func EncoderGetBitrate(ptr unsafe.Pointer, result unsafe.Pointer) {
	e := (*caress.Encoder)(ptr)
	r := (*IntResult)(result)
	br, err := e.GetBitrate()
	r.Value = br
	r.ApiError = *CreateApiError(err)
}

//export EncoderSetComplexity
func EncoderSetComplexity(ptr unsafe.Pointer, comp int32, result unsafe.Pointer) {
	e := (*caress.Encoder)(ptr)
	r := (*ApiError)(result)
	err := e.SetComplexity(comp)
	apiErr := CreateApiError(err)
	r.Code = apiErr.Code
	r.Data = apiErr.Data
}

//export EncoderGetComplexity
func EncoderGetComplexity(ptr unsafe.Pointer, result unsafe.Pointer) {
	e := (*caress.Encoder)(ptr)
	r := (*IntResult)(result)
	comp, err := e.GetComplexity()
	r.Value = comp
	r.ApiError = *CreateApiError(err)
}

//export EncoderSetSignal
func EncoderSetSignal(ptr unsafe.Pointer, signal int32, result unsafe.Pointer) {
	e := (*caress.Encoder)(ptr)
	r := (*ApiError)(result)
	err := e.SetSignal(signal)
	apiErr := CreateApiError(err)
	r.Code = apiErr.Code
	r.Data = apiErr.Data
}

//export EncoderGetSignal
func EncoderGetSignal(ptr unsafe.Pointer, result unsafe.Pointer) {
	e := (*caress.Encoder)(ptr)
	r := (*IntResult)(result)
	signal, err := e.GetSignal()
	r.Value = signal
	r.ApiError = *CreateApiError(err)
}

//export EncoderSetInBandFEC
func EncoderSetInBandFEC(ptr unsafe.Pointer, enable bool, result unsafe.Pointer) {
	e := (*caress.Encoder)(ptr)
	r := (*ApiError)(result)
	err := e.SetInBandFEC(enable)
	apiErr := CreateApiError(err)
	r.Code = apiErr.Code
	r.Data = apiErr.Data
}

//export EncoderGetInBandFEC
func EncoderGetInBandFEC(ptr unsafe.Pointer, result unsafe.Pointer) {
	e := (*caress.Encoder)(ptr)
	r := (*BoolResult)(result)
	enabled, err := e.GetInBandFEC()
	r.Value = enabled
	r.ApiError = *CreateApiError(err)
}

//export EncoderSetPacketLossPercentage
func EncoderSetPacketLossPercentage(ptr unsafe.Pointer, percentage int32, result unsafe.Pointer) {
	e := (*caress.Encoder)(ptr)
	r := (*ApiError)(result)
	err := e.SetPacketLossPercentage(percentage)
	apiErr := CreateApiError(err)
	r.Code = apiErr.Code
	r.Data = apiErr.Data
}

//export EncoderGetPacketLossPercentage
func EncoderGetPacketLossPercentage(ptr unsafe.Pointer, result unsafe.Pointer) {
	e := (*caress.Encoder)(ptr)
	r := (*IntResult)(result)
	percentage, err := e.GetPacketLossPercentage()
	r.Value = percentage
	r.ApiError = *CreateApiError(err)
}

//export Decode
func Decode(
	ptr unsafe.Pointer,
	fec bool,
	buffer unsafe.Pointer,
	bufferLen int32,
	pcm unsafe.Pointer,
	pcmLen int32,
	result unsafe.Pointer) {
	r := (*IntResult)(result)
	d := (*caress.Decoder)(ptr)
	var b []byte
	var p []int16
	pointerToSlice(buffer, bufferLen, unsafe.Pointer(&b))
	pointerToSlice(pcm, pcmLen, unsafe.Pointer(&p))
	l, err := d.Decode(b, p, fec)
	r.Value = int32(l)
	r.ApiError = *CreateApiError(err)
}

//export DecodeFloat
func DecodeFloat(
	ptr unsafe.Pointer,
	fec bool,
	buffer unsafe.Pointer,
	bufferLen int32,
	pcm unsafe.Pointer,
	pcmLen int32,
	result unsafe.Pointer) {
	r := (*IntResult)(result)
	d := (*caress.Decoder)(ptr)
	var b []byte
	var p []float32
	pointerToSlice(buffer, bufferLen, unsafe.Pointer(&b))
	pointerToSlice(pcm, pcmLen, unsafe.Pointer(&p))
	l, err := d.DecodeFloat(b, p, fec)
	r.Value = int32(l)
	r.ApiError = *CreateApiError(err)
}

//export SetMaxAttenuation
func SetMaxAttenuation(ptr unsafe.Pointer, maxAttenuationDB float64) {
	rn := (*caress.NoiseReducer)(ptr)
	rn.SetAttenuationDB(maxAttenuationDB)
}

//export ChangeRnnModel
func ChangeRnnModel(ptr unsafe.Pointer, modelCode byte) {
	rn := (*caress.NoiseReducer)(ptr)
	rn.ChangeRnnModel(GetRnnoiseModelName(RnnoiseModelCode(modelCode)))
}

//export DestroyNoiseReducer
func DestroyNoiseReducer(ptr unsafe.Pointer) {
	d := (*Data)(ptr)
	rn := (*caress.NoiseReducer)(d.Ptr)
	rn.Destroy()
	delete(noiseReducers, d.Ptr)
	rn = nil
	d.Ptr = nil
}

//export DestroyEncoder
func DestroyEncoder(ptr unsafe.Pointer) {
	d := (*Data)(ptr)
	enc := (*caress.Encoder)(d.Ptr)
	enc.Destroy()
	delete(encoders, d.Ptr)
	enc = nil
	d.Ptr = nil
}

//export DestroyDecoder
func DestroyDecoder(ptr unsafe.Pointer) {
	d := (*Data)(ptr)
	dec := (*caress.Decoder)(d.Ptr)
	dec.Destroy()
	delete(decoders, d.Ptr)
	dec = nil
	d.Ptr = nil
}

func pointerToSlice(source unsafe.Pointer, length int32, dist unsafe.Pointer) {
	slice := (*reflect.SliceHeader)(dist)
	slice.Len = int(length)
	slice.Cap = int(length)
	slice.Data = uintptr(source)
}

func init() {
	noiseReducers = map[unsafe.Pointer]*caress.NoiseReducer{}
	encoders = map[unsafe.Pointer]*caress.Encoder{}
	decoders = map[unsafe.Pointer]*caress.Decoder{}
}

func main() {}
