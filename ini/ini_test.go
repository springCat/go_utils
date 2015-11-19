package ini
import (
	"testing"
	"gopkg.in/ini.v1"
	"fmt"
)

//import (
//	"fmt"
//	"github.com/astaxie/beego/config"
//	"testing"
//)
//
//func TestInt(t *testing.T) {
//	iniconf, err := config.NewConfig("ini", "testini.conf")
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	name := iniconf.String("product::name")
//	name2 := iniconf.String("release::name")
//	name3 := iniconf.String("test::name")
//
//	fmt.Println("name:", name)
//	fmt.Println("name2:", name2)
//	fmt.Println("name3:", name3)
//}
type Note struct {
	Content string
	Cities  []string
}


func TestIni(t *testing.T) {
	cfg, err := ini.Load("testini.conf")
	if err != nil {
		fmt.Println(err)
	}

//	s, _ = cfg.NewSection("springcat")
//	s.

	fmt.Println("cfg:", cfg.Section("test").Key("IMPORT_PATH").String())


	n := new(Note)
	err = cfg.Section("Note").MapTo(n)
	fmt.Println("n:", n)
}