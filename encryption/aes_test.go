package encryption

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncrypt(t *testing.T) {
	v, err := AESEncrypt("1111111111111111", []byte("springcat111111111111"))
	if err != nil {
		t.Error(err)
	}

	value, err := AESDecrypt("1111111111111111", []byte(v))
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "springcat111111111111", value)

	key := GenerateRandomKey(10)

	fmt.Println("string(key):", string(Encode(key)))

}
