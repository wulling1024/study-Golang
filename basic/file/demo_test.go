package file

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

/**
文件读取的方式选择
 */

func Test_Main(t *testing.T){

}

func readByOS(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 4096)
	var nbytes int

	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF{
			panic(err)
		}

		if n == 0{
			break
		}
		nbytes += n
	}
	return nbytes
}

func readByBufio(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 4096)
	var nbytes int
	rd := bufio.NewReader(file)

	for {
		n, err := rd.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if n == 0{
			break
		}
		nbytes += n
	}
	return nbytes
}

func readByIoutil(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)

	nbytes := len(bytes)
	return nbytes
}

