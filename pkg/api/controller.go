package api

import (
	"fmt"

	"github.com/aniketmore311/golang-api-template/pkg/errors"
	"github.com/aniketmore311/golang-api-template/pkg/middleware/adapter"
	"github.com/gin-gonic/gin"
)

type DemoController struct {
}

func (dc *DemoController) HealthHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "healthy"})
}

func (dc *DemoController) ExpectedError(c *gin.Context) error {
	err := fmt.Errorf("expected internal error")
	return errors.NewAPIError(err, 400, errors.BadRequest, "customer facing error")
	// return errors.NewAPIError(err, errors.BadRequest, "customer facing error") // status code will be derived from error code
	// errors.ErrorBuilder.WithError(err).WithCode(errors.BadRequest).WithMessage("customer facing error")
	// errors.ErrorBuilder.WithError(err).WithCode(errors.BadRequest) // message will be defaulted to err.Error()
}

func (dc *DemoController) UnExpectedError(c *gin.Context) error {
	return fmt.Errorf("unexpected internal error")
}

func (dc *DemoController) RegisterRoutes(ge *gin.Engine) {
	ge.GET("/health", dc.HealthHandler)
	ge.GET("/error/public", adapter.ErrorAdapter(dc.ExpectedError))
	ge.GET("/error/internal", adapter.ErrorAdapter(dc.UnExpectedError))
}
