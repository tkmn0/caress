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

type Array struct {
	Ptr    unsafe.Pointer
	Length uint32
}

type ArrayResult struct {
	ResultArray Array
	ApiError    ApiError
}

func CreateApiError(err error) *ApiError {
	if err != nil {
		return &ApiError{
			Code: byte(caressOk),
		}
	} else {
		b := []byte(err.Error())
		return &ApiError{
			Code:          byte(FromErrorToErrorCode(err)),
			Message:       unsafe.Pointer(&b[0]),
			MessageLength: uint32(len(b)),
		}
	}
}
