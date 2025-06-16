package mail

import (
	"fmt"

	"gopkg.in/mail.v2"
)

func SendPaymentEmail(to string, amount float64) error {
	m := mail.NewMessage()
	m.SetHeader("From", "noreply@example.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Погашение платежа")
	m.SetBody("text/html", "С вашего счета списано "+fmt.Sprintf("%.2f", amount)+" по кредиту.")

	d := mail.NewDialer("mailhog", 1025, "", "")
	return d.DialAndSend(m)
}

func SendTest() {
	m := mail.NewMessage()
	m.SetHeader("From", "test@example.com")
	m.SetHeader("To", "om@example.com")
	m.SetHeader("Subject", "Test")
	m.SetBody("text/plain", "GoMail works!")
}
