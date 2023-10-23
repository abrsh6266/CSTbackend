package services

import (
    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailService struct {
    SendGridAPIKey string
}

func NewEmailService(apiKey string) *EmailService {
    return &EmailService{
        SendGridAPIKey: apiKey,
    }
}

func (es *EmailService) SendRegistrationEmail(toEmail, username string) error {
    from := mail.NewEmail("Abrham", "abrsh6265@gmail.com")
    to := mail.NewEmail("Recipient", toEmail)
    subject := "Thanks for Registering!"
    content := mail.NewContent("text/html", "Hello "+username+",<br>Thank you for registering on our website!")
	
    message := mail.NewV3MailInit(from, subject, to, content)
    request := sendgrid.GetRequest(es.SendGridAPIKey, "/v3/mail/send", "https://api.sendgrid.com")
    request.Method = "POST"
    request.Body = mail.GetRequestBody(message)

    _, err := sendgrid.API(request)
    return err
}
