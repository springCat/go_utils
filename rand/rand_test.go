package rand

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	s, _ := IntRange(3, 10)
	fmt.Println("s:", s)
}
