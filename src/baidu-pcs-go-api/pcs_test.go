package api

import (
	//"bytes"
	//"io"
	//"os"
	"testing"
)

var client = NewClient("3.8bc38a96fd0a499451a6044b0b0213fd.2592000.1385128807.1345060016-1558780")

/*
func Test_pcs(t *testing.T) {
	aa, err := client.QuotaInfo()
	t.Log(aa.String(), err)
}
*/
func Test_upload(t *testing.T) {

	uc, err := client.FileUploadOverwrite("/apps/baidu_download/333.go", "./pcs.go")
	t.Log(uc, err)

}

func Test_Download(t *testing.T) {

	uc, err := client.FileDownload("/apps/baidu_download/333.go")
	t.Log(uc, err)

	for {

		b := make([]byte, 1024)
		c, err := uc.Read(b)
		if err != nil {
			t.Log("Read Over" + err.Error())
			return
		} else {
			t.Log(b[:c])
		}
	}
}

// 创建文件夹
func Test_mkdir(t *testing.T) {
	u, err := client.Mkdir("/apps/baidu_download/download_by_able")

	t.Log(string(u), err)
}
