package encryption

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPwd(t *testing.T) {
	pwd,salt := EncodePwdStr("12345678")

	assert.Equal(t,true,EqualPwdStr("12345678",pwd,salt))
}
