package util

import (
	"fmt"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"time"
)

var log *logs.BeeLogger
var bm cache.Cache
var err error

func init() {

	log = logs.NewLogger(10000)
	log.SetLogger("console", "")

	bm, err = cache.NewCache("memory", `{"interval":60}`)
	if nil != err {
		log.Info("creat cache object is fail ")
	}

}
func memory() {
	bm.Put("astaxie", 1, 10)
	time.Sleep(10 * time.Second)
	v, ok := bm.Get("astaxie").(int)
	fmt.Println(v, ok)
	log.Info("===================  value  is  null ?", v, ok)
	bm.IsExist("astaxie")
	bm.Delete("astaxie")
}
