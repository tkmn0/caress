package caress

/*
#cgo CFLAGS: -I/usr/local/include
#cgo LDFLAGS: -L//usr/local/lib -lrnnoise-nu

#include <stdio.h>
#include <math.h>
#include "rnnoise-nu.h"

DenoiseState *create_rnnoise(){
	return rnnoise_create(NULL);
}

DenoiseState *create_rnnoise_with(const char *name){
	RNNModel *model = NULL;
	model = rnnoise_get_model("lq");

	return rnnoise_create(model);
}

void process_frame_float(DenoiseState *state, int frameSize, float *out, const float *in){
	for (int i = 0; i < frameSize; i++) out[i] *= 0x7fff;
	rnnoise_process_frame(state, out, in);
	for (int i = 0; i < frameSize; i++) out[i] /= 0x7fff;
}

void process_frame(DenoiseState *state, int frameSize, short *out, const short *in){
	float x[frameSize];
	for (int i=0;i<frameSize;i++) x[i] = in[i];
	rnnoise_process_frame(state, x, x);
    for (int i=0;i<frameSize;i++) out[i] = x[i];
}

void setup(DenoiseState *state, double attenuation, int sampleRate){
	int sample_rate = sampleRate;
	float max_attenuation = pow(10, -attenuation/10);
    rnnoise_set_param(state, RNNOISE_PARAM_MAX_ATTENUATION, max_attenuation);
    rnnoise_set_param(state, RNNOISE_PARAM_SAMPLE_RATE, sample_rate);
}
*/
import "C"
import (
	"unsafe"
)

type NoiseReducer struct {
	rnnStates []*C.DenoiseState
}

func NewNoiseReducer(numChannels int, sampleRate uint32, attenuationDB float64, model RnnoiseModel) *NoiseReducer {
	states := []*C.DenoiseState{}
	var state *C.DenoiseState
	for i := 0; i < numChannels; i++ {
		if model == None {
			state = C.create_rnnoise()
		} else {
			state = C.create_rnnoise_with(C.CString(string(model)))
		}
		C.setup(state, C.double(attenuationDB), C.int(sampleRate))
		states = append(states, state)
	}

	return &NoiseReducer{
		rnnStates: states,
	}
}

func (r *NoiseReducer) ProcessFrame(frame []int16, channel int) {
	if len(frame) == 0 {
		return
	}
	C.process_frame(r.rnnStates[channel], C.int(len(frame)), (*C.short)(unsafe.Pointer(&frame[0])), (*C.short)(unsafe.Pointer(&frame[0])))
}

func (r *NoiseReducer) ProcessFrameFloat(frame []float32, channel int) {
	if len(frame) == 0 {
		return
	}
	C.process_frame_float(r.rnnStates[channel], C.int(len(frame)), (*C.float)(&frame[0]), (*C.float)(&frame[0]))
}

func (r *NoiseReducer) Destroy() {
	for _, state := range r.rnnStates {
		C.rnnoise_destroy(state)
	}
}
