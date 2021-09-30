package mysqlTest

import "fmt"

// SqlInjectDemo sql注入示例:根据name字段查询user表的函数
func SqlInjectDemo(name string) {
	//name inject to sql
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var u user
	err := db.QueryRow(sqlStr).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	fmt.Printf("user:%#v\n", u)
}
