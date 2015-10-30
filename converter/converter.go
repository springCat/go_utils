package converter
import (
	"strconv"
	"fmt"
	"encoding/json"
	"encoding/xml"
)


func ToInt64(s string) (int64,error) {
	i,err := strconv.ParseInt(s,0,64)
	return i,err
}

func ToInt32(s string) (int32,error) {
	i,err := strconv.ParseInt(s,0,32)
	if err != nil {
		return 0,err
	}
	return int32(i),err
}

func ToFloat64(s string) (float64,error) {
	f,err := strconv.ParseFloat(s,64)
	return  f,err
}

func ToFloat32(s string) (float32,error) {
	f,err := strconv.ParseFloat(s,64)
	if err != nil {
		return 0,err
	}
	return float32(f),err
}

func ToBool(s string) (bool,error) {
	f,err := strconv.ParseBool(s)
	return f,err
}

func ToString(obj interface{}) string {
	res := fmt.Sprintf("%v", obj)
	return string(res)
}

func ToJson(obj interface{}) (string, error) {
	res, err := json.Marshal(obj)
	if err != nil {
		res = []byte("")
	}
	return string(res), err
}

func ToXml(obj interface{}) (string, error) {
	res, err := xml.Marshal(obj)
	if err != nil {
		res = []byte("")
	}

	return string(res), err
}