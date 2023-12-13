// webhook.go

package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	if AppConfigInstance.EnableWebhookAuth && !authenticateWebhook(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 记录接收到的webhook请求日志
	log.Println("Received webhook request")

	// 设置超时时间为10秒
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	// 读取请求体
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		// 记录错误日志
		log.Printf("Error reading request body: %v\n", err)
		return
	}

	// 根据消息类型执行相应的操作
	switch strings.ToLower(AppConfigInstance.MessageType) {
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
		log.Printf("Invalid message type: %s\n", AppConfigInstance.MessageType)
		return
	}

	// 在实际应用中，你可能需要根据消息执行其他操作
	// 这里只是演示了根据消息类型发送邮件

	// 记录接收到的请求日志
	log.Printf("Processed %s message: %s\n", AppConfigInstance.MessageType, body)
}
