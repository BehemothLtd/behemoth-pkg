package config

import (
	"net/url"
	"strings"

	"github.com/BehemothLtd/behemoth-pkg/golang/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func getAllowedDomains() []string {
	domains := []string{}
	if customDomains := utils.GetEnv("ALLOWED_DOMAINS", ""); customDomains != "" {
		domains = append(domains, strings.Split(customDomains, ",")...)
	}
	return domains
}

func isAllowedOrigin(origin string, allowedDomains []string) bool {
	if (utils.IsDevelopmentEnv() || utils.IsLocalEnv()) && strings.HasPrefix(origin, "http://localhost") {
		return true
	}

	u, err := url.Parse(origin)
	if err != nil {
		return false
	}

	for _, domain := range allowedDomains {
		if u.Host == domain || strings.HasSuffix(u.Host, "."+domain) {
			return true
		}
	}
	return false
}

func CorsConfig() gin.HandlerFunc {
	allowedDomains := getAllowedDomains()
	log.Debug().Msg("Allowed domains: " + strings.Join(allowedDomains, ", "))

	config := cors.Config{
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return isAllowedOrigin(origin, allowedDomains)
		},
	}

	return cors.New(config)
}
