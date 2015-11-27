package encryption

import (
	"bytes"
	"crypto/md5"
	"log"
)

const saltSize = 10

func EncodePwd(password []byte) (pwd, salt []byte) {

	randomKey := GenerateRandomKey(saltSize)

	cipher := md5.New()
	cipher.Write(randomKey)
	result := cipher.Sum(password)
	salt = Encode(randomKey)
	pwd = Encode(result)
	return
}

func EqualPwd(password []byte, dbPassword, dbSalt []byte) bool {
	rawSalt, err := Decode(dbSalt)
	if err != nil {
		log.Println("err:", err)
		return false
	}
	rawPassword, err := Decode(dbPassword)
	if err != nil {
		log.Println("err:", err)
		return false
	}
	cipher := md5.New()
	cipher.Write(rawSalt)
	result := cipher.Sum([]byte(password))

	return bytes.Equal(result, rawPassword)
}

func EncodePwdStr(password string) (pwd, salt string) {
	p,s := EncodePwd([]byte(password))
	pwd = string(p)
	salt = string(s)
	return
}

func EqualPwdStr(password string, dbPassword, dbSalt string) bool {
	return EqualPwd([]byte(password), []byte(dbPassword), []byte(dbSalt))
}
