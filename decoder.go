package caress

import (
	"github.com/tkmn0/caress/opus"
)

type Decoder struct {
	decoder  *opus.Decoder
	channels int
}

func NewDecoder(sampleRate uint32, channels uint16) (*Decoder, error) {
	var err int32
	d := opus.DecoderCreate(int32(sampleRate), int32(channels), &err)

	if err != opus.Ok {
		return nil, errorInitialize
	} else {
		return &Decoder{
			decoder:  d,
			channels: int(channels),
		}, nil
	}
}

func (d *Decoder) Decode(buffer []byte, pcm []int16, fec bool) (int, error) {
	if d.decoder == nil {
		return 0, errorUnInitialized
	}
	if len(buffer) == 0 {
		return 0, errorNoDataSupplied
	}
	if len(pcm) == 0 {
		return 0, errorNoTargetbuffer
	}
	if cap(pcm)%d.channels != 0 {
		return 0, errorSuppliedDataSize
	}

	var v int32 = 0
	if fec {
		v = 1
	}
	n := opus.Decode(d.decoder, data(buffer), int32(len(buffer)), pcm, int32(cap(pcm)/d.channels), v)

	if n < 0 {
		return 0, errorDecode
	}
	return int(n), nil
}

func (d *Decoder) DecodeFloat(buffer []byte, pcm []float32, fec bool) (int, error) {
	if d.decoder == nil {
		return 0, errorUnInitialized
	}
	if len(buffer) == 0 {
		return 0, errorNoDataSupplied
	}
	if len(pcm) == 0 {
		return 0, errorNoTargetbuffer
	}
	if cap(pcm)%d.channels != 0 {
		return 0, errorSuppliedDataSize
	}

	var v int32 = 0
	if fec {
		v = 1
	}
	n := opus.DecodeFloat(d.decoder, data(buffer), int32(len(buffer)), pcm, int32(cap(pcm)/d.channels), v)

	if n < 0 {
		return 0, errorDecode
	}
	return int(n), nil
}
