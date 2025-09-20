package router

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"api-gateway/internal/config"
	"api-gateway/internal/middleware"
	"api-gateway/internal/proxy"

	"github.com/gin-gonic/gin"
)

func New(cfg *config.Config, proxies map[string]*proxy.ProxyWrap) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.RequestID())
	r.Use(middleware.Logger())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "run"})
	})

	r.Any("/api/v1/:service/*proxyPath", func(c *gin.Context) {
		service := c.Param("service")
		p, ok := proxies[service]
		if !ok {
			c.JSON(http.StatusBadGateway, gin.H{"error": "unknown service"})
			return
		}

		ctx, cancel := context.WithTimeout(c.Request.Context(), cfg.ProxyTimeout)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)

		// delete "api/v1" in url
		parsed, err := url.Parse(strings.Replace(c.Request.URL.String(), "api/v1/"+service+"/", "", 1))

		if err != nil {
			return
		}

		c.Request.URL = parsed

		p.ServeHTTP(c.Writer, c.Request)
	})

	return r
}
