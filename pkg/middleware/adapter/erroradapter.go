package adapter

import (
	"errors"

	internerrors "github.com/aniketmore311/golang-api-template/pkg/errors"
	"github.com/aniketmore311/golang-api-template/pkg/middleware/middleware"
	"github.com/aniketmore311/golang-api-template/pkg/types"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrorAdapter(handler types.HandlerFuncWithErr) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := handler(c)
		if err != nil {
			traceID := c.GetString(middleware.TraceIDKey)
			path := c.GetString(middleware.PathKey)
			method := c.GetString(middleware.MethodKey)
			var apiError *internerrors.APIError
			if errors.As(err, &apiError) {
				c.JSON(apiError.GetStatus(), gin.H{
					"status":  apiError.GetStatus(),
					"code":    apiError.GetCode(),
					"detail":  apiError.GetDetail(),
					"traceId": traceID,
				})
				zap.L().Error(
					apiError.Error(),
					zap.String("traceId", traceID),
					zap.String("path", path),
					zap.String("method", method),
				)
			} else {
				c.JSON(500, gin.H{"traceId": traceID, "status": 500, "code": "internal_server_error", "detail": "internal server error"})
				zap.L().Error(
					err.Error(),
					zap.String("traceId", traceID),
					zap.String("path", path),
					zap.String("method", method),
				)
			}
		}
	}
}
