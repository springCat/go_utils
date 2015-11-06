package ini

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"testing"
)

func TestInt(t *testing.T) {
	iniconf, err := config.NewConfig("ini", "testini.conf")
	if err != nil {
		t.Fatal(err)
	}

	name := iniconf.String("product::name")
	name2 := iniconf.String("release::name")
	name3 := iniconf.String("test::name")

	fmt.Println("name:", name)
	fmt.Println("name2:", name2)
	fmt.Println("name3:", name3)
}
