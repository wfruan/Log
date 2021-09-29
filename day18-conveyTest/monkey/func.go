package monkey

import (
	_ "bou.ke/monkey"
	"fmt"
)

func MyFunc(uid int64) string  {
	u,err := varys.GetInfoByUID(uid)
	if err != nil {
		return "welcome"
	}

	return fmt.Sprintf("hello %s\n",u.Name)
}
