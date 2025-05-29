package main

import (
	"github.com/aniketmore311/golang-api-template/pkg/api"
	"github.com/aniketmore311/golang-api-template/pkg/middleware/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg := zap.NewProductionConfig()
	cfg.DisableStacktrace = true
	logger, _ := cfg.Build()
	zap.ReplaceGlobals(logger)
	r := gin.Default()
	r.Use(middleware.TraceIDMiddleware())
	dc := api.DemoController{}
	dc.RegisterRoutes(r)
	r.Run()
}
