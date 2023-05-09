package mailserver

import (
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

var sendMap []string

func SendMail() {
	em := email.NewEmail()
	em.From = "18651684163@163.com"
	em.To = []string{"xushilong@bianfeng.com"}
	em.Subject = "报错测试"
	em.Text = []byte("报错了报错了！！！")

	//设置服务器相关的配置
	err := em.Send("smtp.163.com:25", smtp.PlainAuth("", "18651684163@163.com", "OUKKKZYEMKZWPPVF", "smtp.163.com"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("send successfully ... ")
}
