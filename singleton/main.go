package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct {
	connectionString string
	connections      int
	createdAt        time.Time
}

func (db *Database) Connect() {
	db.connections++
	fmt.Printf("[Database] Connection #%d established to %s\n", db.connections, db.connectionString)
}

func (db *Database) Query(query string) {
	fmt.Printf("[Database] Executing query: %s\n", query)
}

func (db *Database) GetInfo() string {
	return fmt.Sprintf("Database instance created at %s with %d connections",
		db.createdAt.Format("15:04:05"), db.connections)
}

var (
	databaseInstance *Database
	once             sync.Once
)

func GetDatabaseInstance() *Database {
	once.Do(func() {
		fmt.Println("[Singleton] Creating database instance (thread-safe with sync.Once)...")
		databaseInstance = &Database{
			connectionString: "localhost:5432/myapp",
			connections:      0,
			createdAt:        time.Now(),
		}
	})
	return databaseInstance
}

type Logger struct {
	logLevel  string
	createdAt time.Time
	logCount  int
}

func (l *Logger) Log(level, message string) {
	l.logCount++
	timestamp := time.Now().Format("15:04:05")
	fmt.Printf("[%s] [%s] #%d: %s\n", timestamp, level, l.logCount, message)
}

func (l *Logger) Info(message string) {
	l.Log("INFO", message)
}

func (l *Logger) Error(message string) {
	l.Log("ERROR", message)
}

func (l *Logger) GetInfo() string {
	return fmt.Sprintf("Logger instance created at %s with %d logs",
		l.createdAt.Format("15:04:05"), l.logCount)
}

var loggerInstance = &Logger{
	logLevel:  "INFO",
	createdAt: time.Now(),
	logCount:  0,
}

func GetLoggerInstance() *Logger {
	return loggerInstance
}

type ConfigManager struct {
	mu       sync.RWMutex
	settings map[string]string
}

func (c *ConfigManager) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.settings[key] = value
	fmt.Printf("[ConfigManager] Set %s = %s\n", key, value)
}

func (c *ConfigManager) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.settings[key]
}

func (c *ConfigManager) GetAll() map[string]string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	result := make(map[string]string)
	for k, v := range c.settings {
		result[k] = v
	}
	return result
}

var (
	configInstance *ConfigManager
	configOnce     sync.Once
)

func GetConfigManager() *ConfigManager {
	configOnce.Do(func() {
		fmt.Println("[Singleton] Creating config manager instance...")
		configInstance = &ConfigManager{
			settings: make(map[string]string),
		}
		configInstance.settings["app_name"] = "MyApp"
		configInstance.settings["version"] = "1.0.0"
	})
	return configInstance
}

func simulateConcurrentAccess() {
	var wg sync.WaitGroup
	
	fmt.Println("\n--- Testing Thread Safety: Multiple Goroutines ---")
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			db := GetDatabaseInstance()
			fmt.Printf("Goroutine %d got database instance: %p\n", id, db)
		}(i)
	}
	
	wg.Wait()
}

func main() {
	fmt.Println("=== Singleton Pattern Demo ===\n")

	fmt.Println("--- Lazy Initialization with sync.Once (Thread-Safe) ---")
	db1 := GetDatabaseInstance()
	fmt.Printf("Database instance 1: %p\n", db1)
	db1.Connect()
	db1.Query("SELECT * FROM users")

	fmt.Println()
	db2 := GetDatabaseInstance()
	fmt.Printf("Database instance 2: %p\n", db2)
	db2.Connect()
	db2.Query("SELECT * FROM orders")

	fmt.Printf("\nBoth references point to the same instance: %v\n", db1 == db2)
	fmt.Printf("%s\n", db1.GetInfo())

	fmt.Println("\n--- Eager Initialization ---")
	logger1 := GetLoggerInstance()
	fmt.Printf("Logger instance 1: %p\n", logger1)
	logger1.Info("Application started")
	logger1.Error("Sample error message")

	fmt.Println()
	logger2 := GetLoggerInstance()
	fmt.Printf("Logger instance 2: %p\n", logger2)
	logger2.Info("Another log message")

	fmt.Printf("\nBoth references point to the same instance: %v\n", logger1 == logger2)
	fmt.Printf("%s\n", logger1.GetInfo())

	fmt.Println("\n--- Config Manager Singleton ---")
	config1 := GetConfigManager()
	fmt.Printf("Config instance 1: %p\n", config1)
	config1.Set("database_host", "localhost")
	config1.Set("database_port", "5432")

	fmt.Println()
	config2 := GetConfigManager()
	fmt.Printf("Config instance 2: %p\n", config2)
	fmt.Printf("Reading from config2 - database_host: %s\n", config2.Get("database_host"))
	fmt.Printf("Reading from config2 - app_name: %s\n", config2.Get("app_name"))

	fmt.Printf("\nBoth references point to the same instance: %v\n", config1 == config2)

	simulateConcurrentAccess()

	fmt.Println("\n--- Summary ---")
	fmt.Println("All singleton instances maintain single shared state")
	fmt.Println("Thread-safe implementations prevent race conditions")
	fmt.Println("Memory addresses confirm single instance per type")
}
