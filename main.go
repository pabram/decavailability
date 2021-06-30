package main

import (
  "fmt"
  "net/smtp"
)
func main() {
	sendMail()
}



func sendMail() {
  // Sender data DO NOT HARDCODE PASSWORD
  from := "decavailability@gmail.com"
  password := "BambiIsNotDead"

  // Receiver email address.
  to := []string{
    "piotr.abramowicz93@gmail.com",
  }

  // smtp server configuration.
  smtpHost := "smtp.gmail.com"
  smtpPort := "587"

  // Message.
  message := []byte("This is a test email message.")
  
  // Authentication.
  auth := smtp.PlainAuth("", from, password, smtpHost)

  fmt.Println("Authorized")
  
  // Sending email.
  err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println("Email Sent Successfully!")
}