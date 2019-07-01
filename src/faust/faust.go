package faust

import "C"

type faustDspFactory struct {
	dspFactoryPtr    llvm_dsp_factory
	shakeyDspFactory string
}
