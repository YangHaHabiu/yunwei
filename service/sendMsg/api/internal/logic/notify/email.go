package notify

import (
	"crypto/tls"
	"fmt"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"net"
	"net/smtp"
	"strings"
)

type PEmailConfig struct {
	Host     string
	Port     int
	User     string
	Pwd      string
	From     string
	NickName string
}

type PEmail struct {
	Config  *PEmailConfig
	Subject string
	Body    string
	To      string
	Format  string
	Mid     int
}

type EmailAccount struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	config *PEmailConfig
)

func SendToEmailChan(to, subject, body, jsonContent string) error {
	var content EmailAccount
	err := jsonx.Unmarshal([]byte(jsonContent), &content)
	if err != nil {
		logx.Error("解析邮件消息内容失败，原因：" + err.Error())
		return err
	}

	config = &PEmailConfig{
		Host: content.Host,
		From: content.Username,
		Port: content.Port,
		User: content.Username,
		Pwd:  content.Password,
	}
	email := &PEmail{
		Config:  config,
		Body:    body,
		Subject: subject,
		Format:  "txt",
		To:      to,
	}
	if err := email.SendToEmail(); err != nil {
		logx.Errorf("发送邮件失败，失败帐号[%s] 原因:%s\n", content.Username, err.Error())
		return err
	}
	return nil

}

func (pe *PEmail) SendToEmail() error {
	auth := smtp.PlainAuth("", pe.Config.User, pe.Config.Pwd, pe.Config.Host)
	var contentType string
	if pe.Format == "html" {
		contentType = "Content-Type: text/" + pe.Format + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + pe.To + "\r\nCc: " + pe.Config.User + "\r\nBcc: " + pe.Config.User + "\r\nFrom: " + pe.Config.NickName + "<" + pe.Config.User +
		">\r\nSubject: " + pe.Subject + "\r\n" + contentType + "\r\n\r\n" + pe.Body)
	sendTo := strings.Split(pe.To, ",")
	err := SendMailUsingTLS(
		fmt.Sprintf("%s:%d", pe.Config.Host, pe.Config.Port),
		auth,
		pe.Config.From,
		sendTo,
		msg,
	)
	return err
}

func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		//log.Println("Dialing Error:", err)
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func SendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {
	c, err := Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
