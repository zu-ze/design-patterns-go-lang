package main

import "fmt"

type CPU struct{}

func (c *CPU) Freeze() {
	fmt.Println("[CPU] Freezing processor")
}

func (c *CPU) Jump(position int64) {
	fmt.Printf("[CPU] Jumping to position %d\n", position)
}

func (c *CPU) Execute() {
	fmt.Println("[CPU] Executing instructions")
}

type Memory struct{}

func (m *Memory) Load(position int64, data string) {
	fmt.Printf("[Memory] Loading '%s' at position %d\n", data, position)
}

type HardDrive struct{}

func (h *HardDrive) Read(lba int64, size int) string {
	fmt.Printf("[HardDrive] Reading %d bytes from sector %d\n", size, lba)
	return "BOOT_DATA"
}

type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

func (c *ComputerFacade) Start() {
	fmt.Println("\n[ComputerFacade] Starting computer...\n")
	c.cpu.Freeze()
	c.memory.Load(0x00, "BOOT_SECTOR")
	c.cpu.Jump(0x00)
	bootData := c.hardDrive.Read(0, 1024)
	c.memory.Load(0x100, bootData)
	c.cpu.Execute()
	fmt.Println("\n[ComputerFacade] Computer started successfully!")
}

type VideoFile struct {
	filename string
	codec    string
}

func NewVideoFile(filename string) *VideoFile {
	fmt.Printf("[VideoFile] Loading file: %s\n", filename)
	return &VideoFile{filename: filename}
}

type VideoCodec interface {
	GetType() string
}

type MPEG4Codec struct{}

func (m *MPEG4Codec) GetType() string {
	return "MPEG4"
}

type OggCodec struct{}

func (o *OggCodec) GetType() string {
	return "Ogg"
}

type CodecFactory struct{}

func (c *CodecFactory) Extract(file *VideoFile) VideoCodec {
	fmt.Printf("[CodecFactory] Extracting codec from file\n")
	if file.filename[len(file.filename)-3:] == "mp4" {
		return &MPEG4Codec{}
	}
	return &OggCodec{}
}

type BitrateReader struct{}

func (b *BitrateReader) Read(file *VideoFile, codec VideoCodec) string {
	fmt.Printf("[BitrateReader] Reading bitrate for %s codec\n", codec.GetType())
	return "BUFFER_DATA"
}

func (b *BitrateReader) Convert(buffer string, codec VideoCodec) string {
	fmt.Printf("[BitrateReader] Converting buffer using %s codec\n", codec.GetType())
	return "CONVERTED_DATA"
}

type AudioMixer struct{}

func (a *AudioMixer) Fix(result string) string {
	fmt.Println("[AudioMixer] Fixing audio synchronization")
	return "FIXED_" + result
}

type VideoConverter struct {
	codecFactory  *CodecFactory
	bitrateReader *BitrateReader
	audioMixer    *AudioMixer
}

func NewVideoConverter() *VideoConverter {
	return &VideoConverter{
		codecFactory:  &CodecFactory{},
		bitrateReader: &BitrateReader{},
		audioMixer:    &AudioMixer{},
	}
}

func (v *VideoConverter) Convert(filename, format string) string {
	fmt.Printf("\n[VideoConverter] Starting conversion of %s to %s format\n\n", filename, format)
	
	file := NewVideoFile(filename)
	sourceCodec := v.codecFactory.Extract(file)
	
	var destinationCodec VideoCodec
	if format == "mp4" {
		destinationCodec = &MPEG4Codec{}
	} else {
		destinationCodec = &OggCodec{}
	}
	
	buffer := v.bitrateReader.Read(file, sourceCodec)
	result := v.bitrateReader.Convert(buffer, destinationCodec)
	result = v.audioMixer.Fix(result)
	
	fmt.Printf("\n[VideoConverter] Conversion complete!\n")
	return result
}

type Account struct {
	balance float64
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Debit(amount float64) bool {
	if a.balance >= amount {
		a.balance -= amount
		fmt.Printf("[Account] Debited $%.2f. New balance: $%.2f\n", amount, a.balance)
		return true
	}
	fmt.Printf("[Account] Insufficient funds. Balance: $%.2f\n", a.balance)
	return false
}

func (a *Account) Credit(amount float64) {
	a.balance += amount
	fmt.Printf("[Account] Credited $%.2f. New balance: $%.2f\n", amount, a.balance)
}

type SecurityValidator struct{}

func (s *SecurityValidator) Validate(accountNumber string, pin string) bool {
	fmt.Printf("[SecurityValidator] Validating account %s\n", accountNumber)
	return pin == "1234"
}

type LedgerService struct{}

func (l *LedgerService) RecordTransaction(accountNumber string, amount float64, transactionType string) {
	fmt.Printf("[LedgerService] Recording %s transaction of $%.2f for account %s\n", 
		transactionType, amount, accountNumber)
}

type NotificationService struct{}

func (n *NotificationService) SendNotification(accountNumber string, message string) {
	fmt.Printf("[NotificationService] Sending notification to %s: %s\n", accountNumber, message)
}

type BankingFacade struct {
	accounts            map[string]*Account
	securityValidator   *SecurityValidator
	ledgerService       *LedgerService
	notificationService *NotificationService
}

func NewBankingFacade() *BankingFacade {
	accounts := make(map[string]*Account)
	accounts["ACC001"] = &Account{balance: 1000.00}
	accounts["ACC002"] = &Account{balance: 500.00}
	
	return &BankingFacade{
		accounts:            accounts,
		securityValidator:   &SecurityValidator{},
		ledgerService:       &LedgerService{},
		notificationService: &NotificationService{},
	}
}

func (b *BankingFacade) Withdraw(accountNumber, pin string, amount float64) bool {
	fmt.Printf("\n[BankingFacade] Processing withdrawal of $%.2f from %s\n\n", amount, accountNumber)
	
	if !b.securityValidator.Validate(accountNumber, pin) {
		fmt.Println("[BankingFacade] Security validation failed!")
		return false
	}
	
	account, exists := b.accounts[accountNumber]
	if !exists {
		fmt.Println("[BankingFacade] Account not found!")
		return false
	}
	
	if account.Debit(amount) {
		b.ledgerService.RecordTransaction(accountNumber, amount, "WITHDRAWAL")
		b.notificationService.SendNotification(accountNumber, 
			fmt.Sprintf("Withdrawal of $%.2f completed", amount))
		fmt.Println("\n[BankingFacade] Withdrawal successful!")
		return true
	}
	
	fmt.Println("\n[BankingFacade] Withdrawal failed!")
	return false
}

func (b *BankingFacade) Deposit(accountNumber, pin string, amount float64) bool {
	fmt.Printf("\n[BankingFacade] Processing deposit of $%.2f to %s\n\n", amount, accountNumber)
	
	if !b.securityValidator.Validate(accountNumber, pin) {
		fmt.Println("[BankingFacade] Security validation failed!")
		return false
	}
	
	account, exists := b.accounts[accountNumber]
	if !exists {
		fmt.Println("[BankingFacade] Account not found!")
		return false
	}
	
	account.Credit(amount)
	b.ledgerService.RecordTransaction(accountNumber, amount, "DEPOSIT")
	b.notificationService.SendNotification(accountNumber, 
		fmt.Sprintf("Deposit of $%.2f completed", amount))
	
	fmt.Println("\n[BankingFacade] Deposit successful!")
	return true
}

func (b *BankingFacade) GetBalance(accountNumber, pin string) float64 {
	fmt.Printf("\n[BankingFacade] Checking balance for %s\n\n", accountNumber)
	
	if !b.securityValidator.Validate(accountNumber, pin) {
		fmt.Println("[BankingFacade] Security validation failed!")
		return -1
	}
	
	account, exists := b.accounts[accountNumber]
	if !exists {
		fmt.Println("[BankingFacade] Account not found!")
		return -1
	}
	
	balance := account.GetBalance()
	fmt.Printf("\n[BankingFacade] Current balance: $%.2f\n", balance)
	return balance
}

func main() {
	fmt.Println("=== Facade Pattern Demo ===")

	fmt.Println("\n--- Example 1: Computer Boot System ---")
	fmt.Println("\nWithout Facade, client would need to:")
	fmt.Println("- Create CPU, Memory, HardDrive")
	fmt.Println("- Call CPU.Freeze()")
	fmt.Println("- Call Memory.Load() with correct parameters")
	fmt.Println("- Call CPU.Jump() to correct position")
	fmt.Println("- Call HardDrive.Read() with correct sector")
	fmt.Println("- Call Memory.Load() again with boot data")
	fmt.Println("- Call CPU.Execute()")
	
	fmt.Println("\nWith Facade:")
	computer := NewComputerFacade()
	computer.Start()

	fmt.Println("\n--- Example 2: Video Conversion ---")
	fmt.Println("\nWithout Facade, client would need to:")
	fmt.Println("- Understand video codecs (MPEG4, Ogg, etc.)")
	fmt.Println("- Create CodecFactory, BitrateReader, AudioMixer")
	fmt.Println("- Extract source codec")
	fmt.Println("- Create destination codec")
	fmt.Println("- Read and convert bitrate")
	fmt.Println("- Fix audio synchronization")
	
	fmt.Println("\nWith Facade:")
	converter := NewVideoConverter()
	converter.Convert("video.ogg", "mp4")

	fmt.Println("\n--- Example 3: Banking System ---")
	fmt.Println("\nWithout Facade, client would need to:")
	fmt.Println("- Validate security with SecurityValidator")
	fmt.Println("- Get Account object")
	fmt.Println("- Perform debit/credit operations")
	fmt.Println("- Record transaction in LedgerService")
	fmt.Println("- Send notification via NotificationService")
	
	fmt.Println("\nWith Facade:")
	bank := NewBankingFacade()
	
	bank.GetBalance("ACC001", "1234")
	bank.Withdraw("ACC001", "1234", 200.00)
	bank.Deposit("ACC001", "1234", 150.00)
	bank.GetBalance("ACC001", "1234")
	
	fmt.Println("\n--- Summary ---")
	fmt.Println("Facade provides a simplified interface to complex subsystems")
	fmt.Println("Clients interact with one simple interface instead of many complex ones")
	fmt.Println("Reduces coupling between client and subsystem")
	fmt.Println("Subsystems remain accessible for advanced users who need direct access")
}
