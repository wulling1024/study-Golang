package time

import (
	"fmt"
	"testing"
	"time"
)

func TestDemo2(t *testing.T) {
	start := time.Now()

	time.Sleep(1)

	fmt.Println(int64(time.Now().Sub(start) / time.Microsecond))
}
