package main

import (
	"fmt"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for recipient := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"
		formatedMsg := fmt.Sprintf("To:%s\r\nSubject:Test Email\r\n\r\n %s", recipient.Email, "Just Testing our email")
		msg := []byte(formatedMsg)
		err := smtp.SendMail(smtpHost+":"+smtpPort, nil, "ani@group.com", []string{recipient.Email}, msg)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(50 * time.Millisecond)
	}
}
