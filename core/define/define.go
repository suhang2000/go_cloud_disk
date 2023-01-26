package define

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.RegisteredClaims
}

// JwtKey secret key of JWT token
var JwtKey = []byte("cloud-disk-key")

// CodeLength mail code length
var CodeLength = 6

// CodeExpireTime expire time of email code
var CodeExpireTime = time.Second * 300
