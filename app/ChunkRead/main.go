package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

const DATA_PATH = "../../data/measurements.txt"
const CHAR_LB = '\n'
const DATA_CHAN_BUFSIZE = 1024     // buffers data for goroutines -> 4K
const FILE_CHUNKSIZE = 1024 * 1024 // MB

var wg sync.WaitGroup

func main() {
	f, err := os.Create("profile.pprof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	start := time.Now()
	fmt.Println("reading")

	defer func() {
		timeDiff := time.Since(start)
		fmt.Println(timeDiff.Microseconds(), "ms")
		fmt.Println(timeDiff.Seconds(), "s")

	}()
	file, err := os.Open(DATA_PATH)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	dataChan := make(chan []byte, DATA_CHAN_BUFSIZE)

	wg.Add(1)
	go func(data <-chan []byte) {
		defer wg.Done()

		for ch := range data {
			reader := bytes.NewReader(ch)
			scanner := bufio.NewScanner(reader)
			for scanner.Scan() {
				//fmt.Println(scanner.Text())
			}
		}
	}(dataChan)

	buf := make([]byte, FILE_CHUNKSIZE)
	var leftover []byte
	for {
		// read a chunk
		_, err = file.Read(buf)
		if err == io.EOF {
			// handle last chunk
			buf = append(leftover, buf...)
			dataChan <- buf
			break
		} else if err != nil {
			panic("alarm")
		}
		//append leftover bytes to the buffer
		buf = append(leftover, buf...)

		//process chunk without new underlying arrays
		var data []byte
		data, leftover = ProcessChunk(buf)
		buf = make([]byte, FILE_CHUNKSIZE) // new underlying array since we are passing around slices
		dataChan <- data
	}
	close(dataChan)
	wg.Wait()
}

// NO EOL DATA ALLOWED IN THIS FUNCTION
func ProcessChunk(chunk []byte) (data, leftover []byte) {
	dataSize := len(chunk)
	lastIndex := dataSize - 1
	//iterate top bottom
	for i := lastIndex; i >= 0; i-- {
		if chunk[i] == CHAR_LB {
			if i == lastIndex {
				data = chunk
				break
			}
			leftover = chunk[i:dataSize]
			data = chunk[0 : i+1]
			break
		}
	}
	return
}
