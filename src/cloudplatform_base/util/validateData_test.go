package util

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

/*
1.模拟支撑平台检验码生成逻辑
*/
func TestCheckMsg(t *testing.T) {
	s := "张俊芳"
	s = strings.ToUpper(hex.EncodeToString([]byte(s)))
	//E5BCA0E4BF8AE88AB3
	//E5BCA0E4BF8AE88AB3
	var temp int64 = 0
	var m int64 = 0
	lens := len(s)
	for i := 0; i < lens; i = i + 2 {
		//		fmt.Println(s[i : i+2])
		m, _ = strconv.ParseInt(s[i:i+2], 16, 0)
		m = m ^ temp
		temp = m
	}
	fmt.Println("---------------------------", strconv.Itoa(int(temp)))
}
func TestSSLValidate(t *testing.T) {
	pool := x509.NewCertPool()
	caCertPath := "certs/cert_server/ca.crt"
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)
	cliCrt, err := tls.LoadX509KeyPair("certs/cert_server/client.crt", "certs/cert_server/client.key")
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
			//InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://server:8081")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func TestInvokeWebservice(t *testing.T) {
	//postStr := CreateSOAPXml2("http://tempuri.org/", "HelloWorld", "FK")
	//PostWebService("https://192.168.111.140:8443/cpcs/services/MessageForwardYt?wsdl", "", postStr)
	//http://webservice.webxml.com.cn/WebServices/MobileCodeWS.asmx?wsdl
	//PostWebService("http://webservice.webxml.com.cn/webservices/qqOnlineWebService.asmx?wsdl", "", postStr)
}

//POST到webService
func PostWebService(url string, method string, value string) string {
	res, err := http.Post(url, "text/xml; charset=utf-8", bytes.NewBuffer([]byte(value)))
	fmt.Println("-------------------------", err)
	defer res.Body.Close()
	//这里随便传递了点东西
	if err != nil {
		fmt.Println("post error", err)
	}
	data, err := ioutil.ReadAll(res.Body)
	//取出主体的内容
	if err != nil {
		fmt.Println("read error", err)
	}
	fmt.Printf("result----%s", data)
	//直接xml从字符串中读取
	//input := `<Person><FirstName>Xu</FirstName><LastName>Xinhua</LastName></Person>`
	// inputReader := strings.NewReader(input)
	// 从文件读取，如可以如下：
	// content, err := ioutil.ReadFile("studygolang.xml")
	// decoder := xml.NewDecoder(bytes.NewBuffer(content))

	var decoder *xml.Decoder = xml.NewDecoder(bytes.NewBuffer(data))
	ParserXMLinMemory(decoder)
	return string(data)
}
func ParserXMLinMemory(decoder *xml.Decoder) {
	var now time.Time = time.Now()
	nowNanosecond := now.Nanosecond()
	var t xml.Token
	var err error
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			{
				name := token.Name.Local
				fmt.Printf("\r\nToken name: %s\n", name)
				for _, attr := range token.Attr {
					attrName := attr.Name.Local
					attrValue := attr.Value
					fmt.Printf("\r\nAn attribute is: %s %s\n", attrName, attrValue)
				}
			}
		// 处理元素结束（标签）
		case xml.EndElement:
			{
				fmt.Printf("\r\nToken of '%s' end\n", token.Name.Local)
			}
		// 处理字符数据（这里就是元素的文本）
		case xml.CharData:
			{
				//if token
				content := string([]byte(token))
				fmt.Printf("\r\nThis is the content: %v\n", content)
			}
		default:
			// ...
		}
	}
	fmt.Println("执行XML文件解析总耗时：", time.Now().Nanosecond()-nowNanosecond)
}
func CreateSOAPXml2(nameSpace string, methodName string, valueStr string) string {
	var soapBody bytes.Buffer
	soapBody.WriteString("<?xml version=\"1.0\" encoding=\"utf-8\"?>")
	soapBody.WriteString("<SOAP:Envelope xmlns:SOAP=\"http://schemas.xmlsoap.org/soap/envelope/\">")
	//soapBody.WriteString(" <SOAP:Header/>")
	soapBody.WriteString("  <SOAP:Body>")
	soapBody.WriteString("     <qqCheckOnline xmlns=\"http://WebXml.com.cn/\">")
	soapBody.WriteString("        <qqCode>419692181</qqCode>")
	soapBody.WriteString("     </qqCheckOnline>")
	soapBody.WriteString("  </SOAP:Body>")
	soapBody.WriteString("</SOAP:Envelope>")
	return soapBody.String()
}
func CreateSOAPXml(nameSpace string, methodName string, valueStr string) string {
	var soapBody bytes.Buffer
	soapBody.WriteString("<?xml version=\"1.0\" encoding=\"utf-8\"?>")
	soapBody.WriteString("<SOAP:Envelope xmlns:SOAP=\"http://schemas.xmlsoap.org/soap/envelope/\">")
	soapBody.WriteString(" <SOAP:Header/>")
	soapBody.WriteString("  <SOAP:Body>")
	soapBody.WriteString("    <ns1:postMessage xmlns:ns1=\"https://192.168.111.140:8443/cpcs/services/MessageForwardYt\">")
	soapBody.WriteString("      <billNumber>BD00360134</billNumber>")
	soapBody.WriteString("      <billType>ServiceOrder</billType>")
	soapBody.WriteString("      <opType>1</opType>")
	soapBody.WriteString("      <RequestTime>2015/10/19 15:13:02</RequestTime>")
	soapBody.WriteString("      <RequestUser>0000134009</RequestUser>")
	soapBody.WriteString("      <RequestType>CREATE</RequestType>")
	soapBody.WriteString("    </ns1:postMessage>")
	soapBody.WriteString("  </SOAP:Body>")
	soapBody.WriteString("</SOAP:Envelope>")
	return soapBody.String()
}
