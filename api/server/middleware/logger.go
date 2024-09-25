package middleware

import (
	"github.com/drizzleent/em/internal/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LogMiddleware(ctx *gin.Context) {
	logger.Info("request", zap.String("method", ctx.Request.Method), zap.Any("req", ctx.Request.URL))
	ctx.Next()

}
