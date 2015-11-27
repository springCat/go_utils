package joinner
import (
"testing"
)

func TestIntJoin(t *testing.T) {
	it := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := IntJoin(it, ",","{","}")
	test(t,s,"{1,2,3,4,5,6,7,8,9,10}")
}

func TestInt32Join(t *testing.T) {
	it := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := Int32Join(it, ",","{","}")
	test(t,s,"{1,2,3,4,5,6,7,8,9,10}")
}

func TestInt64Join(t *testing.T) {
	it := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := Int64Join(it, ",","{","}")
	test(t,s,"{1,2,3,4,5,6,7,8,9,10}")
}

func TestFloat32Join(t *testing.T) {
	it := []float32{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.101}
	s := Float32Join(it, ",","{","}")
	test(t,s,"{1.1,2.2,3.3,4.4,5.5,6.6,7.7,8.8,9.9,10.101}")
}

func TestFloat64Join(t *testing.T) {
	it := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.101}
	s := Float64Join(it, ",","{","}")
	test(t,s,"{1.1,2.2,3.3,4.4,5.5,6.6,7.7,8.8,9.9,10.101}")
}

type user struct {
	Id int64
	Name string
}

func (u user) ToJoinStr() string {
	return u.Name
}

func TestStructJoin(t *testing.T) {
	s := make([]Join,0)
	u1 := &user{
		Id:1,
		Name:"小明1",
	}
	u2 := &user{
		Id:2,
		Name:"小明2",
	}
	s = append(s,u1)
	s = append(s,u2)
	ss := StructJoin(s,",","{","}")
	test(t,ss,"{小明1,小明2}")
}


func test(t *testing.T, test interface{}, expected string) {
	res := (test)
	if res != expected {
		t.Log("Case ToString: expected ", expected, " when result is ", res)
		t.FailNow()
	}
}
