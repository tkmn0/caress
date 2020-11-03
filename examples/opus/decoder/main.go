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
	fmt.Println("read ogg file and write to wav")

	infile_path := flag.String("infile", "", "wav file to read")
	flag.Parse()

	file, _ := os.Open(*infile_path)
	reader, header, err := ogg.NewOggReaderWith(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	fmt.Println("sample rate:", header.SampleRate)
	fmt.Println("num channels:", header.Channels)

	d, err := caress.NewDecoder(header.SampleRate, uint16(header.Channels))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	allPcm := make([]int16, 1024)
	// read all data
	for {
		buffer, _, err := reader.ParseNextPage()
		if err != nil {
			break
		}

		pcm := make([]int16, 1024)
		n, err := d.Decode(buffer, pcm, false)
		if err != nil {
			fmt.Println(err.Error())
		}
		pcm = pcm[:n]
		allPcm = append(allPcm, pcm...)
	}

	outfile, err := os.OpenFile("../output.wav", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}

	w := wav.NewWriter(outfile, uint32(len(allPcm)), uint16(header.Channels), header.SampleRate, 16)
	if err != nil {
		fmt.Println(err)
	}

	samples := make([]wav.Sample, len(allPcm))
	for i, pcm := range allPcm {
		samples[i].Values[0] = int(pcm)
	}
	w.WriteSamples(samples)
}
