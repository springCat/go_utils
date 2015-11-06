package id

import (
	"encoding/hex"
	"github.com/satori/go.uuid"
)

func UUID36() string {
	return uuid.NewV4().String()
}

func UUID32() string {
	return noUnderLine(uuid.NewV4())
}

func ObjectIdHex(s string) ObjectId {
	return objectIdHex(s)
}

func NewObjectId() string {
	return newObjectId().Hex()
}

func noUnderLine(u uuid.UUID) string {
	buf := make([]byte, 36)

	hex.Encode(buf[0:8], u[0:4])
	//buf[8] = dash
	hex.Encode(buf[8:12], u[4:6])
	//buf[13] = dash
	hex.Encode(buf[12:16], u[6:8])
	//buf[18] = dash
	hex.Encode(buf[16:20], u[8:10])
	//buf[23] = dash
	hex.Encode(buf[20:], u[10:])

	return string(buf)
}
