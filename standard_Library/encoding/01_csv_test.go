package encoding

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestCSV(t *testing.T) {

}

//readFile -> Read data from a CSV file
func readFile() {
	file, err := os.Open("source.csv")
	if err != nil {
		fmt.Printf("Filed open failed, error:%v\n", err)
		panic(err)
	}
	reader := csv.NewReader(bufio.NewReader(file))
	reader.FieldsPerRecord = -1

	for {
		record, err := reader.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Printf("End of read file")
				break
			}
			fmt.Printf("File read failed, error=%v\n", err)
		}
		// process data
		go func(data string) {

		}(record[0])
	}
}

//writerFile -> Write data to a CSV file
func writerFile() {
	file, err := os.Create("result.csv")
	if err != nil {
		fmt.Printf("Filed create failed, error:%v\n", err)
		panic(err)
	}
	writer := csv.NewWriter(bufio.NewWriter(file))

	// process data
	items := []string{}
	err = writer.Write(items)
	if err != nil {
		fmt.Printf("Cache write error:%v\n", err)
		panic(err)
	}

	writer.Flush()
}
