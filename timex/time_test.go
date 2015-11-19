package timex

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

)

func TestNow(t *testing.T) {
	now := Now()
	time := FromUnix(now)
	l := ToUnix(time)
	assert.Equal(t, now, l)
}


func TestTIme(t *testing.T) {

	fmt.Println("day0",Today0())
	fmt.Println("time:", NextDay0())
}
