package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tkmn0/caress"
	"github.com/tkmn0/caress/examples/opus/util/ogg"
	"github.com/youpy/go-wav"
)

func main() {
	infile_path := flag.String("infile", "", "wav file to read")
	flag.Parse()

	file, _ := os.Open(*infile_path)
	reader := wav.NewReader(file)
	defer file.Close()

	format, err := reader.Format()
	if err != nil {
		fmt.Println(err.Error())
	}

	var numChannels uint16 = format.NumChannels
	var sampleRate uint32 = format.SampleRate

	w, err := ogg.NewOggWriter("../output.ogg", sampleRate, numChannels)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer w.Close()

	e, err := caress.NewEncoder(sampleRate, numChannels, caress.ApplicationVoip)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	samples := make([]wav.Sample, 1024)
	frameSize := 480

	// once read all sampels
	for {
		read, err := reader.ReadSamples(uint32(frameSize))
		samples = append(samples, read...)
		if err != nil {
			break
		}
	}

	// resampling => let frame size to 480
	for i := 0; i < len(samples); i += frameSize {
		resampled := samples[i : i+frameSize]

		pcm := make([]int16, frameSize)

		for i, sample := range resampled {
			pcm[i] = int16(sample.Values[0])
		}

		buffer := make([]byte, 1024)
		// encoding
		n, err := e.Encode(pcm, buffer)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// write encoded data to ogg
		buffer = buffer[:n]
		w.WriteData(buffer)
	}
}
