package log

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

var (
	buf bytes.Buffer
	logger = log.New(&buf, "logger: ", log.LUTC | log.LstdFlags | log.Llongfile)
)

func Test_demo1(t *testing.T) {
	logger.Println("QQQ")
	fmt.Println(&buf)
}
