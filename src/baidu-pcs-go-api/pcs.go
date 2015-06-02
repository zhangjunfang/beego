package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func NewClient(access_token string) *Pcs {
	p := new(Pcs)
	p.access_token = access_token
	return p
}

//Pcs 类
type Pcs struct {
	access_token string
}

//获取空间配额
func (p Pcs) Quota(method string) ([]byte, error) {
	return HttpGet(api_url + "quota?access_token=" + p.access_token + "&method=" + method)
}

//空间配额信息。
func (p Pcs) QuotaInfo() (*QuotaInfoCallback, error) {
	quota := new(QuotaInfoCallback)

	b, err := p.Quota("info")
	if err != nil {
		return quota, err
	}
	if err := json.Unmarshal(b, quota); err == nil {
		return quota, err
	} else {
		return quota, err
	}
}

//上传文件
func (p Pcs) FileUpload(path string, fileurl string, ondup string) ([]byte, error) {
	if _, err := os.Stat(fileurl); os.IsNotExist(err) {
		return []byte{}, err
	}

	u := api_c_url + "file?method=upload&access_token=" + p.access_token
	u = u + "&path=" + url.QueryEscape(path)
	u = u + "&ondup=" + ondup

	buf := bytes.NewBuffer(nil)
	w := multipart.NewWriter(buf)
	w.CreateFormFile("file", fileurl)

	r, err := http.Post(u, w.FormDataContentType(), buf)
	if err != nil {
		return []byte{}, err
	}
	switch r.StatusCode {

	case http.StatusOK:
		b := readClose(r.Body)

		return b, nil

	default:
		b := readClose(r.Body)
		return b, errors.New("FileUpload Error:" + r.Status + " url :" + u)

	}
	return []byte{}, errors.New("other")
}

//覆盖模式上传
func (p Pcs) FileUploadOverwrite(path string, file string) (*UploadCallback, error) {
	call := new(UploadCallback)

	b, err := p.FileUpload(path, file, "overwrite")
	if err != nil {
		return call, errors.New(err.Error() + ":: " + string(b))
	}
	err = json.Unmarshal(b, call)
	return call, err
}

//自动重命名
func (p Pcs) FileUploadNewcopy(path string, file string) (*UploadCallback, error) {
	call := new(UploadCallback)

	b, err := p.FileUpload(path, file, "newcopy")
	if err != nil {
		return call, errors.New(err.Error() + ":: " + string(b))
	}
	err = json.Unmarshal(b, call)
	return call, err
}

//下载
func (p Pcs) FileDownload(path string) (io.ReadCloser, error) {
	u := api_d_url + "file?method=download&access_token=" + p.access_token
	u = u + "&path=" + url.QueryEscape(path)
	return HttpGetRetReadCloser(u)
}

//创建目录
func (p Pcs) Mkdir(path string) ([]byte, error) {
	u := api_url + "file?method=mkdir&access_token=" + p.access_token
	u = u + "&path=" + url.QueryEscape(path)

	return HttpGet(u)
}

//目录或文件信息。
func (p Pcs) Meta(path string) ([]byte, error) {
	u := api_url + "file?method=meta&access_token=" + p.access_token
	u = u + "&path=" + url.QueryEscape(path)

	return HttpGet(u)
}

//多个目录或文件信息
/*
func (p Pcs) PostMeta(path string) ([]byte, error) {
	u := api_url + "file?method=meta&access_token=" + p.access_token
	u = u + "&path=" + url.QueryEscape(path)

	return HttpGet(u)
}
*/

func (p Pcs) List(path string) ([]byte, error) {
	u := api_url + "file?method=list&access_token=" + p.access_token
	u = u + "&path=" + url.QueryEscape(path)
	return HttpGet(u)
}

func (p Pcs) Move(from string, to string) {
}
func (p Pcs) MoveMore(list []map[string]string) {
}
func (p Pcs) Copy(from string, to string) {

}

func (p Pcs) CopyMore(list []map[string]string) {

}

func (p Pcs) Delete(path string) {

}

func (p Pcs) DeleteMore(list []map[string]string) {
}

func (p Pcs) Search(path string, wd string, re int) ([]byte, error) {
	u := api_url + "file?method=search&access_token=" + p.access_token
	u = u + "&path=" + url.QueryEscape(path)
	u = u + "&wd=" + url.QueryEscape(wd)
	u = u + "&re=" + strconv.Itoa(re)
	return HttpGet(u)
}

//视频转码支持
func (p Pcs) VideoConvert(path string, convert string) ([]byte, error) {
	u := api_url + "file?method=streaming&access_token=" + p.access_token
	u = u + "&path=" + url.QueryEscape(path)
	u = u + "&type=" + convert
	return HttpGet(u)
}

//离线下载任务
func (p Pcs) AddOfflineTask(source string, savePath string, rateLimit int, timeout int, expires int, callback string) ([]byte, error) {
	u := api_url + "services/cloud_dl?method=add_task&access_token=" + p.access_token
	u = u + "&save_path=" + url.QueryEscape(savePath)
	u = u + "&source_url=" + url.QueryEscape(source)
	if rateLimit > -1 {
		u = u + "&rate_limit=" + strconv.Itoa(rateLimit)
	}
	if timeout > -1 {
		u = u + "&timeout=" + strconv.Itoa(timeout)
	}
	if expires > -1 {
		u = u + "&timeout=" + strconv.Itoa(expires)
	}
	if len(callback) > 0 {
		u = u + "&callback=" + url.QueryEscape(callback)
	}

	buf := bytes.NewBuffer(nil)
	form := multipart.NewWriter(buf)
	return HttpPost(u, form, buf)
}

func (p Pcs) GetOfflineTaskList(asc int, needTaskInfo int) ([]byte, error) {
	u := api_url + "services/cloud_dl?method=list_task&access_token=" + p.access_token
	if asc >= 0 {
		u = u + "&ask=" + strconv.Itoa(asc)
	}
	if needTaskInfo > -1 {
		u = u + "&need_task_info=" + strconv.Itoa(needTaskInfo)
	}
	buf := bytes.NewBuffer(nil)
	form := multipart.NewWriter(buf)
	return HttpPost(u, form, buf)
}
func (p Pcs) GetOfflineTaskInfo(taskIds string, opType int, expires int) ([]byte, error) {
	u := api_url + "services/cloud_dl?method=query_task&access_token=" + p.access_token
	if len(taskIds) > 0 {
		u = u + "&task_ids=" + taskIds
	}
	if opType >= 0 {
		u = u + "&opType=" + strconv.Itoa(opType)
	}
	if expires > -1 {
		u = u + "&timeout=" + strconv.Itoa(expires)
	}
	buf := bytes.NewBuffer(nil)
	form := multipart.NewWriter(buf)
	return HttpPost(u, form, buf)
}

func readClose(r io.ReadCloser) []byte {
	bb := bytes.NewBuffer(nil)
	for {
		b := make([]byte, 1024)
		if c, err := r.Read(b); err == nil {
			bb.Write(b[:c])
		} else {
			goto STOP
		}
	}
STOP:
	return bb.Bytes()

}

func HttpGet(u string) ([]byte, error) {
	r, err := http.Get(u)
	b := readClose(r.Body)
	if err != nil {
		return b, errors.New(err.Error() + string(b))
	} else {
		return b, nil
	}
}

func HttpPost(u string, w *multipart.Writer, buf *bytes.Buffer) ([]byte, error) {
	r, err := http.Post(u, w.FormDataContentType(), buf)
	b := readClose(r.Body)
	if err != nil {
		return b, errors.New(err.Error() + string(b))
	} else {
		return b, nil
	}
}

func HttpGetRetReadCloser(u string) (io.ReadCloser, error) {
	r, err := http.Get(u)
	return r.Body, err
}
