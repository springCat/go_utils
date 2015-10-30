package collections
import (
	"testing"
)


func TestContains(t *testing.T) {
	t.Parallel()
	a :=  []int{ 1,2,3,4,5,6}
	if !Contains(a,1) {
		t.Error("a should have 1")
		t.FailNow()
	}
	if Contains(a,7) {
		t.Error("a should not have 7")
		t.FailNow()
	}
	if Contains(a,"1") {
		t.Error("a should not have 7")
		t.FailNow()
	}
	if Contains(a,"a") {
		t.Error("a should not have 7")
		t.FailNow()
	}



	type user struct {
		Name string `xml:"name"`
		Age  int64	`xml:"age"`
		Money float64 `xml:"money"`
	}

	u := user{
		Name:"小明",
		Age:12,
		Money:1000.0,
	}

	u1 := user{
		Name:"小明1",
		Age:13,
		Money:1001.0,
	}

	u2 := user{
		Name:"小明",
		Age:12,
		Money:1000.0,
	}

	us := make([]user,0)
	us = append(us,u)

	if !Contains(us,u) {
		t.Error("us should  have u")
		t.FailNow()
	}

	if Contains(us,u1) {
		t.Error("us should not have u1")
		t.FailNow()
	}

	if !Contains(us,u2) {
		t.Error("us should have u2")
		t.FailNow()
	}
}
No newline at end of file
