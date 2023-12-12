// util.go
package main

import (
	"log"
	"os"
	"strconv"
)

func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// 处理转换错误
		log.Printf("Error converting string to integer: %v\n", err)
		return 0
	}
	return i
}
