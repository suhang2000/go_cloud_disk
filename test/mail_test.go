package test

import (
	"crypto/tls"
	"errors"
	"github.com/jordan-wright/email"
	"go_cloud_disk/core/define"
	"net/smtp"
	"net/textproto"
	"testing"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unknown fromServer")
		}
	}
	return nil, nil
}

func TestSendEmail(t *testing.T) {
	mail := "xx@gmail.com"
	code := "123123"
	e := &email.Email{
		To:      []string{mail},
		From:    "Cloud Disk Sender <" + define.ConfigEmail.Email + ">",
		Subject: "send code",
		HTML:    []byte("your code: <h1>" + code + "</h1>"),
		Headers: textproto.MIMEHeader{},
	}

	auth := LoginAuth(define.ConfigEmail.Email, define.ConfigEmail.Password)
	err := e.SendWithStartTLS("smtp.office365.com:587", auth, &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.office365.com",
	})
	if err != nil {
		t.Fatal(err)
	}
}
