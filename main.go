// main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	fromEmail     string
	toEmail       string
	smtpServer    string
	smtpPort      int
	emailPassword string
	messageType   string
)

func init() {
	// 获取并保存环境变量值
	fromEmail = getEnvWithDefault("FROM_EMAIL", "from@163.com")
	toEmail = getEnvWithDefault("TO_EMAIL", "test@163.com")
	smtpServer = getEnvWithDefault("SMTP_SERVER", "10.1.1.1")
	smtpPort = atoi(getEnvWithDefault("SMTP_PORT", "25"))
	emailPassword = getEnvWithDefault("EMAIL_PASSWORD", "sadads1")
	messageType = getEnvWithDefault("WEBHOOK_MESSAGE_TYPE", "string")
}

func main() {
	http.HandleFunc("/webhook", handleWebhook)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	// 记录接收到的webhook请求日志
	log.Println("Received webhook request")

	// 设置超时时间为10秒
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	// 读取请求体
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		// 记录错误日志
		log.Printf("Error reading request body: %v\n", err)
		return
	}

	// 根据消息类型执行相应的操作
	switch strings.ToLower(messageType) {
	case "string":
		sendEmail("Received string message: "+string(body), ctx)
	case "json":
		var jsonData map[string]interface{}
		err := json.Unmarshal(body, &jsonData)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			// 记录错误日志
			log.Printf("Error decoding JSON: %v\n", err)
			return
		}
		sendEmail("Received JSON message: "+fmt.Sprintf("%+v", jsonData), ctx)
	default:
		http.Error(w, "Invalid message type", http.StatusBadRequest)
		// 记录错误日志
		log.Printf("Invalid message type: %s\n", messageType)
		return
	}

	// 在实际应用中，你可能需要根据消息执行其他操作
	// 这里只是演示了根据消息类型发送邮件

	// 记录接收到的请求日志
	log.Printf("Processed %s message: %s\n", messageType, body)
}
