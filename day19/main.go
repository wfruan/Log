package main

import "day19/mysqlTest"

func main()  {
	//mysqlTest.Connect()
	mysqlTest.Run()
	//mysqlTest.QueryMultiRowDemo()
	//mysqlTest.QueryRowDemo()
	//mysqlTest.InsertRowDemo()
	//mysqlTest.DeleteRowDemo()
	mysqlTest.UpdateRowDemo()
	//mysqlTest.PrepareInsertDemo()

	//危险
	//mysqlTest.SqlInjectDemo("xxx' or 1=1#")
	mysqlTest.TransactionDemo()
}
