package mysqlTest

import (
	"database/sql"
	"fmt"

	_"github.com/go-sql-driver/mysql"
)

func Connect() {
	// DSN:Data Source Name
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Connect error")
		panic(err)
	}
	fmt.Println("OK")
	defer db.Close()  // 注意这行代码要写在上面err判断的下面
}
