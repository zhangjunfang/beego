package api

import (
	"fmt"
)

//Quota 配额信息。
type QuotaInfoCallback struct {
	Quota      int64 //磁盘配额
	Used       int64 // 已使用量
	Request_id int64 `request_id`
}

func (q QuotaInfoCallback) String() string {
	return fmt.Sprintf("quota %d  , used %d , reqid %d ", q.Quota, q.Used, q.Request_id)
}

//提交反馈
type UploadCallback struct {
	Path  string
	Size  uint64
	Ctime uint64
	Mtime uint64
	Md5   string
	fs_id uint64
}
