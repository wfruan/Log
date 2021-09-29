package monkey

import (
	"bou.ke/monkey"
	"strings"
	"testing"
)

func TestMyFunc(t *testing.T) {
	monkey.Patch(varys.GetInfoByUID, func(int64)(*varys.UserInfo, error) {
		return &varys.UserInfo{Name: "liwenzhou"}, nil
	})

	ret := MyFunc(123)
	if !strings.Contains(ret, "liwenzhou"){
		t.Fatal()
	}
}