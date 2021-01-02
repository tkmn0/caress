package caress

import "errors"

// AudioEngine
const (
	// ApplicationVoip as defined in opus/opus_defines.h:193
	ApplicationVoip = 2048
	// ApplicationAudio as defined in opus/opus_defines.h:196
	ApplicationAudio = 2049
	// ApplicationRestrictedLowdelay as defined in opus/opus_defines.h:199
	ApplicationRestrictedLowdelay = 2051

	SignalAuto  = 1000
	SignalVoice = 3001
	SignalMusic = 3002
)

var (
	ErrorInitialize                    = errors.New("initialize error")
	ErrorUnInitialized                 = errors.New("uninitialized error")
	ErrorNoDataSupplied                = errors.New("no data supplied error")
	ErrorNoTargetbuffer                = errors.New("no target buffer error")
	ErrorSuppliedDataSize              = errors.New("input buffer length must be multiple of channels")
	ErrorEncode                        = errors.New("error to encode")
	ErrorDecode                        = errors.New("no data decoded")
	ErrorSetBitrate                    = errors.New("set bitrate error")
	ErrorGetBitrate                    = errors.New("get bitrate error")
	ErrorSetBitrateInvalidSize         = errors.New("set bitrate with invalid size error: capable max bitrate is 512000, min bitrate is 500")
	ErrorSetComplexity                 = errors.New("set complexity error")
	ErrorSetComplexityInvalidSize      = errors.New("set complexity with invalid size error: complexity value range is 0 to 10")
	ErrorGetComplexity                 = errors.New("get complexity error")
	ErrorSetSignal                     = errors.New("set signal error")
	ErrorSetSignalInvalidValue         = errors.New("set signal error with invalid value: signal should be 1000, 3001 or 3002")
	ErrorGetSignal                     = errors.New("get signal error")
	ErrorSetInBandFEC                  = errors.New("set in band fec error")
	ErrorGetInBandFEC                  = errors.New("get in band fec error")
	ErrorSetPacketLossPerc             = errors.New("set packet loss percentage error")
	ErrorSetPacketLossPercInvalidValue = errors.New("set packet loss percentage error: percentage range should be 0 to 100")
	ErrorGetPacketLossPerc             = errors.New("get packet loss percentage error")
)

// NoiseReducer
type RnnoiseModel string

const (
	RnnoiseModelGeneral          = RnnoiseModel("mp")
	RnnoiseModelGeneralRecording = RnnoiseModel("cb")
	RnnoiseModelVoice            = RnnoiseModel("lq")
	RnnoiseModelVoiceRecording   = RnnoiseModel("bd")
	RnnoiseModelSpeech           = RnnoiseModel("orig")
	RnnoiseModelSpeechRecording  = RnnoiseModel("sh")
	RnnoiseModelNone             = "none"
)
