package encryption
import (
"crypto/aes"
	"go_utils/synx"
	"crypto/cipher"
)


var AESKeyMap = aesCipher{synx.NewRWMutexMap()}

type aesCipher struct {
	m *synx.RWMutexMap
}

func (a aesCipher) key(key string) (cipher.Block,error) {
	block,exist := a.m.Get(key)
	if !exist {
		b, err := aes.NewCipher([]byte(key))
		if err != nil {
			return nil,err
		}
		a.m.Put(key,b)
		return b,nil
	}
	return block.(cipher.Block),nil
}

func AESEncrypt(key string,plaintext []byte) (string, error){
	block,err := AESKeyMap.key(key)
	if err != nil {
		return "",err
	}
	raw,err := Encrypt(block,plaintext)
	if err != nil {
		return "",err
	}
	v := Encode(raw)
	return string(v),nil
}

func AESDecrypt(key string,ciphertext []byte) (string, error) {
	block,err := AESKeyMap.key(key)
	if err != nil {
		return "",err
	}
	raw,err := Decode(ciphertext)
	if err != nil {
		return "",err
	}
	v,err := Decrypt(block,raw)
	if err != nil {
		return "",err
	}
	return string(v),nil
}
