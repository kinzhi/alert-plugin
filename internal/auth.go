// auth.go

package internal

import (
	"log"
	"net/http"
	"strings"
)

func authenticateWebhook(r *http.Request) bool {
	// 如果认证未启用，直接返回true
	if !AppConfigInstance.EnableWebhookAuth {
		return true
	}

	// 从请求头中获取Authorization信息
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		log.Println("Webhook authentication failed: Authorization header is missing")
		return false
	}

	// 检查Bearer Token
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		log.Println("Webhook authentication failed: Invalid Authorization header format")
		return false
	}

	// 验证Token
	if tokenParts[1] != AppConfigInstance.WebhookToken {
		log.Println("Webhook authentication failed: Invalid Bearer Token")
		return false
	}

	log.Println("Webhook authentication successful")
	return true
}
