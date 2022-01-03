package database

import (
	Common "Go_Study/Asset_Admin/common"
	DB_common "Go_Study/Asset_Admin/database/common"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {

	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", DB_common.DNS_test)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func DB_Asset_Add_NameInfo(data Common.RequestInfo) bool {
	// 初始化数据库连接
	err := initDB()

	// 初始化数据库连接失败
	if err != nil {
		fmt.Println("Database Init Failed!, err: ", err)
		return false
	}

	stmt, err := db.Prepare(DB_common.SQL_Asset_Add_NameInfo)

	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.PmName, data.DesignName, data.FinalName, data.Sort, data.Recording)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}

	return true

}
