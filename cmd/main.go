package main

import (
	"github.com/aniketmore311/golang-api-template/pkg/api"
	"github.com/aniketmore311/golang-api-template/pkg/middleware/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	r := gin.Default()
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	r.Use(middleware.TraceIDMiddleware())
	dc := api.DemoController{}
	dc.RegisterRoutes(r)
	r.Run()
}
