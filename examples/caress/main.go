package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tkmn0/caress"
	"github.com/tkmn0/caress/examples/opus/util/ogg"
	"github.com/youpy/go-wav"
)

var noise_reduced_wav_path = "./out/0_noise_reduced.wav"
var encoded_ogg_path = "./out/1_noise_reduced_encoded.ogg"
var decoded_wav_path = "./out/2_noise_reduced_encoded_decoded.wav"

func main() {
	fmt.Println("read wav file") // should be samplerate 48000, mono, 16bit pcm
	infile_path := flag.String("infile", "", "wav file to read")
	flag.Parse()

	file, _ := os.Open(*infile_path)
	reader := wav.NewReader(file)
	defer file.Close()

	format, err := reader.Format()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("numChannels:", format.NumChannels)
	fmt.Println("sampleRate:", format.SampleRate)
	fmt.Println("bitsPerSample:", format.BitsPerSample)
	fmt.Println("byteRate", format.ByteRate)
	fmt.Println("blockAlign", format.BlockAlign)

	var numChannels uint16 = format.NumChannels
	var sampleRate uint32 = format.SampleRate
	var bitsPerSample uint16 = format.BitsPerSample

	samples := make([]wav.Sample, 1024)

	// NoiseReducer can reduce with frame size 480
	frameSize := 480

	// once read all sampels
	for {
		read, err := reader.ReadSamples(uint32(frameSize))
		samples = append(samples, read...)
		if err != nil {
			break
		}
	}
	fmt.Println("numSamples:", len(samples))

	// reduce noise
	s, err := reduceNoise(samples, numChannels, sampleRate, bitsPerSample, frameSize)
	if err != nil {
		fmt.Println("noise reduction error:", err)
		return
	}

	// encode sampels and get encoded data chunks
	data, err := encodeSamples(s, sampleRate, numChannels, frameSize)
	if err != nil {
		fmt.Println("encode error:", err)
		return
	}

	// decode encoded data chunks
	err = decodeData(data, sampleRate, numChannels, bitsPerSample, frameSize)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}

	fmt.Println("all process done! you can check outputs in output folder")
}

func reduceNoise(samples []wav.Sample, numChannels uint16, sampleRate uint32, bitsPerSample uint16, frameSize int) ([]wav.Sample, error) {
	outfile, err := os.OpenFile(noise_reduced_wav_path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer func() {
		outfile.Close()
	}()

	writer := wav.NewWriter(outfile, uint32(len(samples)), numChannels, sampleRate, bitsPerSample)
	rnn := caress.NewNoiseReducer(int(numChannels), sampleRate, 20, caress.RnnoiseModelVoice)
	defer rnn.Destroy()

	allResampled := []wav.Sample{}
	// resampling => let frame size to 480
	for i := 0; i < len(samples); i += frameSize {
		resampled := samples[i : i+frameSize]
		outSample := make([]wav.Sample, frameSize)

		// each channel
		for c := 0; c < int(numChannels); c++ {
			pcm := make([]int16, frameSize)
			for i, sample := range resampled {
				pcm[i] = int16(sample.Values[c])
			}

			// rnnnoise
			rnn.ProcessFrame(pcm, c)

			for i := range outSample {
				outSample[i].Values[c] = int(pcm[i])
			}
		}
		writer.WriteSamples(outSample)
		allResampled = append(allResampled, outSample...)
	}

	return allResampled, nil
}

func encodeSamples(samples []wav.Sample, sampleRate uint32, numChannels uint16, frameSize int) ([][]byte, error) {
	e, err := caress.NewEncoder(sampleRate, numChannels, caress.ApplicationVoip)
	if err != nil {
		return nil, err
	}

	w, err := ogg.NewOggWriter(encoded_ogg_path, sampleRate, numChannels)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer w.Close()

	data := [][]byte{}

	// resampling => let frame size to 480
	for i := 0; i < len(samples); i += frameSize {
		resampled := samples[i : i+frameSize]

		pcm := make([]int16, frameSize)

		for i, sample := range resampled {
			pcm[i] = int16(sample.Values[0])
		}

		buffer := make([]byte, 1024*4)
		// encoding
		n, err := e.Encode(pcm, buffer)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			// write encoded data to ogg
			buffer = buffer[:n]
			w.WriteData(buffer)

			data = append(data, buffer)
		}
	}
	return data, nil
}

func decodeData(data [][]byte, sampleRate uint32, numChannels uint16, bitsPerSample uint16, frameSize int) error {
	d, err := caress.NewDecoder(sampleRate, numChannels)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	allPcm := make([]int16, 1024)

	for i := 0; i < len(data); i++ {
		pcm := make([]int16, 1024)
		n, err := d.Decode(data[i], pcm, false)

		if err != nil {
			fmt.Println(err.Error())
		}
		pcm = pcm[:n]
		allPcm = append(allPcm, pcm...)
	}

	outfile, err := os.OpenFile(decoded_wav_path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	w := wav.NewWriter(outfile, uint32(len(allPcm)), numChannels, sampleRate, bitsPerSample)
	if err != nil {
		return err
	}

	samples := make([]wav.Sample, len(allPcm))
	for i, pcm := range allPcm {
		samples[i].Values[0] = int(pcm)
	}
	w.WriteSamples(samples)
	return nil
}
