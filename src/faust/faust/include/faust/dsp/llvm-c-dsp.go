package dsp

/*
#include "llvm-c-dsp.h"
*/
import "C"

import "unsafe"

type llvm_dsp_factory unsafe.Pointer
type llvm_dsp unsafe.Pointer

func getLibFaustVersion() string {
	ver := C.getCLibFaustVersion()
	defer C.free(unsafe.Pointer(ver))
	return C.GoString(ver)
}

func getDSPMachineTarget() string {
	cs := C.getCDSPMachineTarget()
	defer C.free(unsafe.Pointer(cs))
	return C.GoString(cs)
}

func getDSPFactoryFromSHAKey(sha_key string) llvm_dsp_factory {
	cs := C.CString(sha_key)
	defer C.free(unsafe.Pointer(cs))
	pointr * llvm_dsp_factory = C.getCDSPFactoryFromSHAKey(cs)
	return pointr
}

func createDSPFactoryFromFile(filename string, argc int32, argv []string, target string, error_msg string, opt_level int32) llvm_dsp_factory {
	cfilename := C.CString(filename)
	cArray := C.malloc(C.size_t(len(argv)) * C.size_t(unsafe.Sizeof(uintptr(0))))

	cargv := (*[1<<30 - 1]*C.char)(cArray)

	for index, value := range argv {
		cargv[index] = C.CString(value)
		defer C.free(unsafe.Pointer(value))
	}
	defer C.free(unsafe.Pointer(cArray), unsafe.Pointer(cargv))

	ctarget := C.CString(target)
	cerror_msg := C.CString(error_msg)
	defer C.free(unsafe.Pointer(cfilename), unsafe.Pointer(cargc), unsafe.Pointer(cstring), unsafe.Pointer(cerror_msg))
	pointr * llvm_dsp_factory = C.createCDSPFactoryFromFile(cfilename, argc, cargv, ctarget, cerror_msg, opt_level)
	return pointr
}

func createDSPFactoryFromString(name_app string, dsp_content string, argc int32, argv []string, target string, error_msg string, opt_level int32) llvm_dsp_factory {
	cname_app := C.CString(name_app)
	cdsp_content := C.CString(dsp_content)
	cArray := C.malloc(C.size_t(len(argv)) * C.size_t(unsafe.Sizeof(uintptr(0))))

	cargv := (*[1<<30 - 1]*C.char)(cArray)

	for index, value := range argv {
		cargv[index] = C.CString(value)
		defer C.free(unsafe.Pointer(value))
	}
	defer C.free(unsafe.Pointer(cArray), unsafe.Pointer(cargv))

	cerror_msg := C.CString(error_msg)
	defer C.free(unsafe.Pointer(cname_app), unsafe.Pointer(cargc), unsafe.Pointer(cstring), unsafe.Pointer(cerror_msg))
	pointr * llvm_dsp_factory = C.createCDSPFactoryFromString(cname_app, cdsp_content, argc, cargv, ctarget, cerror_msg, opt_level)
	return pointr
}

func deleteCDSPFactory(llvm_dsp_factory *factory) bool {

}
