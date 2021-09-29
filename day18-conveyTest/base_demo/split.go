package base_demo

import "strings"

func Split(s,sep string) (res []string)  {
	i := strings.Index(s,sep)
	for i > -1 {
		res = append(res,s[:i])
		s=s[i+len(sep):]
		i=strings.Index(s,sep)
	}
	res = append(res,s)
	return
}

