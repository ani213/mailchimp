package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for recipient := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"
		// formatedMsg := fmt.Sprintf("To:%s\r\nSubject:Test Email\r\n\r\n %s", recipient.Email, "Just Testing our email")
		// msg := []byte(formatedMsg)
		msg, err := ExcuteTemplat(recipient)
		if err != nil {
			fmt.Println(err)
		}
		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "ani@group.com", []string{recipient.Email}, []byte(msg))
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func ExcuteTemplat(r Recipient) (string, error) {
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, r); err != nil {
		return "", err
	}
	return tpl.String(), nil
}
