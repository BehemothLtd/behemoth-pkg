package constants

import (
	"time"
)

type CtxKey string

const (
	// Context Keys
	CtxRequestID              = "requestId"
	CtxRemoteIP               = "remoteIp"
	CtxGin             CtxKey = "ApiGinContext"
	CtxCurrentJwtClaim        = "CurrentJwtClaim"
)

// const (
// 	// Auths
// 	AuthorizationHeader             = "Authorization"
// 	AuthorizationTokenCookie        = "DevCommAuthorizationCookie"
// 	AuthorizationRefreshTokenCookie = "DevCommAuthorizationRefreshTokenCookie"
// 	ClubAccessTokenCookie           = AppName + "-" + "ClubAccessTokenCookie"
// 	ClubRefreshTokenCookie          = AppName + "-" + "ClubRefreshTokenCookie"
// 	CookieExpires                   = 3600
// )

const (
	// Regex Formats
	EmailFormat       = `\A([^@\s]+)@((?:[-a-z0-9]+\.)+[a-z]{2,})\z`
	UrlFormat         = `(\A(http:\/\/|https:\/\/)(www\.)?([-a-zA-Z0-9@:%_\+~#=\-]{1,256}\.[a-zA-Z0-9()]{1,6}|localhost(:\d{1,5})?)|\Atel:\d{8,14}\z|\A{{var[1-9][0-9]*}}\z)`
	PhoneNumberFormat = `(?:\+84|0084|0)[1235789][0-9]{1,2}[0-9]{7}(?:[^\d]+|$)`
)

const (
	// Time Constants
	DefaultJwtTokenExpiration        = 150000 * time.Hour    // short expiration time: 15 mins
	DefaultJwtRefreshTokenExpiration = 7 * 24000 * time.Hour // long expiration time: 7 days
	HHMM_TimeFormat                  = "15:04"
	YYYYMMDD_HHMM_DateTimeFormat     = "2006-01-02 15:04"
)

const (
	// Lengths
	MaxLength30     = 30
	MaxStringLength = 255

	MinPasswordLength = 8
	MaxPasswordLength = 72

	FileMaxSize = 5 * 1024 * 1024 // 5MB
	MinEntryFee = 50000
	Min0        = 0
	Max100      = 100
)

const (
	RoleAdmin    = "admin"
	RoleClubUser = "club_user"
	RolePublic   = "public"
)
