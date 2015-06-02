// 处理json相关操作
package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// 读取一段json数据，将其写入struct f中
func GetJson(dat *string, s interface{}) {
	err := json.Unmarshal([]byte(*dat), s)
	if err != nil {
		panic("Get json failed")
	}
}

// 将struct s转换为json格式
func SetJson(s interface{}) string {
	dat, err := json.Marshal(s)
	if err != nil {
		panic("Set json failed")
	}
	return string(dat)
}

// 从指定文件中读取json数据
func ReadJson(path string, s interface{}) {
	dat, err1 := ioutil.ReadFile(path)
	if err1 != nil {
		panic("Json file fails to open")
	}
	err2 := json.Unmarshal(dat, s)
	if err2 != nil {
		panic("Create json failed")
	}
}

// 将json数据写入指定的文件
func WriteJson(path string, dat *string) {
	_, err0 := os.Stat(path)
	if err0 != nil || !os.IsExist(err0) {
		os.Create(path)
	}
	err := ioutil.WriteFile(path, []byte(*dat), 0644)
	if err != nil {
		panic("Create json file failed")
	}
}
