package webservice

import (
	"bytes"
	base "cloudplatform_base/base"
	"cloudplatform_base/util"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"strconv"
	"strings"
	"time"
)

const (
	// 虚拟服务站ID
	serviceStationId = "87430178-0931-4f68-9b0c-936d67343097"
	//和宇通接口数据传输密钥
	transferMsgKey = "0f6b0bab-958f-490b-8aed-60e3f01934f7"
	//和宇通接口数据传输编码
	transferMsgEncoded = "UTF-8"
)

var bm cache.Cache

func init() {
	var err error
	bm, err = cache.NewCache("redis", `{"conn":"127.0.0.1:6379"}`)
	if nil != err {
		fmt.Println("连接redis数据出错：%V", err)
	}
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
	/*
		解密数据：
		1.base64 解密
		2.TripleDES解密
	*/
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
	/*
	 存储数据到redis中
	*/
	var buffer bytes.Buffer
	buffer.WriteString(base.LEFT_BRACKET)

	//		msg.append(stationId).append("$"); //服务站id。宇通消息会转发到这个服务站
	buffer.WriteString(m.ServiceStationSap)
	buffer.WriteString(base.DOLLAR)
	buffer.WriteString(strings.Replace(base.Rand().Hex(), "-", "", -1))
	buffer.WriteString(base.DOLLAR)
	buffer.WriteString("Y$")
	buffer.WriteString(strconv.Itoa(time.Now().Second()))
	buffer.WriteString(base.DOLLAR)
	stationId, _ := bm.Get(m.ServiceStationSap).(string)
	buffer.WriteString(stationId)
	buffer.WriteString(base.DOLLAR)
	json, _ := json.Marshal([]byte(fmt.Sprint("%v", m)))
	buffer.WriteString(base64.StdEncoding.EncodeToString(json))
	buffer.WriteString(base.DOLLAR)
	buffer.WriteString(util.CheckMsg(buffer.String()))
	buffer.WriteString(base.RIGHT_BRACKET)
	bm.Put(m.ServiceStationSap, buffer.String(), -1)
	/*
	 宇通接口的响应报文
	*/
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
