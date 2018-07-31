package main

import (
	"gopkg.in/gomail.v2"
	"crypto/tls"
	"strings"
)

type EmailMessage struct {
	EmailSMTPServerHost string
	EmailSMTPServerPort int
	EmailUserName       string
	EmailPassword       string
	Message             string
	From                string
	To                  string
	CC                  string
	Subject             string
	Body                string
	Attach              string
}

func (em *EmailMessage) Send() error {
	m := gomail.NewMessage()
	m.SetHeader("From", em.From)
	m.SetHeader("To", strings.Split(em.To, "|")...)
	m.SetHeader("Cc", strings.Split(em.CC, "|")...)
	m.SetHeader("Subject", em.Subject)
	m.SetBody("text/html", em.Body)
	if FileExists(em.Attach) {
		m.Attach(em.Attach)
	}

	d := gomail.NewDialer(em.EmailSMTPServerHost, em.EmailSMTPServerPort, em.EmailUserName, em.EmailPassword)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		panic(err)
		return err
	} else {
		return nil
	}
}
