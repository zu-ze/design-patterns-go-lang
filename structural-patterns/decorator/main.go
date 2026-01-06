package main

import (
	"fmt"
	"strings"
	"time"
)

type Coffee interface {
	GetDescription() string
	GetCost() float64
}

type SimpleCoffee struct{}

func (s *SimpleCoffee) GetDescription() string {
	return "Simple Coffee"
}

func (s *SimpleCoffee) GetCost() float64 {
	return 2.00
}

type CoffeeDecorator struct {
	coffee Coffee
}

type MilkDecorator struct {
	CoffeeDecorator
}

func NewMilkDecorator(coffee Coffee) *MilkDecorator {
	return &MilkDecorator{
		CoffeeDecorator: CoffeeDecorator{coffee: coffee},
	}
}

func (m *MilkDecorator) GetDescription() string {
	return m.coffee.GetDescription() + ", Milk"
}

func (m *MilkDecorator) GetCost() float64 {
	return m.coffee.GetCost() + 0.50
}

type SugarDecorator struct {
	CoffeeDecorator
}

func NewSugarDecorator(coffee Coffee) *SugarDecorator {
	return &SugarDecorator{
		CoffeeDecorator: CoffeeDecorator{coffee: coffee},
	}
}

func (s *SugarDecorator) GetDescription() string {
	return s.coffee.GetDescription() + ", Sugar"
}

func (s *SugarDecorator) GetCost() float64 {
	return s.coffee.GetCost() + 0.25
}

type WhippedCreamDecorator struct {
	CoffeeDecorator
}

func NewWhippedCreamDecorator(coffee Coffee) *WhippedCreamDecorator {
	return &WhippedCreamDecorator{
		CoffeeDecorator: CoffeeDecorator{coffee: coffee},
	}
}

func (w *WhippedCreamDecorator) GetDescription() string {
	return w.coffee.GetDescription() + ", Whipped Cream"
}

func (w *WhippedCreamDecorator) GetCost() float64 {
	return w.coffee.GetCost() + 0.75
}

type CaramelDecorator struct {
	CoffeeDecorator
}

func NewCaramelDecorator(coffee Coffee) *CaramelDecorator {
	return &CaramelDecorator{
		CoffeeDecorator: CoffeeDecorator{coffee: coffee},
	}
}

func (c *CaramelDecorator) GetDescription() string {
	return c.coffee.GetDescription() + ", Caramel"
}

func (c *CaramelDecorator) GetCost() float64 {
	return c.coffee.GetCost() + 0.60
}

type DataSource interface {
	WriteData(data string)
	ReadData() string
}

type FileDataSource struct {
	filename string
	data     string
}

func NewFileDataSource(filename string) *FileDataSource {
	return &FileDataSource{
		filename: filename,
		data:     "",
	}
}

func (f *FileDataSource) WriteData(data string) {
	f.data = data
	fmt.Printf("[FileDataSource] Writing to %s: %s\n", f.filename, data)
}

func (f *FileDataSource) ReadData() string {
	fmt.Printf("[FileDataSource] Reading from %s\n", f.filename)
	return f.data
}

type DataSourceDecorator struct {
	wrappee DataSource
}

type EncryptionDecorator struct {
	DataSourceDecorator
}

func NewEncryptionDecorator(source DataSource) *EncryptionDecorator {
	return &EncryptionDecorator{
		DataSourceDecorator: DataSourceDecorator{wrappee: source},
	}
}

func (e *EncryptionDecorator) WriteData(data string) {
	encrypted := e.encrypt(data)
	fmt.Printf("[EncryptionDecorator] Encrypting data\n")
	e.wrappee.WriteData(encrypted)
}

func (e *EncryptionDecorator) ReadData() string {
	data := e.wrappee.ReadData()
	fmt.Printf("[EncryptionDecorator] Decrypting data\n")
	return e.decrypt(data)
}

func (e *EncryptionDecorator) encrypt(data string) string {
	return "ENCRYPTED[" + data + "]"
}

func (e *EncryptionDecorator) decrypt(data string) string {
	return strings.TrimPrefix(strings.TrimSuffix(data, "]"), "ENCRYPTED[")
}

type CompressionDecorator struct {
	DataSourceDecorator
}

func NewCompressionDecorator(source DataSource) *CompressionDecorator {
	return &CompressionDecorator{
		DataSourceDecorator: DataSourceDecorator{wrappee: source},
	}
}

func (c *CompressionDecorator) WriteData(data string) {
	compressed := c.compress(data)
	fmt.Printf("[CompressionDecorator] Compressing data\n")
	c.wrappee.WriteData(compressed)
}

func (c *CompressionDecorator) ReadData() string {
	data := c.wrappee.ReadData()
	fmt.Printf("[CompressionDecorator] Decompressing data\n")
	return c.decompress(data)
}

func (c *CompressionDecorator) compress(data string) string {
	return "COMPRESSED[" + data + "]"
}

func (c *CompressionDecorator) decompress(data string) string {
	return strings.TrimPrefix(strings.TrimSuffix(data, "]"), "COMPRESSED[")
}

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) {
	fmt.Printf("[Email] Sending: %s\n", message)
}

type NotifierDecorator struct {
	wrappee Notifier
}

type SMSDecorator struct {
	NotifierDecorator
}

func NewSMSDecorator(notifier Notifier) *SMSDecorator {
	return &SMSDecorator{
		NotifierDecorator: NotifierDecorator{wrappee: notifier},
	}
}

func (s *SMSDecorator) Send(message string) {
	s.wrappee.Send(message)
	fmt.Printf("[SMS] Also sending: %s\n", message)
}

type SlackDecorator struct {
	NotifierDecorator
}

func NewSlackDecorator(notifier Notifier) *SlackDecorator {
	return &SlackDecorator{
		NotifierDecorator: NotifierDecorator{wrappee: notifier},
	}
}

func (s *SlackDecorator) Send(message string) {
	s.wrappee.Send(message)
	fmt.Printf("[Slack] Also sending: %s\n", message)
}

type FacebookDecorator struct {
	NotifierDecorator
}

func NewFacebookDecorator(notifier Notifier) *FacebookDecorator {
	return &FacebookDecorator{
		NotifierDecorator: NotifierDecorator{wrappee: notifier},
	}
}

func (f *FacebookDecorator) Send(message string) {
	f.wrappee.Send(message)
	fmt.Printf("[Facebook] Also sending: %s\n", message)
}

type LogDecorator struct {
	NotifierDecorator
}

func NewLogDecorator(notifier Notifier) *LogDecorator {
	return &LogDecorator{
		NotifierDecorator: NotifierDecorator{wrappee: notifier},
	}
}

func (l *LogDecorator) Send(message string) {
	fmt.Printf("[Log] %s - Sending notification\n", time.Now().Format("15:04:05"))
	l.wrappee.Send(message)
	fmt.Printf("[Log] %s - Notification sent\n", time.Now().Format("15:04:05"))
}

func printOrder(coffee Coffee) {
	fmt.Printf("Order: %s\n", coffee.GetDescription())
	fmt.Printf("Total Cost: $%.2f\n", coffee.GetCost())
}

func main() {
	fmt.Println("=== Decorator Pattern Demo ===\n")

	fmt.Println("--- Example 1: Coffee Shop ---")
	
	fmt.Println("\n** Simple Coffee **")
	coffee1 := &SimpleCoffee{}
	printOrder(coffee1)
	
	fmt.Println("\n** Coffee with Milk **")
	coffee2 := NewMilkDecorator(&SimpleCoffee{})
	printOrder(coffee2)
	
	fmt.Println("\n** Coffee with Milk and Sugar **")
	coffee3 := NewSugarDecorator(NewMilkDecorator(&SimpleCoffee{}))
	printOrder(coffee3)
	
	fmt.Println("\n** Deluxe Coffee (Milk, Sugar, Whipped Cream, Caramel) **")
	coffee4 := NewCaramelDecorator(
		NewWhippedCreamDecorator(
			NewSugarDecorator(
				NewMilkDecorator(&SimpleCoffee{}),
			),
		),
	)
	printOrder(coffee4)

	fmt.Println("\n--- Example 2: Data Source with Encryption and Compression ---")
	
	fmt.Println("\n** Plain File Write/Read **")
	plainFile := NewFileDataSource("data.txt")
	plainFile.WriteData("Hello World")
	data := plainFile.ReadData()
	fmt.Printf("Retrieved: %s\n", data)
	
	fmt.Println("\n** Encrypted File Write/Read **")
	encryptedFile := NewEncryptionDecorator(NewFileDataSource("encrypted.txt"))
	encryptedFile.WriteData("Secret Message")
	data = encryptedFile.ReadData()
	fmt.Printf("Retrieved: %s\n", data)
	
	fmt.Println("\n** Compressed and Encrypted File Write/Read **")
	secureFile := NewEncryptionDecorator(
		NewCompressionDecorator(
			NewFileDataSource("secure.txt"),
		),
	)
	secureFile.WriteData("Confidential Data")
	data = secureFile.ReadData()
	fmt.Printf("Retrieved: %s\n", data)

	fmt.Println("\n--- Example 3: Multi-Channel Notifications ---")
	
	fmt.Println("\n** Email Only **")
	emailNotifier := &EmailNotifier{}
	emailNotifier.Send("Server is up!")
	
	fmt.Println("\n** Email + SMS **")
	smsNotifier := NewSMSDecorator(&EmailNotifier{})
	smsNotifier.Send("Server is down!")
	
	fmt.Println("\n** Email + SMS + Slack **")
	multiNotifier := NewSlackDecorator(NewSMSDecorator(&EmailNotifier{}))
	multiNotifier.Send("Deployment complete!")
	
	fmt.Println("\n** All Channels with Logging **")
	fullNotifier := NewLogDecorator(
		NewFacebookDecorator(
			NewSlackDecorator(
				NewSMSDecorator(&EmailNotifier{}),
			),
		),
	)
	fullNotifier.Send("Critical alert!")

	fmt.Println("\n--- Summary ---")
	fmt.Println("Decorators add responsibilities to objects dynamically")
	fmt.Println("Multiple decorators can be stacked/chained")
	fmt.Println("Decorators maintain the same interface as the component")
	fmt.Println("Provides flexible alternative to subclassing")
}
