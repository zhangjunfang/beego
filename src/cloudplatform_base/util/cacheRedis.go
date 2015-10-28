package util

import (
	"database/sql"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	redis "github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

/*
功能介绍：
1.使用第三方的组件github.com/garyburd/redigo/redis，
需要注意:1.直接调用Do方法，是直接连接，没有使用连接池pool
        2.第三方组件提供了连接池函数，需要开发人员自己管理连接
2.使用github.com/astaxie/beego/cache/redis，
是对原来的redis封装，并且在每个方法中都有从连接池中借还操作
*/
func CacheRedis() {
	var conn redis.Conn
	var err error
	redis.DialDatabase(0)
	conn, err = redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()
	CheckError(err)
	conn.Do("AUTH", "zhangboyu")
	conn.Do("SET", "email", "i@RincLiu.com")
	bm, err := cache.NewCache("redis", `{"conn":"127.0.0.1:6379"}`)

	bm.Put("zhang:bo:yu", "0505", 120)
	time.Sleep(1 * time.Second)
	log.Info("获取到的值应该是 nil ? or  %v ", bm.Get("zhang:bo:yu"))
	CheakMysql()
}
func CheckError(err error) {
	if err != nil {
		log.Info(" 获取redis 实例 时 发生错误")
		return
	}
}

type Student struct {
	Id        int64
	StudentId string
	Name      string
	Age       int8
}

func (s *Student) String() string {
	return s.StudentId + s.Name + "--------------------"
}
func CheakMysql() {
	db, err := sql.Open("mysql", "root:@/restbird")
	defer db.Close()
	age := 1
	//rows, err := db.Query("SELECT id, sid,sname,sage FROM tbl_student WHERE id =?", age)
	rows, err := db.Query("SELECT id, sid,sname,sage FROM tbl_student WHERE 1 =?", age)
	if err != nil {
		log.Error(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var zhansan Student

		if err := rows.Scan(&zhansan.Id, &zhansan.StudentId, &zhansan.Name, &zhansan.Age); err != nil {
			log.Error(err.Error())
		}
		log.Info("---------------##############:%v", zhansan)
	}
	if err := rows.Err(); err != nil {
		log.Error(err.Error())
	}
}
