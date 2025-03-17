package cookies

import (
	"behemoth-pkg/golang/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func ExtractTokenFromCookie(c *gin.Context, key string) (string, error) {
	cookie, err := c.Cookie(key)
	if err != nil {
		return "", err
	}
	return cookie, nil
}

func SetCookie(ctx *gin.Context, key string, value string, expiration time.Duration) {
	ctx.SetCookie(key, value, utils.TimeDurationToInt(expiration), "/", utils.GetEnv("DOMAIN", "localhost"), false, true)
}

func ClearCookie(ctx *gin.Context, key string) {
	ctx.SetCookie(key, "", -1, "/", utils.GetEnv("DOMAIN", "localhost"), false, true)
}
