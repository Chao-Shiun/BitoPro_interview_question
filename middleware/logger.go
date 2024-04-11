package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"BitoPro_interview_question/logger"
)

func LoggingMiddleware(log logger.LogInfoFormat) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// 先執行請求
		c.Next()

		// 請求處理完後計算執行時間
		latency := time.Since(startTime)

		// 獲取請求的詳細資訊
		method := c.Request.Method
		path := c.Request.URL.Path
		status := c.Writer.Status()
		clientIP := c.ClientIP()

		// 檢查是否有錯誤發生
		if len(c.Errors) > 0 {
			// 如果有錯誤，使用 Zap 記錄錯誤
			for _, err := range c.Errors {
				log.Errorf("Request error:",
					zap.String("method", method),
					zap.String("path", path),
					zap.Int("status", status),
					zap.String("client_ip", clientIP),
					zap.Duration("latency", latency),
					zap.Error(err),
				)
			}
		} else {
			// 如果沒有錯誤，使用 Zap 記錄請求資訊
			log.Infof("Request Result:",
				zap.String("method", method),
				zap.String("path", path),
				zap.Int("status", status),
				zap.String("client_ip", clientIP),
				zap.Duration("execute time", latency),
			)
		}
	}
}
