package main

import "fmt"

// ======== ABSTRACT FACTORY ========
// Su rol es creasr objectos (productos)

type NotificationFactory interface {
	// SendNotification() // Si la factory tiene SendNotification(), ya no es factory, es un service.
	Sender() Sender
	Notifier() Notifier
}

// ======== PRODUCTS (interfaces) ========

type Sender interface {
	Method() string  // sms, email, etc.
	Channel() string // twilio, ses, etc.
}

type Notifier interface {
	Notify(message string) error
}

// ======== SMS IMPLEMENTATION ========

type SMSFactory struct{}

func (SMSFactory) Sender() Sender { return SMSSender{} }
func (f SMSFactory) Notifier() Notifier {
	return SMSNotifier{
		sender: f.Sender(), // o SMSSender{}
	}
}

// Concrete product: Sender SMS

type SMSSender struct{}

func (SMSSender) Method() string { return "sms" }

func (SMSSender) Channel() string { return "Twilio" }

// Concrete product: Notifier SMS

type SMSNotifier struct {
	sender Sender
}

func (n SMSNotifier) Notify(message string) error {
	fmt.Printf("Enviando %s via %s: %s\n", n.sender.Method(), n.sender.Channel(), message)
	return nil
}

// ========= EMAIL IMPLEMENTATION ==========

type EmailFactory struct{}

func (EmailFactory) Sender() Sender { return EmailSender{} }
func (f EmailFactory) Notifier() Notifier {
	return EmailNotifier{
		sender: f.Sender(),
	}
}

// Concrete product: Sender Email

type EmailSender struct{}

func (EmailSender) Method() string  { return "email" }
func (EmailSender) Channel() string { return "ses" }

// Concrete product: Sender Notifier

type EmailNotifier struct {
	sender Sender
}

func (n EmailNotifier) Notify(message string) error {
	fmt.Printf("Enviando %s via %s: %s\n", n.sender.Method(), n.sender.Channel(), message)
	return nil
}

// ========= FACTORY SELECTOR ==========

func NewNotificationFactory(notificationType string) (NotificationFactory, error) {
	switch notificationType {
	case "sms":
		return SMSFactory{}, nil
	case "email":
		return EmailFactory{}, nil
	default:
		return nil, fmt.Errorf("tipo no soportado")
	}
}

func main() {
	smsFactory, _ := NewNotificationFactory("sms")
	emailFactory, _ := NewNotificationFactory(("email"))

	_ = smsFactory.Notifier().Notify("Notificando mensage por SMS")
	_ = emailFactory.Notifier().Notify("Notificando mensage por Email")
}
