package opus

/*
#cgo pkg-config: opus
#include "opus.h"
#include <stdlib.h>

int
encoder_set_bitrate(OpusEncoder *st, opus_int32 bitrate){
	return opus_encoder_ctl(st, OPUS_SET_BITRATE(bitrate));
}

int
encoder_get_bitrate(OpusEncoder *st, opus_int32 *bitrate){
	return opus_encoder_ctl(st, OPUS_GET_BITRATE(bitrate));
}

int
encoder_set_complexity(OpusEncoder *st, opus_int32 complexity){
	return opus_encoder_ctl(st, OPUS_SET_COMPLEXITY(complexity));
}

int
encoder_get_complexity(OpusEncoder *st, opus_int32 *complexity){
	return opus_encoder_ctl(st, OPUS_GET_COMPLEXITY(complexity));
}

int
encoder_set_signal(OpusEncoder *st, opus_int32 signal){
	return opus_encoder_ctl(st, OPUS_SET_SIGNAL(signal));
}

int
encoder_get_signal(OpusEncoder *st, opus_int32 *signal){
	return opus_encoder_ctl(st, OPUS_GET_SIGNAL(signal));
}

int
encoder_set_inband_fec(OpusEncoder *st, opus_int32 fec){
	return opus_encoder_ctl(st, OPUS_SET_INBAND_FEC(fec));
}

int
encoder_get_inband_fec(OpusEncoder *st, opus_int32 *fec){
	return opus_encoder_ctl(st, OPUS_GET_INBAND_FEC(fec));
}

int
encoder_set_pakcet_loss_perc(OpusEncoder *st, opus_int32 perc){
	return opus_encoder_ctl(st, OPUS_SET_PACKET_LOSS_PERC(perc));
}

int
encoder_get_pakcet_loss_perc(OpusEncoder *st, opus_int32 *perc){
	return opus_encoder_ctl(st, OPUS_GET_PACKET_LOSS_PERC(perc));
}
*/
import "C"
import (
	"unsafe"
)

const (
	MaxBitrate    = 512000
	MinBitrate    = 500
	MaxComplexity = 10
	MinComplexity = 0
	SignalAuto    = 1000
)

// Maximum is 512000, minimum is 500
func EncoderSetBitrate(e *Encoder, bitrate int32) int32 {
	result := C.encoder_set_bitrate((*C.OpusEncoder)(unsafe.Pointer(e)), C.opus_int32(bitrate))
	return int32(result)
}

func EncoderGetBitrate(e *Encoder) (int32, int32) {
	var bitrate C.opus_int32
	result := C.encoder_get_bitrate((*C.OpusEncoder)(unsafe.Pointer(e)), &bitrate)
	return int32(result), int32(bitrate)
}

// The Opus encoder uses its maximum algorithmic complexity setting of 10 by default.
// This means that it does not hesitate to use CPU to give you the best quality encoding at a given bitrate.
// If the CPU usage is too high for the system you are using Opus on, you can try a lower complexity setting.
// The allowed values span from 10 (highest CPU usage and quality) down to 0 (lowest CPU usage and quality).
func EncoderSetComplexity(e *Encoder, complexity int32) int32 {
	result := C.encoder_set_complexity((*C.OpusEncoder)(unsafe.Pointer(e)), C.opus_int32(complexity))
	return int32(result)
}

func EncoderGetComplexity(e *Encoder) (int32, int32) {
	var complexity C.opus_int32
	result := C.encoder_get_complexity((*C.OpusEncoder)(unsafe.Pointer(e)), &complexity)
	return int32(result), int32(complexity)
}

func EncoderSetSignal(e *Encoder, signal int32) int32 {
	result := C.encoder_set_signal((*C.OpusEncoder)(unsafe.Pointer(e)), C.opus_int32(signal))
	return int32(result)
}

func EncoderGetSignal(e *Encoder) (int32, int32) {
	var signal C.opus_int32
	result := C.encoder_get_signal((*C.OpusEncoder)(unsafe.Pointer(e)), &signal)
	return int32(result), int32(signal)
}

func EncoderSetInBandFEC(e *Encoder, fec bool) int32 {
	var v C.opus_int32 = 0
	if fec {
		v = 1
	}
	result := C.encoder_set_inband_fec((*C.OpusEncoder)(unsafe.Pointer(e)), v)
	return int32(result)
}

func EncoderGetInBandFEC(e *Encoder) (int32, bool) {
	var fec C.opus_int32
	result := C.encoder_get_inband_fec((*C.OpusEncoder)(unsafe.Pointer(e)), &fec)
	v := false
	if fec == 1 {
		v = true
	}
	return int32(result), v
}

func EncoderSetPacketLossPerc(e *Encoder, perc int32) int32 {
	result := C.encoder_set_pakcet_loss_perc((*C.OpusEncoder)(unsafe.Pointer(e)), C.opus_int32(perc))
	return int32(result)
}

func EncoderGetPacketLossPerc(e *Encoder) (int32, int32) {
	var perc C.opus_int32 = 0
	result := C.encoder_get_pakcet_loss_perc((*C.OpusEncoder)(unsafe.Pointer(e)), &perc)
	return int32(result), int32(perc)
}
