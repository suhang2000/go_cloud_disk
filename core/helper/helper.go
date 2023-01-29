package helper

import (
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"go_cloud_disk/core/define"
	"math/rand"
	"net/textproto"
	"time"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id int, identity, name string) (string, error) {
	userClaim := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	tokenString, err := token.SignedString(define.JwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(token string) (*define.UserClaim, error) {
	if token == "" {
		return nil, errors.New("unauthorized")
	}
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return define.JwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims.Valid {
		return uc, nil
	} else {
		return nil, errors.New("token is invalid")
	}
}

// SendMailCode
// send mail code by outlook
func SendMailCode(mail, code string) error {
	e := &email.Email{
		To:      []string{mail},
		From:    "Cloud Disk Sender <" + define.EmailSender + ">",
		Subject: "send code",
		HTML:    []byte("your code: <h1>" + code + "</h1>"),
		Headers: textproto.MIMEHeader{},
	}

	auth := define.LoginAuth(define.EmailSender, define.EmailPassword)
	err := e.SendWithStartTLS("smtp.office365.com:587", auth, &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.office365.com",
	})
	if err != nil {
		return err
	}
	return nil
}

func RandCode() string {
	s := []rune("1234567890")
	code := make([]rune, define.CodeLength)
	for i := 0; i < define.CodeLength; i++ {
		code[i] = s[rand.Intn(len(s))]
	}
	return string(code)
}

func init() {
	rand.Seed(time.Now().UnixNano())
	//println(time.Now().UnixNano())
}

func UUID() string {
	return uuid.NewV4().String()
}
