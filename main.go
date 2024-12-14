package main

import (
	"bufio"
	"compress/gzip"
	"encoding/csv"
	"log"
	"os"
	"sync"
)

func main() {
	ProcessFile()
}

func process(workerId int, record []string) {
	_ = workerId
	var newA []string
	newA = append(newA, record...)
	// if len(record) > 0 {
	// 	// log.Printf("Worker %d processando: %s\n", workerId, record[0])
	// }
	_ = newA
}

func ProcessFile() {
	file, err := os.Open("data/largeData.csv.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer gzReader.Close()

	bufferedReader := bufio.NewReaderSize(gzReader, 64*1024)
	csvReader := csv.NewReader(bufferedReader)
	lines := make(chan []string, 100)

	var wg sync.WaitGroup
	numWorkers := 10

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			for record := range lines {
				process(workerId, record)
			}
		}(i)
	}

	for {
		record, err := csvReader.Read()

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf(err.Error())
		}

		lines <- record
	}

	close(lines)
	wg.Wait()
}
