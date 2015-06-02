package util

import (
	"fmt"
	"net/smtp"
	"strings"
)

/*
 *  user : xxxx@163.com login smtp server user
 *  password: xxxxx login smtp server password
 *  host: smtp.example.com:port   smtp.163.com:25
 *  to: example@example.com;example1@163.com;example2@sina.com.cn;...
 *  subject:The subject of mail
 *  body: The content of mail
 *  mailtype: mail type html or text
 */

func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

//func main() {
//	//beego.Run()
//	user := "xxxx@163.com"  //使用那个邮箱发送邮件
//	password := "xxxxxxxx"  //邮箱密码
//	host := "smtp.163.com:25"
//	to := "zhangjunfang0505@163.com;3203317@qq.com"  //接受邮件列表

//	subject := "Test send email by zhangboyu"  //邮件主题
//邮件内容
//	body := `
//    <html>
//    <body>
//    <h3>
//    "Test send email by zhangboyu"
//    </h3>
//    </body>
//    </html>
//    `
//	fmt.Println("send email")
//	err := SendMail(user, password, host, to, subject, body, "html")//发送邮件
//	if err != nil {
//		fmt.Println("send mail error!")
//		fmt.Println(err)
//	} else {
//		fmt.Println("send mail success!")
//	}

//}
