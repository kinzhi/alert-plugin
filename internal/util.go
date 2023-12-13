// util.go
package internal

import (
	"log"
	"os"
	"strconv"
)

// AppConfig 包含应用程序的配置信息
type AppConfig struct {
	FromEmail          string
	ToEmail            string
	SMTPServer         string
	SMTPPort           int
	EmailPassword      string
	NotificationMethod string
	MessageType        string
	EnableWebhookAuth  bool
	WebhookToken       string
}

// AppConfigInstance 包含应用程序的配置信息的实例
var AppConfigInstance AppConfig

// init 用于初始化应用程序配置
func init() {
	// 获取并保存环境变量值到 AppConfigInstance
	AppConfigInstance.FromEmail = GetEnvWithDefault("FROM_EMAIL", "from@163.com")
	AppConfigInstance.ToEmail = GetEnvWithDefault("TO_EMAIL", "test@163.com")
	AppConfigInstance.SMTPServer = GetEnvWithDefault("SMTP_SERVER", "10.1.1.1")
	AppConfigInstance.SMTPPort = Atoi(GetEnvWithDefault("SMTP_PORT", "25"))
	AppConfigInstance.EmailPassword = GetEnvWithDefault("EMAIL_PASSWORD", "sadads1")
	AppConfigInstance.NotificationMethod = GetEnvWithDefault("NOTIFICATION_METHOD", "smtp")
	AppConfigInstance.MessageType = GetEnvWithDefault("WEBHOOK_MESSAGE_TYPE", "string")
	AppConfigInstance.EnableWebhookAuth = GetEnvWithDefault("ENABLE_WEBHOOK_AUTH", "false") == "true"
	AppConfigInstance.WebhookToken = GetEnvWithDefault("WEBHOOK_TOKEN", "")

}

func GetEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// 处理转换错误
		log.Printf("Error converting string to integer: %v\n", err)
		return 0
	}
	return i
}
