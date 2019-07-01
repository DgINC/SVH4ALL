package faust

import "C"

import (
	"github.com/DgINC/SVH4ALL/src/faust/faust/include/dsp"
)

type faustDspFactory struct {
	ptrDspFactory    llvm_dsp_factory
	shakeyDspFactory string
	nameDspFactory string
}

type fausDspFactoryList {
	
}
