package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/checkout-off-configs-api/src/api/service"
)

type Banana interface {
	//...
}
type BananaImpl struct{}

// START 2 OMIT

type SiteConfigController struct {
	SiteConfigService *service.SiteConfigService
}

// SiteConfig creates a new SiteConfigCtrl
func SiteConfig() *SiteConfigController {
	return &SiteConfigController{
		SiteConfigService: service.SiteConfig(),
	}
}

func (ctrl *SiteConfigController) GetBySite(c *gin.Context) {
	// ...
	// END 2 OMIT
}
