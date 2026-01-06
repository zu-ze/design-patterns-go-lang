package main

import (
	"fmt"
	"strings"
)

type MediaPlayer interface {
	Play(mediaType, fileName string) error
}

type AudioPlayer struct{}

func (a *AudioPlayer) PlayAudio(fileName string) {
	fmt.Printf("[AudioPlayer] Playing audio file: %s\n", fileName)
}

type VideoPlayer struct{}

func (v *VideoPlayer) PlayVideo(fileName string) {
	fmt.Printf("[VideoPlayer] Playing video file: %s\n", fileName)
}

type VLCPlayer struct{}

func (v *VLCPlayer) PlayVLC(fileName string) {
	fmt.Printf("[VLCPlayer] Playing VLC media: %s\n", fileName)
}

type MediaAdapter struct {
	audioPlayer *AudioPlayer
	videoPlayer *VideoPlayer
	vlcPlayer   *VLCPlayer
}

func NewMediaAdapter() *MediaAdapter {
	return &MediaAdapter{
		audioPlayer: &AudioPlayer{},
		videoPlayer: &VideoPlayer{},
		vlcPlayer:   &VLCPlayer{},
	}
}

func (m *MediaAdapter) Play(mediaType, fileName string) error {
	switch strings.ToLower(mediaType) {
	case "mp3", "wav", "audio":
		m.audioPlayer.PlayAudio(fileName)
		return nil
	case "mp4", "avi", "video":
		m.videoPlayer.PlayVideo(fileName)
		return nil
	case "vlc", "mkv":
		m.vlcPlayer.PlayVLC(fileName)
		return nil
	default:
		return fmt.Errorf("unsupported media type: %s", mediaType)
	}
}

type AdvancedMediaPlayer struct {
	adapter *MediaAdapter
}

func NewAdvancedMediaPlayer() *AdvancedMediaPlayer {
	return &AdvancedMediaPlayer{
		adapter: NewMediaAdapter(),
	}
}

func (a *AdvancedMediaPlayer) Play(mediaType, fileName string) error {
	fmt.Printf("\n[AdvancedMediaPlayer] Received request to play %s: %s\n", mediaType, fileName)
	return a.adapter.Play(mediaType, fileName)
}

type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

type StripePayment struct{}

func (s *StripePayment) ChargeCard(amount float64) error {
	fmt.Printf("[Stripe] Charging card: $%.2f\n", amount)
	return nil
}

type PayPalPayment struct{}

func (p *PayPalPayment) SendPayment(amount float64) error {
	fmt.Printf("[PayPal] Sending payment: $%.2f\n", amount)
	return nil
}

type CryptoPayment struct{}

func (c *CryptoPayment) TransferCrypto(amountUSD float64) error {
	btcAmount := amountUSD / 50000.0
	fmt.Printf("[Crypto] Transferring %.6f BTC (equivalent to $%.2f)\n", btcAmount, amountUSD)
	return nil
}

type StripeAdapter struct {
	stripe *StripePayment
}

func NewStripeAdapter() *StripeAdapter {
	return &StripeAdapter{
		stripe: &StripePayment{},
	}
}

func (s *StripeAdapter) ProcessPayment(amount float64) error {
	return s.stripe.ChargeCard(amount)
}

type PayPalAdapter struct {
	paypal *PayPalPayment
}

func NewPayPalAdapter() *PayPalAdapter {
	return &PayPalAdapter{
		paypal: &PayPalPayment{},
	}
}

func (p *PayPalAdapter) ProcessPayment(amount float64) error {
	return p.paypal.SendPayment(amount)
}

type CryptoAdapter struct {
	crypto *CryptoPayment
}

func NewCryptoAdapter() *CryptoAdapter {
	return &CryptoAdapter{
		crypto: &CryptoPayment{},
	}
}

func (c *CryptoAdapter) ProcessPayment(amount float64) error {
	return c.crypto.TransferCrypto(amount)
}

func processPayment(processor PaymentProcessor, amount float64) {
	fmt.Printf("\nProcessing payment of $%.2f...\n", amount)
	if err := processor.ProcessPayment(amount); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Payment processed successfully!")
	}
}

type LegacyRectangle struct {
	x1, y1, x2, y2 float64
}

func (r *LegacyRectangle) Draw(x1, y1, x2, y2 float64) {
	r.x1, r.y1, r.x2, r.y2 = x1, y1, x2, y2
	fmt.Printf("[LegacyRectangle] Drawing rectangle from (%.1f,%.1f) to (%.1f,%.1f)\n", x1, y1, x2, y2)
}

type Shape interface {
	Draw(x, y, width, height float64)
	GetInfo() string
}

type RectangleAdapter struct {
	legacy *LegacyRectangle
}

func NewRectangleAdapter() *RectangleAdapter {
	return &RectangleAdapter{
		legacy: &LegacyRectangle{},
	}
}

func (r *RectangleAdapter) Draw(x, y, width, height float64) {
	x2 := x + width
	y2 := y + height
	r.legacy.Draw(x, y, x2, y2)
}

func (r *RectangleAdapter) GetInfo() string {
	return fmt.Sprintf("Rectangle Adapter wrapping legacy drawing system")
}

type Circle struct{}

func (c *Circle) Draw(x, y, width, height float64) {
	radius := width / 2
	centerX := x + radius
	centerY := y + radius
	fmt.Printf("[Circle] Drawing circle at center (%.1f,%.1f) with radius %.1f\n", centerX, centerY, radius)
}

func (c *Circle) GetInfo() string {
	return "Modern circle implementation"
}

func drawShape(shape Shape, x, y, width, height float64) {
	fmt.Printf("\nDrawing shape at position (%.1f,%.1f) with size %.1fx%.1f\n", x, y, width, height)
	shape.Draw(x, y, width, height)
}

func main() {
	fmt.Println("=== Adapter Pattern Demo ===\n")

	fmt.Println("--- Example 1: Media Player Adapter ---")
	player := NewAdvancedMediaPlayer()
	
	player.Play("mp3", "song.mp3")
	player.Play("mp4", "movie.mp4")
	player.Play("vlc", "documentary.mkv")
	player.Play("wav", "podcast.wav")
	
	err := player.Play("unknown", "file.xyz")
	if err != nil {
		fmt.Printf("\n[AdvancedMediaPlayer] Error: %v\n", err)
	}

	fmt.Println("\n--- Example 2: Payment Processor Adapters ---")
	
	stripeAdapter := NewStripeAdapter()
	processPayment(stripeAdapter, 99.99)
	
	paypalAdapter := NewPayPalAdapter()
	processPayment(paypalAdapter, 149.50)
	
	cryptoAdapter := NewCryptoAdapter()
	processPayment(cryptoAdapter, 500.00)

	fmt.Println("\n--- Example 3: Legacy Shape System Adapter ---")
	
	circle := &Circle{}
	drawShape(circle, 10, 10, 50, 50)
	
	rectangleAdapter := NewRectangleAdapter()
	drawShape(rectangleAdapter, 100, 100, 80, 60)
	
	fmt.Println("\n--- Summary ---")
	fmt.Println("Adapters allow incompatible interfaces to work together")
	fmt.Println("Legacy systems can be integrated without modification")
	fmt.Println("Different payment systems work through a common interface")
	fmt.Println("Media players with different APIs unified under one interface")
}
