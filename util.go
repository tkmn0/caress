package caress

import "unsafe"

func data(b []byte) string {
	hdr := (*sliceHeader)(unsafe.Pointer(&b))
	return *(*string)(unsafe.Pointer(&stringHeader{
		Data: hdr.Data,
		Len:  hdr.Len,
	}))
}

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

type stringHeader struct {
	Data uintptr
	Len  int
}
