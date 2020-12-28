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
	errorInitialize                    = errors.New("initialize error")
	errorUniInitialized                = errors.New("uninitialized error")
	errorNoDataSupplied                = errors.New("no data supplied error")
	errorNoTargetbuffer                = errors.New("no target buffer error")
	errorSuppliedDataSize              = errors.New("input buffer length must be multiple of channels")
	errorEncode                        = errors.New("error to encode")
	errorDecode                        = errors.New("no data decoded")
	errorSetBitrate                    = errors.New("set bitrate error")
	errorGetBitrate                    = errors.New("get bitrate error")
	errorSetBitrateInvalidSize         = errors.New("set bitrate with invalid size error: capable max bitrate is 512000, min bitrate is 500")
	errorSetComplexity                 = errors.New("set complexity error")
	errorSetComplexityInvalidSize      = errors.New("set complexity with invalid size error: complexity value range is 0 to 10")
	errorGetComplexity                 = errors.New("get complexity error")
	errorSetSignal                     = errors.New("set signal error")
	errorSetSignalInvalidValue         = errors.New("set signal error with invalid value: signal should be 1000, 3001 or 3002")
	errorGetSignal                     = errors.New("get signal error")
	errorSetInBandFEC                  = errors.New("set in band fec error")
	errorGetInBandFEC                  = errors.New("get in band fec error")
	errorSetPacketLossPerc             = errors.New("set packet loss percentage error")
	errorSetPacketLossPercInvalidValue = errors.New("set packet loss percentage error: percentage range should be 0 to 100")
	errorGetPacketLossPerc             = errors.New("get packet loss percentage error")
)

// NoiseReducer
type RnnoiseModel string

const (
	General          = RnnoiseModel("mp")
	GeneralRecording = RnnoiseModel("cb")
	Voice            = RnnoiseModel("lq")
	VoiceRecording   = RnnoiseModel("bd")
	Speech           = RnnoiseModel("orig")
	SpeechRecording  = RnnoiseModel("sh")
	None             = "none"
)
