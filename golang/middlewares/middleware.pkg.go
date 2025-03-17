package middlewares

import (
	"behemoth-pkg/golang/constants"
	"behemoth-pkg/golang/cookies"
	"behemoth-pkg/golang/jwts"
	translator "behemoth-pkg/golang/translators"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := uuid.New().String()
		c.Set(constants.CtxRequestID, requestId)
		c.Writer.Header().Set("X-Request-Id", requestId)
		realIp, exists := c.Get("x-real-ip")
		if !exists {
			realIp = c.RemoteIP()
		}
		c.Set(constants.CtxRemoteIP, realIp)

		c.Next()
	}
}

func SetClientLanguage() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang, _ := cookies.ExtractTokenFromCookie(c, "language")
		translator.ClientLanguage = lang
		c.Next()
	}
}

func JwtTokenCheck(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwts.JwtTokenCheckAndSetToCtx(c, role)
		c.Next()
	}
}

func IncludeGinCtxIntoCtx() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), constants.CtxGin, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
