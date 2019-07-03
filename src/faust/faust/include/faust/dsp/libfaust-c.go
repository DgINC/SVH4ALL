package dsp

/*
#cgo pkg-config: libfaust
#include <stdlib.h>
#include "libfaust-c.h"
*/
import "C"
import "unsafe"

func generateSHA1(data string, sha_key string) {
	cdata := C.CString(data)
	csha_key := C.CString(sha_key)
	defer C.free(unsafe.Pointer(cdata))
	defer C.free(unsafe.Pointer(csha_key))
	_ := C.generateCSHA1(cdata, csha_key)
}
