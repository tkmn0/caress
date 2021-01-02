package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"reflect"

	"github.com/tkmn0/caress"
)

func main() {
	infile_path := flag.String("infile", "", "wav file to read")
	flag.Parse()

	file, _ := os.Open(*infile_path)
	defer file.Close()

	outFile, _ := os.Create("./processed.raw")
	defer outFile.Close()

	rnn := caress.NewNoiseReducer(1, 48000, 30, caress.RnnoiseModelVoice)
	defer rnn.Destroy()

	buffer := make([]byte, 480*2)

	for {
		_, err := file.Read(buffer)
		if err != nil {
			break
		}

		pcm := createInt16(buffer)
		original := createInt16(buffer)
		fmt.Println(reflect.DeepEqual(pcm, original))

		rnn.ProcessFrame(pcm, 0)

		fmt.Println(reflect.DeepEqual(pcm, original))
		data := createBytes(pcm)

		outFile.Write(data)
	}
}

func createInt16(data []byte) []int16 {
	rawPcm := make([]int16, len(data)/2)
	for i := 0; i < len(data); i += 2 {
		var value int16
		buffer := bytes.NewReader(data[i : i+2])
		binary.Read(buffer, binary.LittleEndian, &value)
		rawPcm[i/2] = value
	}

	return rawPcm
}

func createBytes(pcm []int16) []byte {
	data := make([]byte, len(pcm)*2)

	for i := 0; i < len(pcm); i++ {
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.LittleEndian, pcm[i])

		data[i*2] = buf.Bytes()[0]
		data[i*2+1] = buf.Bytes()[1]
	}

	return data
}
