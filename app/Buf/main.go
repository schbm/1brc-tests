package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

const DATA_PATH = "../../data/measurements.txt"
const CHAR_LB = '\n'
const DATA_CHAN_BUFSIZE = 1_000_000

var wg sync.WaitGroup

func main() {
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

	dataChan := make(chan string, DATA_CHAN_BUFSIZE)

	wg.Add(1)
	go func(data <-chan string) {
		defer wg.Done()
		for _ = range data {
		}
	}(dataChan)

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		dataChan <- line
	}
	close(dataChan)
	wg.Wait()
}
