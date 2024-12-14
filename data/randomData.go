package data

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"time"
)

func RandSeq(n int) string {
	letters := []rune("1234567890")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

func WriteDataOnCsv() {
	file, err := os.Create("largeData.csv")
	if err != nil {
		log.Fatal(err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 1000000; i++ {
		generateData := RandSeq(11)
		if err := writer.Write([]string{generateData}); err != nil {
			panic(err)
		}
	}
}
