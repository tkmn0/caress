package caress

import (
	"github.com/tkmn0/caress/opus"
)

type Encoder struct {
	encoder     *opus.Encoder
	numChannels int
}

func NewEncoder(sampleRate uint32, channels uint16, application int32) (*Encoder, error) {
	var err int32
	e := opus.EncoderCreate(int32(sampleRate), int32(channels), application, &err)

	if err != opus.Ok {
		return nil, ErrorInitialize
	} else {
		return &Encoder{
			encoder:     e,
			numChannels: int(channels),
		}, nil
	}
}

func (e *Encoder) Encode(pcm []int16, buffer []byte) (int, error) {
	if e.encoder == nil {
		return 0, ErrorUnInitialized
	}
	if len(pcm) == 0 {
		return 0, ErrorNoDataSupplied
	}
	if len(buffer) == 0 {
		return 0, ErrorNoTargetbuffer
	}

	// libopus talks about samples as 1 sample containing multiple channels. So
	// e.g. 20 samples of 2-channel data is actually 40 raw data points.
	if len(pcm)%e.numChannels != 0 {
		return 0, ErrorSuppliedDataSize
	}

	samples := len(pcm) / e.numChannels
	n := opus.Encode(e.encoder, pcm, int32(samples), buffer, int32(cap(buffer)))

	if n < 0 {
		return 0, ErrorEncode
	}

	return int(n), nil
}

func (e *Encoder) EncodeFloat(pcm []float32, buffer []byte) (int, error) {
	if e.encoder == nil {
		return 0, ErrorUnInitialized
	}
	if len(pcm) == 0 {
		return 0, ErrorNoDataSupplied
	}
	if len(buffer) == 0 {
		return 0, ErrorNoTargetbuffer
	}

	// libopus talks about samples as 1 sample containing multiple channels. So
	// e.g. 20 samples of 2-channel data is actually 40 raw data points.
	if len(pcm)%e.numChannels != 0 {
		return 0, ErrorSuppliedDataSize
	}

	samples := len(pcm) / e.numChannels
	n := opus.EncodeFloat(e.encoder, pcm, int32(samples), buffer, int32(cap(buffer)))

	if n < 0 {
		return 0, ErrorEncode
	}

	return int(n), nil
}

func (e *Encoder) Setbitrate(bitrate int32) error {
	if e.encoder == nil {
		return ErrorUnInitialized
	}

	if bitrate > opus.MaxBitrate {
		return ErrorSetBitrateInvalidSize
	}

	if bitrate < opus.MinBitrate {
		return ErrorSetBitrateInvalidSize
	}

	result := opus.EncoderSetBitrate(e.encoder, bitrate)
	if result != opus.Ok {
		return ErrorSetBitrate
	}
	return nil
}

func (e *Encoder) GetBitrate() (int32, error) {
	if e.encoder == nil {
		return 0, ErrorUnInitialized
	}
	result, bitrate := opus.EncoderGetBitrate(e.encoder)
	if result != opus.Ok {
		return 0, ErrorGetBitrate
	}
	return bitrate, nil
}

// if cpu usage is so high, you can make complexity value low.
// complexity value range is 0 to 10
func (e *Encoder) SetComplexity(complexity int32) error {
	if e.encoder == nil {
		return ErrorUnInitialized
	}
	if complexity > opus.MaxComplexity {
		return ErrorSetComplexityInvalidSize
	}
	if complexity < opus.MinComplexity {
		return ErrorSetComplexityInvalidSize
	}
	result := opus.EncoderSetComplexity(e.encoder, complexity)
	if result != opus.Ok {
		return ErrorSetComplexity
	}
	return nil
}

func (e *Encoder) GetComplexity() (int32, error) {
	if e.encoder == nil {
		return 0, ErrorUnInitialized
	}
	result, complexity := opus.EncoderGetComplexity(e.encoder)
	if result != opus.Ok {
		return 0, ErrorGetComplexity
	}
	return complexity, nil
}

func (e *Encoder) SetSignal(signal int32) error {
	if e.encoder == nil {
		return ErrorUnInitialized
	}
	if !(signal == opus.SignalAuto || signal == opus.SignalMusic || signal == opus.SignalVoice) {
		return ErrorSetSignalInvalidValue
	}
	result := opus.EncoderSetSignal(e.encoder, signal)
	if result != opus.Ok {
		return ErrorSetSignal
	}
	return nil
}

func (e *Encoder) GetSignal() (int32, error) {
	if e.encoder == nil {
		return 0, ErrorUnInitialized
	}
	result, signal := opus.EncoderGetSignal(e.encoder)
	if result != opus.Ok {
		return 0, ErrorGetSignal
	}
	return signal, nil
}

// this is for "Forwad error Correction". this config could be useful for udp networking.
// see https://ddanilov.me/how-to-enable-in-band-fec-for-opus-codec/
// default value is false
func (e *Encoder) SetInBandFEC(enable bool) error {
	if e.encoder == nil {
		return ErrorUnInitialized
	}
	result := opus.EncoderSetInBandFEC(e.encoder, enable)
	if result != opus.Ok {
		return ErrorSetInBandFEC
	}
	return nil
}

func (e *Encoder) GetInBandFEC() (bool, error) {
	if e.encoder == nil {
		return false, ErrorUnInitialized
	}
	result, enabled := opus.EncoderGetInBandFEC(e.encoder)
	if result != opus.Ok {
		return false, ErrorGetInBandFEC
	}
	return enabled, nil
}

// default is 0
// percentage range is 0 to 100
func (e *Encoder) SetPacketLossPercentage(perc int32) error {
	if e.encoder == nil {
		return ErrorUnInitialized
	}
	if perc > 100 || perc < 0 {
		return ErrorSetPacketLossPercInvalidValue
	}
	result := opus.EncoderSetPacketLossPerc(e.encoder, perc)
	if result != opus.Ok {
		return ErrorSetPacketLossPerc
	}
	return nil
}

func (e *Encoder) GetPacketLossPercentage() (int32, error) {
	if e.encoder == nil {
		return 0, ErrorUnInitialized
	}
	result, perc := opus.EncoderGetPacketLossPerc(e.encoder)
	if result != opus.Ok {
		return 0, ErrorGetPacketLossPerc
	}
	return perc, nil
}
