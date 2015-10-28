package webservice

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"errors"
	//"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"strings"
)

const (
	// 虚拟服务站ID
	serviceStationId = "87430178-0931-4f68-9b0c-936d67343097"
	//和宇通接口数据传输密钥
	transferMsgKey = "0f6b0bab-958f-490b-8aed-60e3f01934f7"
	//和宇通接口数据传输编码
	transferMsgEncoded = "UTF-8"
)

func init() {

}

type Message struct {
	//服务站sap代码
	ServiceStationSap string
	//请求时间
	RequestTime string
	//单据类型
	BillType string
	//单据号
	BillNumber string
	//操作类型
	OpType string
}
type MessageResponse struct {
	//接口 返回状态
	ReturnStatus string
	//接口 返回值
	ReturnValue string
}
type MessageForward interface {
	MessageForwardYt() (mr MessageResponse, err error)
}

func (m *Message) MessageForwardYt() (mr MessageResponse, err error) {

	mr.ReturnStatus = "S"
	mr.ReturnValue = ""
	sap := m.ServiceStationSap
	rt := m.RequestTime
	bt := m.BillType
	bn := m.BillNumber
	ot := m.OpType
	//校验字段是否为空
	err = errors.New("传输数据格式不正确！")
	if bn == "" || strings.Trim(bn, " ") == "" {
		mr.ReturnStatus = "F"
		mr.ReturnValue = "billNumber is null"
		return
	}
	if "" == bt || strings.Trim(bt, " ") == "" {
		mr.ReturnStatus = "F"
		mr.ReturnValue = "billType is null"
		return
	}
	if "" == ot || strings.Trim(ot, " ") == "" {
		mr.ReturnStatus = "F"
		mr.ReturnValue = "opType is null"
		return
	}
	if "" == rt || strings.Trim(rt, " ") == "" {
		mr.ReturnStatus = "F"
		mr.ReturnValue = "requestTime is null"
		return
	}
	if sap == "" || strings.Trim(sap, " ") == "" {
		mr.ReturnStatus = "F"
		mr.ReturnValue = "serviceStationSap is null"
		return
	}
	b, err := Decode64(m.BillNumber)
	var bb = []byte(transferMsgKey)
	var block cipher.Block
	block, err = des.NewTripleDESCipher(bb)
	if err != nil {
		return mr, err
	}
	blockMode := cipher.NewCBCDecrypter(block, bb[:8])
	origData := make([]byte, len(b))
	blockMode.CryptBlocks(origData, b)
	origData = PKCS5UnPadding(origData)
	return mr, nil
}
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
func Decode64(s string) (b []byte, err error) {
	b, err = base64.StdEncoding.DecodeString(s)
	return
}
