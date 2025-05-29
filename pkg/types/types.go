package types

import (
	"github.com/gin-gonic/gin"
)

type HandlerFuncWithErr = func(c *gin.Context) error
