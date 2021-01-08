package caress

import "testing"

func TestDecoderInit(t *testing.T) {
	dec, err := NewDecoder(48000, 1)
	if err != nil || dec == nil {
		t.Errorf("Error creating new decoder: %v", err)
	}
	dec, err = NewDecoder(12345, 1)
	if err == nil || dec != nil {
		t.Errorf("Expected error for illegal samplerate 12345")
	}
}

func TestDecoderUnitialized(t *testing.T) {
	var dec Decoder
	_, err := dec.Decode(nil, nil, false)
	if err != ErrorUnInitialized {
		t.Errorf("decode Expected \"uninitialized error\" error: %v", err)
	}

	_, err = dec.DecodeFloat(nil, nil, false)
	if err != ErrorUnInitialized {
		t.Errorf("DecodeFloat Expected \"uninitialized error\" error: %v", err)
	}
}

func TestDeocoderNoDataSupplied(t *testing.T) {
	dec, err := NewDecoder(48000, 1)
	if err != nil {
		t.Errorf("decoder create error (this is expected to be passed, check your environment)")
	}
	buffer := make([]int16, 1024)

	_, err = dec.Decode(nil, buffer, false)
	if err != ErrorNoDataSupplied {
		t.Errorf("Expected \"no data supplied error\" error: %v", err)
	}

	bufferFloat := make([]float32, 1024)
	_, err = dec.DecodeFloat(nil, bufferFloat, false)
	if err != ErrorNoDataSupplied {
		t.Errorf("Expected \"no data supplied error\" error: %v", err)
	}
}

func TestDecoderNoTargetBuffer(t *testing.T) {
	dec, err := NewDecoder(48000, 1)
	if err != nil {
		t.Errorf("decoder create error (this is expected to be passed, check your environment)")
	}

	buffer := make([]byte, 1024)
	_, err = dec.Decode(buffer, nil, false)
	if err != ErrorNoTargetbuffer {
		t.Errorf("Expected \"no target buffer error\" error: %v", err)
	}

	_, err = dec.DecodeFloat(buffer, nil, false)
	if err != ErrorNoTargetbuffer {
		t.Errorf("Expected \"no target buffer error\" error: %v", err)
	}
}

func TestDecoderInvalidPcmSizeDecode(t *testing.T) {
	dec, _ := NewDecoder(48000, 2)

	pcm := make([]int16, 255)
	buffer := make([]byte, 256)
	_, err := dec.Decode(buffer, pcm, false)
	if err != ErrorSuppliedDataSize {
		t.Errorf("Expected \"input buffer length must be multiple of channels\" error: %v", err)
	}

	pcmFloat := make([]float32, 255)
	_, err = dec.DecodeFloat(buffer, pcmFloat, false)
	if err != ErrorSuppliedDataSize {
		t.Errorf("Expected \"input buffer length must be multiple of channels\" error: %v", err)
	}
}
