package timex

import (
	"errors"
	"github.com/springCat/go_utils/types"
	"time"
)

type Timex struct {
	time.Time
}

func NewTimex() *Timex {
	return new(Timex)
}

func (timex *Timex) FromTime(t time.Time) *Timex {
	timex.Time = t
	return timex
}

func (timex *Timex) FromUnix(i int64) *Timex {
	timex.Time = time.Unix(i, 0).In(UTC8)
	return timex
}

func (timex *Timex) Now() *Timex {
	timex.Time = time.Now()
	return timex
}

func (timex *Timex) Unix() int64 {
	return ToUnix(timex.Time)
}

func (timex Timex) MarshalJSON() ([]byte, error) {
	if y := timex.Year(); y < 0 || y > 9999 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	i := timex.In(UTC8).Unix()
	b, err := types.Int64ToByte(i)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (timex *Timex) UnmarshalJSON(data []byte) (err error) {
	i, err := types.ByteToInt64(data)
	if err != nil {
		return err
	}
	timex.Time = FromUnix(i)
	return nil
}
