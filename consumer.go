package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()

	smtpHost := "localhost"
	smtpPort := "1025"
	addr := smtpHost + ":" + smtpPort

	for recipient := range ch {
		msg, err := executeTemplate(recipient)
		if err != nil {
			log.Printf("Worker %d: error parsing template for %s: %v", id, recipient.Email, err)
			continue
		}

		fmt.Printf("Worker %d: Sending email to %s\n", id, recipient.Email)

		err = smtp.SendMail(addr, nil, "tom@gmail.com", []string{recipient.Email}, []byte(msg))
		if err != nil {
			log.Printf("Worker %d: failed to send to %s: %v", id, recipient.Email, err)
			continue
		}

		time.Sleep(50 * time.Millisecond)
		fmt.Printf("Worker %d: Sent email to %s\n", id, recipient.Email)
	}
}
