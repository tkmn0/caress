package caress

import "github.com/tkmn0/caress/opus"

func PcmSoftClip(pcm []float32, frameSize int32, channels int32, softClipMem []float32) {
	opus.PcmSoftClip(pcm, frameSize, channels, softClipMem)
}
