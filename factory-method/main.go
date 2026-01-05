package main

import "fmt"

type Notification interface {
	Send(message string) error
	GetType() string
}

type EmailNotification struct {
	recipient string
}

func (e *EmailNotification) Send(message string) error {
	fmt.Printf("[EMAIL] Sending to %s: %s\n", e.recipient, message)
	return nil
}

func (e *EmailNotification) GetType() string {
	return "Email"
}

type SMSNotification struct {
	phoneNumber string
}

func (s *SMSNotification) Send(message string) error {
	fmt.Printf("[SMS] Sending to %s: %s\n", s.phoneNumber, message)
	return nil
}

func (s *SMSNotification) GetType() string {
	return "SMS"
}

type PushNotification struct {
	deviceID string
}

func (p *PushNotification) Send(message string) error {
	fmt.Printf("[PUSH] Sending to device %s: %s\n", p.deviceID, message)
	return nil
}

func (p *PushNotification) GetType() string {
	return "Push"
}

type NotificationFactory interface {
	CreateNotification() Notification
}

type EmailNotificationFactory struct {
	recipient string
}

func (f *EmailNotificationFactory) CreateNotification() Notification {
	return &EmailNotification{recipient: f.recipient}
}

type SMSNotificationFactory struct {
	phoneNumber string
}

func (f *SMSNotificationFactory) CreateNotification() Notification {
	return &SMSNotification{phoneNumber: f.phoneNumber}
}

type PushNotificationFactory struct {
	deviceID string
}

func (f *PushNotificationFactory) CreateNotification() Notification {
	return &PushNotification{deviceID: f.deviceID}
}

func sendNotification(factory NotificationFactory, message string) {
	notification := factory.CreateNotification()
	fmt.Printf("Created %s notification\n", notification.GetType())
	notification.Send(message)
}

func main() {
	fmt.Println("=== Factory Method Pattern Demo ===\n")

	emailFactory := &EmailNotificationFactory{recipient: "user@example.com"}
	sendNotification(emailFactory, "Your order has been shipped!")

	fmt.Println()

	smsFactory := &SMSNotificationFactory{phoneNumber: "+1234567890"}
	sendNotification(smsFactory, "Your verification code is 123456")

	fmt.Println()

	pushFactory := &PushNotificationFactory{deviceID: "device-abc-123"}
	sendNotification(pushFactory, "You have a new message")
}
