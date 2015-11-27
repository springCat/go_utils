package prj

import (
	"fmt"
	"testing"
)

//func TestLoadConf(t *testing.T) {
//	conf := LoadDefaultConf()
//	fmt.Println("conf:", conf)
//	dbConf := conf.Db
//	s := pgurl(dbConf.User,dbConf.Pwd,dbConf.Address,dbConf.Database)
//	fmt.Println("s:", s)
//}



func TestCurrPath(t *testing.T) {
	fmt.Println("CurrPath:", CurrPath())
}

func TestLocalAddr(t *testing.T) {
	fmt.Println("LocalAddr:", LocalAddr())
}

func TestPublicAddr(t *testing.T) {
	fmt.Println("PublicAddr:", PublicAddr())
}

func TestVersionCompare(t *testing.T) {
	v1 := "1.2.1"
	v2 := "1.2.13"
	result,err := VersionCompare(v1,v2)
	fmt.Println("err:", err)
	fmt.Println("result:", result)
}