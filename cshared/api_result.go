package main

import (
	"unsafe"
)

type ApiError struct {
	Code          byte
	Message       unsafe.Pointer
	MessageLength uint32
}

type PointerResult struct {
	Ptr      unsafe.Pointer
	ApiError ApiError
}

func CreateApiError(errorCode CaressErrorCode, err error) *ApiError {
	b := []byte(err.Error())
	return &ApiError{
		Code:          byte(errorCode),
		Message:       unsafe.Pointer(&b[0]),
		MessageLength: uint32(len(b)),
	}
}
