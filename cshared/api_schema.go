package main

import "C"
import (
	"unsafe"

	"github.com/tkmn0/caress"
)

type Data struct {
	Ptr    unsafe.Pointer
	Length uint32
}

func (d *Data) StringValue() string {
	return string(d.ByteValue())
}

func (d *Data) ByteValue() []byte {
	return C.GoBytes(d.Ptr, C.int(d.Length))
}

type ApiError struct {
	Code byte
	Data Data
}

type PointerResult struct {
	Ptr      unsafe.Pointer
	ApiError ApiError
}

type EncodeDecodeResult struct {
	Length   int32
	ApiError ApiError
}

type NoiseReducerConfig struct {
	NumChannels int32
	SampleRate  int32
	Attenuation float64
	Model       byte
}

type EncoderConfig struct {
	SampleRate  uint32
	Channels    uint16
	Application int32
}

type DecoderConfig struct {
	SampleRate uint32
	Channels   uint16
}

func CreateApiError(err error) *ApiError {
	if err == nil {
		return &ApiError{
			Code: byte(caressOk),
		}
	} else {
		b := []byte(err.Error())
		return &ApiError{
			Code: byte(FromErrorToErrorCode(err)),
			Data: Data{
				Ptr:    unsafe.Pointer(&b[0]),
				Length: uint32(len(b)),
			},
		}
	}
}

type RnnoiseModelCode byte

const (
	RnnoiseModelGeneral RnnoiseModelCode = iota
	RnnoiseModelGeneralRecording
	RnnoiseModelVoice
	RnnoiseModelVoiceRecording
	RnnoiseModelSpeech
	RnnoiseModelSpeechRecording
	None
)

func GetRnnoiseModelName(code RnnoiseModelCode) caress.RnnoiseModel {
	switch code {
	case RnnoiseModelGeneral:
		return caress.RnnoiseModelGeneral
	case RnnoiseModelGeneralRecording:
		return caress.RnnoiseModelGeneralRecording
	case RnnoiseModelVoice:
		return caress.RnnoiseModelVoice
	case RnnoiseModelVoiceRecording:
		return caress.RnnoiseModelVoiceRecording
	case RnnoiseModelSpeech:
		return caress.RnnoiseModelSpeech
	case RnnoiseModelSpeechRecording:
		return caress.RnnoiseModelSpeechRecording
	case None:
		return caress.RnnoiseModelNone
	default:
		return caress.RnnoiseModelNone
	}
}
