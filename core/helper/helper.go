package helper

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	"go_cloud_disk/core/define"
	"net/textproto"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id int, identity, name string) (string, error) {
	userClaim := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	tokenString, err := token.SignedString(define.JwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// SendMailCode
// send mail code by outlook
func SendMailCode(mail, code string) error {
	e := &email.Email{
		To:      []string{mail},
		From:    "Cloud Disk Sender <" + define.ConfigEmail.Email + ">",
		Subject: "send code",
		HTML:    []byte("your code: <h1>" + code + "</h1>"),
		Headers: textproto.MIMEHeader{},
	}

	auth := define.LoginAuth(define.ConfigEmail.Email, define.ConfigEmail.Password)
	err := e.SendWithStartTLS("smtp.office365.com:587", auth, &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.office365.com",
	})
	if err != nil {
		return err
	}
	return nil
}
