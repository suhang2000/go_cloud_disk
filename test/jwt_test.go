package test

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"testing"
	"time"
)

type UserClaims struct {
	Name     string `json:"name"`
	Identity string `json:"identity"`
	jwt.RegisteredClaims
}

var testKey = []byte("test_key")

// generate token
func TestGenerateToken(t *testing.T) {
	// newWithClaims
	// create the claims
	userClaim := &UserClaims{
		Name:     "test_name",
		Identity: "test_identity",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}
	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	tokenString, err := token.SignedString(testKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tokenString)
	//	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzdF9uYW1lIiwiaWRlbnRpdHkiOiJ0ZXN0X2lkZW50aXR5IiwiZXhwIjoxNjc0NTAxODQ1fQ.DZ652LbDFRzMSVkQk54oEYBEe_-CI__2flONDrAFRgk
}

// parse token
func TestParseToken(t *testing.T) {
	// ParseWithClaims
	userClaim := new(UserClaims)
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzdF9uYW1lIiwiaWRlbnRpdHkiOiJ0ZXN0X2lkZW50aXR5IiwiZXhwIjoxNjc0NTAxODQ1fQ.DZ652LbDFRzMSVkQk54oEYBEe_-CI__2flONDrAFRgk"
	token, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return testKey, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if token.Valid {
		fmt.Println(userClaim)
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		fmt.Println(claims)
	} else {
		t.Fatal(err)
	}
}
