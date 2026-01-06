package main

import (
	"fmt"
	"time"
)

type Image interface {
	Display()
}

type RealImage struct {
	filename string
}

func NewRealImage(filename string) *RealImage {
	image := &RealImage{filename: filename}
	image.loadFromDisk()
	return image
}

func (r *RealImage) loadFromDisk() {
	fmt.Printf("[RealImage] Loading image from disk: %s (expensive operation)\n", r.filename)
	time.Sleep(2 * time.Second)
	fmt.Printf("[RealImage] Image loaded: %s\n", r.filename)
}

func (r *RealImage) Display() {
	fmt.Printf("[RealImage] Displaying image: %s\n", r.filename)
}

type ImageProxy struct {
	filename  string
	realImage *RealImage
}

func NewImageProxy(filename string) *ImageProxy {
	fmt.Printf("[ImageProxy] Created proxy for: %s (not loaded yet)\n", filename)
	return &ImageProxy{filename: filename}
}

func (p *ImageProxy) Display() {
	if p.realImage == nil {
		fmt.Println("[ImageProxy] First access - loading real image...")
		p.realImage = NewRealImage(p.filename)
	}
	p.realImage.Display()
}

type Database interface {
	Query(sql string) string
}

type RealDatabase struct {
	connectionString string
}

func NewRealDatabase(connectionString string) *RealDatabase {
	fmt.Printf("[RealDatabase] Connecting to database: %s\n", connectionString)
	return &RealDatabase{connectionString: connectionString}
}

func (d *RealDatabase) Query(sql string) string {
	fmt.Printf("[RealDatabase] Executing query: %s\n", sql)
	return "QUERY_RESULT"
}

type DatabaseProxy struct {
	userRole     string
	realDatabase *RealDatabase
}

func NewDatabaseProxy(userRole string) *DatabaseProxy {
	return &DatabaseProxy{
		userRole:     userRole,
		realDatabase: NewRealDatabase("localhost:5432/mydb"),
	}
}

func (p *DatabaseProxy) Query(sql string) string {
	fmt.Printf("\n[DatabaseProxy] User role: %s attempting query\n", p.userRole)
	
	if !p.checkAccess(sql) {
		fmt.Println("[DatabaseProxy] ACCESS DENIED - Insufficient permissions")
		return "ACCESS_DENIED"
	}
	
	fmt.Println("[DatabaseProxy] Access granted - forwarding to real database")
	result := p.realDatabase.Query(sql)
	p.logQuery(sql)
	
	return result
}

func (p *DatabaseProxy) checkAccess(sql string) bool {
	if p.userRole == "admin" {
		return true
	}
	
	if p.userRole == "user" && (sql[:6] == "SELECT" || sql[:6] == "select") {
		return true
	}
	
	return false
}

func (p *DatabaseProxy) logQuery(sql string) {
	fmt.Printf("[DatabaseProxy] Logging query: %s by user role: %s\n", sql, p.userRole)
}

type VideoDownloader interface {
	Download(videoID string) string
}

type YouTubeDownloader struct{}

func (y *YouTubeDownloader) Download(videoID string) string {
	fmt.Printf("[YouTubeDownloader] Downloading video %s from YouTube (slow operation)...\n", videoID)
	time.Sleep(1 * time.Second)
	return "VIDEO_DATA_" + videoID
}

type CachedYouTubeProxy struct {
	downloader *YouTubeDownloader
	cache      map[string]string
}

func NewCachedYouTubeProxy() *CachedYouTubeProxy {
	return &CachedYouTubeProxy{
		downloader: &YouTubeDownloader{},
		cache:      make(map[string]string),
	}
}

func (p *CachedYouTubeProxy) Download(videoID string) string {
	fmt.Printf("\n[CachedYouTubeProxy] Request for video: %s\n", videoID)
	
	if cachedData, exists := p.cache[videoID]; exists {
		fmt.Printf("[CachedYouTubeProxy] Cache HIT - returning cached data\n")
		return cachedData
	}
	
	fmt.Printf("[CachedYouTubeProxy] Cache MISS - downloading from YouTube\n")
	data := p.downloader.Download(videoID)
	p.cache[videoID] = data
	fmt.Printf("[CachedYouTubeProxy] Cached video: %s\n", videoID)
	
	return data
}

type BankAccount interface {
	Withdraw(amount float64) bool
	GetBalance() float64
}

type RealBankAccount struct {
	accountNumber string
	balance       float64
}

func NewRealBankAccount(accountNumber string, balance float64) *RealBankAccount {
	return &RealBankAccount{
		accountNumber: accountNumber,
		balance:       balance,
	}
}

func (b *RealBankAccount) Withdraw(amount float64) bool {
	if b.balance >= amount {
		b.balance -= amount
		fmt.Printf("[RealBankAccount] Withdrew $%.2f. New balance: $%.2f\n", amount, b.balance)
		return true
	}
	fmt.Printf("[RealBankAccount] Insufficient funds. Balance: $%.2f\n", b.balance)
	return false
}

func (b *RealBankAccount) GetBalance() float64 {
	return b.balance
}

type BankAccountProxy struct {
	pin            string
	realAccount    *RealBankAccount
	attemptCount   int
	locked         bool
}

func NewBankAccountProxy(accountNumber string, balance float64, pin string) *BankAccountProxy {
	return &BankAccountProxy{
		pin:          pin,
		realAccount:  NewRealBankAccount(accountNumber, balance),
		attemptCount: 0,
		locked:       false,
	}
}

func (p *BankAccountProxy) Withdraw(amount float64) bool {
	return p.realAccount.Withdraw(amount)
}

func (p *BankAccountProxy) GetBalance() float64 {
	return p.realAccount.GetBalance()
}

func (p *BankAccountProxy) AuthenticateAndWithdraw(enteredPin string, amount float64) bool {
	fmt.Printf("\n[BankAccountProxy] Authentication attempt...\n")
	
	if p.locked {
		fmt.Println("[BankAccountProxy] Account is LOCKED due to multiple failed attempts")
		return false
	}
	
	if enteredPin != p.pin {
		p.attemptCount++
		fmt.Printf("[BankAccountProxy] Invalid PIN. Attempt %d/3\n", p.attemptCount)
		
		if p.attemptCount >= 3 {
			p.locked = true
			fmt.Println("[BankAccountProxy] Account LOCKED after 3 failed attempts")
		}
		return false
	}
	
	fmt.Println("[BankAccountProxy] Authentication successful")
	p.attemptCount = 0
	return p.Withdraw(amount)
}

type Server interface {
	Request(url string) string
}

type RealServer struct {
	name string
}

func NewRealServer(name string) *RealServer {
	return &RealServer{name: name}
}

func (s *RealServer) Request(url string) string {
	fmt.Printf("[RealServer %s] Processing request: %s\n", s.name, url)
	time.Sleep(500 * time.Millisecond)
	return fmt.Sprintf("Response from %s for %s", s.name, url)
}

type LoadBalancerProxy struct {
	servers      []*RealServer
	currentIndex int
}

func NewLoadBalancerProxy() *LoadBalancerProxy {
	return &LoadBalancerProxy{
		servers: []*RealServer{
			NewRealServer("Server-1"),
			NewRealServer("Server-2"),
			NewRealServer("Server-3"),
		},
		currentIndex: 0,
	}
}

func (p *LoadBalancerProxy) Request(url string) string {
	server := p.servers[p.currentIndex]
	fmt.Printf("\n[LoadBalancerProxy] Routing request to %s (Round-robin)\n", server.name)
	
	p.currentIndex = (p.currentIndex + 1) % len(p.servers)
	
	return server.Request(url)
}

func main() {
	fmt.Println("=== Proxy Pattern Demo ===")

	fmt.Println("\n--- Example 1: Virtual Proxy (Lazy Loading) ---")
	fmt.Println("\nCreating image proxies (images not loaded yet)...")
	image1 := NewImageProxy("photo1.jpg")
	image2 := NewImageProxy("photo2.jpg")
	
	fmt.Println("\nDisplaying first image (triggers loading)...")
	image1.Display()
	
	fmt.Println("\nDisplaying first image again (already loaded)...")
	image1.Display()
	
	fmt.Println("\nDisplaying second image (triggers loading)...")
	image2.Display()

	fmt.Println("\n--- Example 2: Protection Proxy (Access Control) ---")
	
	adminDB := NewDatabaseProxy("admin")
	adminDB.Query("SELECT * FROM users")
	adminDB.Query("DELETE FROM logs WHERE date < '2024-01-01'")
	
	userDB := NewDatabaseProxy("user")
	userDB.Query("SELECT * FROM products")
	userDB.Query("DELETE FROM products WHERE id = 1")

	fmt.Println("\n--- Example 3: Caching Proxy ---")
	
	youtube := NewCachedYouTubeProxy()
	
	fmt.Println("\nFirst request for video:")
	youtube.Download("dQw4w9WgXcQ")
	
	fmt.Println("\nSecond request for same video (cached):")
	youtube.Download("dQw4w9WgXcQ")
	
	fmt.Println("\nRequest for different video:")
	youtube.Download("9bZkp7q19f0")
	
	fmt.Println("\nRequest for first video again (cached):")
	youtube.Download("dQw4w9WgXcQ")

	fmt.Println("\n--- Example 4: Smart Proxy (Authentication) ---")
	
	account := NewBankAccountProxy("ACC123", 1000.00, "1234")
	
	fmt.Printf("Initial balance: $%.2f\n", account.GetBalance())
	
	account.AuthenticateAndWithdraw("0000", 100.00)
	account.AuthenticateAndWithdraw("9999", 100.00)
	account.AuthenticateAndWithdraw("5555", 100.00)
	account.AuthenticateAndWithdraw("1234", 100.00)

	fmt.Println("\n--- Example 5: Remote Proxy (Load Balancer) ---")
	
	loadBalancer := NewLoadBalancerProxy()
	
	fmt.Println("\nMaking multiple requests:")
	loadBalancer.Request("/api/users")
	loadBalancer.Request("/api/products")
	loadBalancer.Request("/api/orders")
	loadBalancer.Request("/api/payments")

	fmt.Println("\n--- Summary ---")
	fmt.Println("Virtual Proxy: Delays expensive object creation until needed")
	fmt.Println("Protection Proxy: Controls access based on permissions")
	fmt.Println("Caching Proxy: Caches results to avoid redundant operations")
	fmt.Println("Smart Proxy: Adds additional logic like authentication")
	fmt.Println("Remote Proxy: Represents object in different location (load balancing)")
	fmt.Println("\nAll proxies maintain the same interface as the real object")
}
