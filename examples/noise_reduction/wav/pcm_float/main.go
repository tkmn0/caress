package main

import (
	"flag"
	"fmt"
	"math"
	"os"

	"github.com/tkmn0/caress"
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

	outfile, err := os.OpenFile("./output.wav", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		outfile.Close()
	}()

	fmt.Println("numChannels:", format.NumChannels)
	fmt.Println("sampleRate:", format.SampleRate)
	fmt.Println("bitsPerSample:", format.BitsPerSample)
	fmt.Println("byteRate", format.ByteRate)
	fmt.Println("blockAlign", format.BlockAlign)

	var numChannels uint16 = format.NumChannels
	var sampleRate uint32 = format.SampleRate
	var bitsPerSample uint16 = format.BitsPerSample

	rnn := caress.NewNoiseReducer(int(numChannels), sampleRate, 20, caress.Voice)
	defer rnn.Destroy()

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

	writerInt16 := wav.NewWriter(outfile, uint32(len(samples)), numChannels, sampleRate, bitsPerSample)
	fmt.Println("numSamples:", len(samples))

	// float32
	// resampling => let frame size to 480
	for i := 0; i < len(samples); i += frameSize {
		resampled := samples[i : i+frameSize]
		outSample := make([]wav.Sample, frameSize)

		// each channel
		for c := 0; c < int(numChannels); c++ {
			pcm := make([]float32, frameSize)
			for i, sample := range resampled {
				pcm[i] = float32(reader.FloatValue(sample, uint(c)))
			}

			// rnnnoise
			rnn.ProcessFrameFloat(pcm, c)

			for i := range outSample {
				outSample[i].Values[c] = int(float64(pcm[i]) * math.Pow(2, float64(format.BitsPerSample)-1))
			}
		}
		writerInt16.WriteSamples(outSample)
	}
}
