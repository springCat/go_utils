package dbx

import (
	"encoding/json"
	"github.com/springCat/go_utils/joinner"
	"log"
)

type Float64Array []float64

func (a *Float64Array) FromDB(data []byte) error {
	l := len(data)
	s := make([]byte, l)
	copy(s[:1], []byte("["))
	copy(s[1:l-1], data[1:l-1])
	copy(s[l-1:], []byte("]"))
	err := json.Unmarshal(s, a)
	if err != nil {
		log.Println("err:", err)
	}
	return err
}

func (a Float64Array) ToDB() ([]byte, error) {
	s := joinner.Float64Join(a, ",", "{", "}")
	return []byte(s), nil
}
