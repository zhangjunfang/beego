package util

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	sq = "SELECT u.id FROM s_user AS u INNER JOIN s_role ON u.id = s_role.id"
)

var db *sql.DB

func init() {
	db, err = sql.Open("mysql", "root:@/cppt")

	if err != nil {
		log.Error(err.Error())
	}
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(20)
	db.Ping()
}

/*
1.数据库连接可以在初始化函数中获取 。注意不能再初始化函数中 延迟关闭数据库连接
2.关闭数据库连接，需要在数据操作入口
3.
4.
*/
func JionTable() {
	//defer db.Close()
	row, err := db.Query(sq)
	fmt.Println("-------查询出错了--------", err)
	defer row.Close()
	for row.Next() {
		var id string
		row.Scan(&id)
		fmt.Println("------id-----------", id)
	}

}
