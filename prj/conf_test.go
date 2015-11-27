package prj

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConf(t *testing.T) {
	conf := LoadDefaultConf()
	conf.IP = "111.111.111.111"
	SaveDefaultConf(conf)
}

func TestWebRootFileStr(t *testing.T) {
	err := SaveWebRootFile("", "1.txt", []byte("1234567890qwertyuwet;oiadj.gva.v a/sdvwep'g"))
	if err != nil {
		fmt.Println("err:", err)
	}
	b, err := LoadWebRootFileStr("", "1.txt")
	if err != nil {
		fmt.Println("err:", err)
	}
	assert.Equal(t, "1234567890qwertyuwet;oiadj.gva.v a/sdvwep'g", string(b))
}

func TestOsFileStr(t *testing.T) {
	err := SaveOsFile("/Users/springcat/Desktop/2.txt", []byte("1234567890qwertyuwet;oiadj.gva.v a/sdvwep'g"))
	if err != nil {
		fmt.Println("err:", err)
	}
	b, err := LoadOsFileStr("/Users/springcat/Desktop/2.txt")
	if err != nil {
		fmt.Println("err:", err)
	}
	assert.Equal(t, "1234567890qwertyuwet;oiadj.gva.v a/sdvwep'g", string(b))
}
