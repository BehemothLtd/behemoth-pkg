package jwts

import (
	"github.com/BehemothLtd/behemoth-pkg/golang/constants"
	"github.com/BehemothLtd/behemoth-pkg/golang/cookies"
	"github.com/BehemothLtd/behemoth-pkg/golang/exceptions"
	"github.com/BehemothLtd/behemoth-pkg/golang/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaim struct {
	Sub  uint32
	Role string
	jwt.RegisteredClaims
}

var jwtSecret = []byte(utils.GetEnv("JWT_SECRET_KEY", "secret"))

// JwtTokenCheckAndSetToCtx extracts JWT from cookie, validates it, and sets user info in context.
func JwtTokenCheckAndSetToCtx(c *gin.Context, target string) {

	jwtClaim, err := ExtractAndReadFromCookie(c, target)

	if err == nil && jwtClaim.Role == target {
		WriteToCtxCurrentUser(c, *jwtClaim)
	}
}

// WriteToCtxCurrentUser stores user ID and role in Gin's context.
func WriteToCtxCurrentUser(c *gin.Context, jwtClaim JwtClaim) {
	c.Set(constants.CtxCurrentJwtClaim, jwtClaim)
}

// ExtractAndReadFromCookie extracts and decodes JWT from cookie.
func ExtractAndReadFromCookie(c *gin.Context, key string) (*JwtClaim, error) {
	token, err := cookies.ExtractTokenFromCookie(c, key)
	if err != nil {
		return nil, err
	}
	return ParseToken(token)
}

// ParseToken decodes and validates a JWT token.
func ParseToken(jwtToken string) (*JwtClaim, error) {
	var jwtClaim JwtClaim
	if err := DecodeJwtToken(jwtToken, &jwtClaim); err != nil {
		return nil, exceptions.NewBadRequestError(nil)
	}
	return &jwtClaim, nil
}

// GenerateJwtToken creates a signed JWT token.
func GenerateJwtToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// DecodeJwtToken validates and extracts claims from a JWT token.
func DecodeJwtToken(tokenString string, userClaim jwt.Claims) error {
	token, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return exceptions.NewBadRequestError(nil)
	}
	return nil
}
