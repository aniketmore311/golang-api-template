package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const TraceIDKey = "TraceID"
const MethodKey = "Method"
const PathKey = "Path"
const TraceIDHeader = "X-Trace-ID"

func TraceIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader(TraceIDHeader)

		// If the trace ID doesn't exist, generate a new one
		if traceID == "" {
			traceID = uuid.New().String()
		}

		// Set the trace ID in the context
		c.Set(TraceIDKey, traceID)
		c.Set(PathKey, c.Request.URL.Path)
		c.Set(MethodKey, c.Request.Method)

		// Add it to the response headers
		c.Writer.Header().Set(TraceIDHeader, traceID)

		// Proceed with the request
		c.Next()
	}
}
