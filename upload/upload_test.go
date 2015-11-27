package upload
import (
"testing"
"strings"
"go_utils/id"
"fmt"
)

func TestNewName(t *testing.T) {
	fileName := "111.jpgx"
	l := len(fileName)
	i := strings.LastIndex(fileName,"." )
	newFileName := make([]byte,(l-i+32+1))
	id := id.UUID32()
	fmt.Println("id:",id )
	copy(newFileName[:32],id)
	copy(newFileName[33:],fileName[i:])
	fmt.Println(string(newFileName))
}
