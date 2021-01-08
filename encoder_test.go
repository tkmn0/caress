package caress

import (
	"testing"

	"github.com/tkmn0/caress/opus"
)

func TestEncoderInit(t *testing.T) {
	enc, err := NewEncoder(48000, 1, ApplicationVoip)
	if err != nil || enc == nil {
		t.Errorf("Error creating new encoder: %v", err)
	}
	enc, err = NewEncoder(12345, 1, ApplicationVoip)
	if err == nil || enc != nil {
		t.Errorf("Expected error for illegal samplerate 12345")
	}
}

func TestEncoderUnitialized(t *testing.T) {
	var enc Encoder
	_, err := enc.Encode(nil, nil)
	if err != ErrorUnInitialized {
		t.Errorf("Encode Expected \"uninitialized error\" error: %v", err)
	}

	_, err = enc.EncodeFloat(nil, nil)
	if err != ErrorUnInitialized {
		t.Errorf("EncodeFloat Expected \"uninitialized error\" error: %v", err)
	}

	err = enc.Setbitrate(600)
	if err != ErrorUnInitialized {
		t.Errorf("SetBitrate Expected \"uninitialized error\" error: %v", err)
	}

	_, err = enc.GetBitrate()
	if err != ErrorUnInitialized {
		t.Errorf("SetBitrate Expected \"uninitialized error\" error: %v", err)
	}

	err = enc.SetComplexity(10)
	if err != ErrorUnInitialized {
		t.Errorf("SetComplexity Expected \"uninitialized error\" error: %v", err)
	}

	_, err = enc.GetComplexity()
	if err != ErrorUnInitialized {
		t.Errorf("GetComplexity Expected \"uninitialized error\" error: %v", err)
	}

	err = enc.SetSignal(SignalAuto)
	if err != ErrorUnInitialized {
		t.Errorf("SetSignal Expected \"uninitialized error\" error: %v", err)
	}

	_, err = enc.GetSignal()
	if err != ErrorUnInitialized {
		t.Errorf("GetSignal Expected \"uninitialized error\" error: %v", err)
	}

	err = enc.SetInBandFEC(true)
	if err != ErrorUnInitialized {
		t.Errorf("SetInBandFEC Expected \"uninitialized error\" error: %v", err)
	}

	_, err = enc.GetInBandFEC()
	if err != ErrorUnInitialized {
		t.Errorf("GetInBandFEC Expected \"uninitialized error\" error: %v", err)
	}

	err = enc.SetPacketLossPercentage(30)
	if err != ErrorUnInitialized {
		t.Errorf("SetPacketLossPercentage Expected \"uninitialized error\" error: %v", err)
	}

	_, err = enc.GetPacketLossPercentage()
	if err != ErrorUnInitialized {
		t.Errorf("GetPacketLossPercentage Expected \"uninitialized error\" error: %v", err)
	}
}

func TestNoDataSupplied(t *testing.T) {
	enc, err := NewEncoder(48000, 1, ApplicationVoip)
	if err != nil {
		t.Errorf("encoder create error (this is expected to be passed, check your environment)")
	}
	buffer := make([]byte, 1024)

	_, err = enc.Encode(nil, buffer)
	if err != ErrorNoDataSupplied {
		t.Errorf("Expected \"no data supplied error\" error: %v", err)
	}

	_, err = enc.EncodeFloat(nil, buffer)
	if err != ErrorNoDataSupplied {
		t.Errorf("Expected \"no data supplied error\" error: %v", err)
	}
}

func TestNoTargetBuffer(t *testing.T) {
	enc, err := NewEncoder(48000, 1, ApplicationVoip)
	if err != nil {
		t.Errorf("encoder create error (this is expected to be passed, check your environment)")
	}

	pcmInt16 := make([]int16, 1024)
	_, err = enc.Encode(pcmInt16, nil)
	if err != ErrorNoTargetbuffer {
		t.Errorf("Expected \"no target buffer error\" error: %v", err)
	}

	pcmFloat32 := make([]float32, 1024)
	_, err = enc.EncodeFloat(pcmFloat32, nil)
	if err != ErrorNoTargetbuffer {
		t.Errorf("Expected \"no target buffer error\" error: %v", err)
	}
}

func TestInvalidPcmSizeEncode(t *testing.T) {
	enc, _ := NewEncoder(48000, 2, ApplicationVoip)

	pcm := make([]int16, 255)
	buffer := make([]byte, 256)
	_, err := enc.Encode(pcm, buffer)
	if err != ErrorSuppliedDataSize {
		t.Errorf("Expected \"input buffer length must be multiple of channels\" error: %v", err)
	}

	pcmFloat := make([]float32, 255)
	_, err = enc.EncodeFloat(pcmFloat, buffer)
	if err != ErrorSuppliedDataSize {
		t.Errorf("Expected \"input buffer length must be multiple of channels\" error: %v", err)
	}
}

func TestEncoderBitrate(t *testing.T) {
	enc, err := NewEncoder(48000, 1, ApplicationVoip)
	if err != nil {
		t.Errorf("encoder create error (this is expected to be passed, check your environment)")
	}
	var targetBitrate int32 = 1000
	err = enc.Setbitrate(targetBitrate)
	if err != nil && err != ErrorSetBitrate {
		t.Errorf("Expected \"set bitrate error\" error: %v", err)
	}

	bitrate, err := enc.GetBitrate()
	if err != nil {
		t.Errorf("get bit rate error: %v", err)
	}
	if bitrate != targetBitrate {
		t.Errorf("Expected bitrate is %v, but result is %v", targetBitrate, bitrate)
	}
}

func TestEncoderSetInvalidBitrate(t *testing.T) {
	enc, err := NewEncoder(48000, 1, ApplicationVoip)
	if err != nil {
		t.Errorf("encoder create error (this is expected to be passed, check your environment)")
	}

	err = enc.Setbitrate(opus.MinBitrate - 100)
	if err != ErrorSetBitrateInvalidSize {
		t.Errorf("Expected %v error: %v", ErrorSetBitrateInvalidSize, err)
	}

	err = enc.Setbitrate(opus.MaxBitrate + 100)
	if err != ErrorSetBitrateInvalidSize {
		t.Errorf("Expected %v error: %v", ErrorSetBitrateInvalidSize, err)
	}
}

func TestEncoderComplexity(t *testing.T) {
	enc, err := NewEncoder(48000, 1, ApplicationVoip)
	if err != nil {
		t.Errorf("encoder create error (this is expected to be passed, check your environment)")
	}

	var targetComplexity int32 = 7
	err = enc.SetComplexity(targetComplexity)
	if err != nil {
		t.Errorf("encoder set complexity error: %v", err)
	}

	c, err := enc.GetComplexity()
	if err != nil {
		t.Errorf("encoder get complexity error: %v", err)
	}

	if c != targetComplexity {
		t.Errorf("Expected complexity is %v, but result is %v", targetComplexity, c)
	}
}

func TestEncoderSetInvalidComplexity(t *testing.T) {
	enc, err := NewEncoder(48000, 1, ApplicationVoip)
	if err != nil {
		t.Errorf("encoder create error (this is expected to be passed, check your environment)")
	}

	err = enc.SetComplexity(opus.MaxComplexity + 1)
	if err != ErrorSetComplexityInvalidSize {
		t.Errorf("Expected %v error: %v", ErrorSetComplexityInvalidSize, err)
	}

	err = enc.SetComplexity(opus.MinComplexity - 1)
	if err != ErrorSetComplexityInvalidSize {
		t.Errorf("Expected %v error: %v", ErrorSetComplexityInvalidSize, err)
	}
}

func TestEncoderSetSignal(t *testing.T) {
	enc, err := NewEncoder(48000, 1, ApplicationVoip)
	if err != nil {
		t.Errorf("encoder create error (this is expected to be passed, check your environment)")
	}

	err = enc.SetSignal(SignalVoice)
	if err != nil {
		t.Errorf("Encoder set signal error")
	}

	signal, err := enc.GetSignal()
	if err != nil {
		t.Errorf("Encoder get signal error")
	}

	if signal != SignalVoice {
		t.Errorf("Expected signal value is %v, but result is %v", SignalVoice, signal)
	}
}

func TestEncoderSetSignalInvalidValue(t *testing.T) {
	enc, err := NewEncoder(48000, 1, ApplicationVoip)
	if err != nil {
		t.Errorf("encoder create error (this is expected to be passed, check your environment)")
	}

	err = enc.SetSignal(6000)
	if err == nil || err != ErrorSetSignalInvalidValue {
		t.Errorf("Expected error is %v, but error is %v", ErrorSetSignalInvalidValue, err)
	}
}

func TestEncoderSetInBandFEC(t *testing.T) {
	enc, err := NewEncoder(48000, 1, ApplicationVoip)
	if err != nil {
		t.Errorf("encoder create error (this is expected to be passed, check your environment)")
	}

	enabled := true
	err = enc.SetInBandFEC(enabled)
	if err != nil {
		t.Errorf("set inband fec error")
	}
	result, err := enc.GetInBandFEC()
	if err != nil {
		t.Errorf("get inband fec error")
	}
	if enabled != result {
		t.Errorf("set iband fec value is %v, but result is %v", enabled, result)
	}
}

func TestEncoderSetPacketLossPercentage(t *testing.T) {
	enc, err := NewEncoder(48000, 1, ApplicationVoip)
	if err != nil {
		t.Errorf("encoder create error (this is expected to be passed, check your environment)")
	}
	var perc int32 = 70
	err = enc.SetPacketLossPercentage(perc)
	if err != nil {
		t.Errorf("set packet loss percentage error")
	}
	result, err := enc.GetPacketLossPercentage()
	if err != nil {
		t.Errorf("get packet loss percentage error")
	}
	if perc != result {
		t.Errorf("set to %v percentage but result is %v percentage", perc, result)
	}
}

func TestEncoderSetPakcetLossInvalidPercentage(t *testing.T) {
	enc, err := NewEncoder(48000, 1, ApplicationVoip)
	if err != nil {
		t.Errorf("encoder create error (this is expected to be passed, check your environment)")
	}

	err = enc.SetPacketLossPercentage(120)
	if err != ErrorSetPacketLossPercInvalidValue {
		t.Errorf("Expected error is %v, but error is %v", ErrorSetPacketLossPercInvalidValue, err)
	}

	err = enc.SetPacketLossPercentage(-230)
	if err != ErrorSetPacketLossPercInvalidValue {
		t.Errorf("Expected error is %v, but error is %v", ErrorSetPacketLossPercInvalidValue, err)
	}
}
