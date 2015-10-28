package util

import (
	"github.com/astaxie/beego/cache"
)

func CacheFile() {
	//配置信息如下所示，配置 CachePath 表示缓存的文件目录，FileSuffix 表示文件后缀，DirectoryLevel 表示目录层级，EmbedExpiry 表示过期设置
	bm, err = cache.NewCache("file", `{"CachePath":"cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}`)
	if nil != err {
		log.Info("creat cache object is fail ")
	}
	bm.Put("astaxie", 1, 120)
	log.Info("----------xxxx---------", bm.Get("astaxie"))
	bm.IsExist("astaxie")
	bm.Delete("astaxie")
}
