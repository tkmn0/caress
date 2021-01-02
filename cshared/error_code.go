package main

import "github.com/tkmn0/caress"

type CaressErrorCode byte

const (
	caressOk CaressErrorCode = iota + 1
	errorInitialize
	errorUnInitialized
	errorNoDataSupplied
	errorNoTargetBuffer
	errorSuppliedDataSize
	errorEncode
	errorDecode
	errorSetBitrate
	errorGetBitrate
	errorSetBitrateInvalidSize
	errorSetComplexity
	errorSetComplexityInvalidSize
	errorGetComplexity
	errorSetSignal
	errorSetSignalInvalidValue
	errorGetSignal
	errorSetInBandFEC
	errorGetInBandFEC
	errorSetPacketLossPerc
	errorSetPacketLossPercInvalidValue
	errorGetPacketLossPerc
	errorUnDefined
)

func FromErrorToErrorCode(e error) CaressErrorCode {
	switch e {
	case caress.ErrorInitialize:
		return errorInitialize
	case caress.ErrorUnInitialized:
		return errorUnInitialized
	case caress.ErrorNoDataSupplied:
		return errorNoDataSupplied
	case caress.ErrorNoTargetbuffer:
		return errorNoTargetBuffer
	case caress.ErrorSuppliedDataSize:
		return errorSuppliedDataSize
	case caress.ErrorEncode:
		return errorEncode
	case caress.ErrorDecode:
		return errorDecode
	case caress.ErrorSetBitrate:
		return errorSetBitrate
	case caress.ErrorGetBitrate:
		return errorGetBitrate
	case caress.ErrorSetBitrateInvalidSize:
		return errorSetBitrateInvalidSize
	case caress.ErrorSetComplexity:
		return errorSetComplexity
	case caress.ErrorSetComplexityInvalidSize:
		return errorSetComplexityInvalidSize
	case caress.ErrorGetComplexity:
		return errorGetComplexity
	case caress.ErrorSetSignal:
		return errorSetSignal
	case caress.ErrorSetSignalInvalidValue:
		return errorSetSignalInvalidValue
	case caress.ErrorGetSignal:
		return errorGetSignal
	case caress.ErrorSetInBandFEC:
		return errorSetInBandFEC
	case caress.ErrorGetInBandFEC:
		return errorGetInBandFEC
	case caress.ErrorSetPacketLossPerc:
		return errorSetPacketLossPerc
	case caress.ErrorSetPacketLossPercInvalidValue:
		return errorSetPacketLossPercInvalidValue
	case caress.ErrorGetPacketLossPerc:
		return errorGetPacketLossPerc
	default:
		return errorUnDefined
	}
}
