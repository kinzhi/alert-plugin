// email.go
package internal

import (
	"context"
	"log"
	"strings"

	"gopkg.in/gomail.v2"
)

func sendEmail(message string, ctx context.Context) {
	// 判断发邮件方式
	emailMethod := AppConfigInstance.NotificationMethod

	// 记录开始发送邮件的日志
	log.Println("Start sending email")

	select {
	case <-ctx.Done():
		// 超时，记录超时日志
		log.Println("Timeout exceeded. Abort sending email.")
		return
	default:
		// 继续执行
	}

	switch strings.ToLower(emailMethod) {
	case "smtp":
		sendSMTPEmail(message)
	default:
		// 记录无效邮件方式的日志
		log.Printf("Invalid email method: %s\n", emailMethod)
	}
}

func sendSMTPEmail(message string) {
	// 设置邮箱服务器信息

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", AppConfigInstance.FromEmail)
	// 从环境变量获取收件人列表，使用逗号分隔
	toEmails := strings.Split(AppConfigInstance.ToEmail, ",")
	mailer.SetHeader("To", toEmails...)
	mailer.SetHeader("Subject", "Webhook Notification")
	mailer.SetBody("text/plain", message)

	dialer := gomail.NewDialer(AppConfigInstance.SMTPServer, AppConfigInstance.SMTPPort, AppConfigInstance.FromEmail, AppConfigInstance.EmailPassword)

	// 发送邮件
	if err := dialer.DialAndSend(mailer); err != nil {
		// 记录发送邮件失败的日志
		log.Printf("Error sending email: %v\n", err)
		return
	}

	// 记录发送邮件成功的日志
	log.Printf("Sent email: %s\n", message)
}
