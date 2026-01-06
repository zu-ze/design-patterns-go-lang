# Singleton Pattern

## Overview

The Singleton Pattern is a creational design pattern that ensures a class has only one instance and provides a global point of access to that instance. It restricts the instantiation of a class to a single object, which is useful when exactly one object is needed to coordinate actions across the system.

## Problem It Solves

When you need exactly one instance of a class throughout your application (for example, a database connection, configuration manager, or logging service), creating multiple instances would be wasteful or could cause issues. The Singleton pattern ensures that only one instance exists and provides controlled access to it.

## When to Use

- When exactly one instance of a class is needed throughout the application
- When you need global access to a shared resource
- When you want to control concurrent access to a shared resource
- When lazy initialization of a resource is desired
- When creating multiple instances would waste resources or cause conflicts

## Common Use Cases

1. **Database Connection Pools**: Single connection manager for the entire application
2. **Configuration Managers**: Single source of configuration data
3. **Logger Instances**: Centralized logging service
4. **Cache Managers**: Single cache instance shared across the application
5. **Thread Pools**: Managing a fixed set of worker threads
6. **Device Drivers**: Single interface to hardware resources
7. **File System Access**: Coordinating file operations
8. **Application State**: Managing global application state

## Structure

The pattern consists of:

1. **Private Instance Variable**: Holds the single instance
2. **Private/Protected Constructor**: Prevents direct instantiation
3. **Public Static Method**: Provides global access point to the instance
4. **Thread Safety Mechanism**: Ensures safe concurrent access (in multi-threaded environments)

## Implementation in Go

Go doesn't have constructors in the traditional sense, so Singleton implementations differ from classical OOP languages. Here are two common approaches:

### 1. Lazy Initialization with sync.Once (Recommended)

Thread-safe, lazy initialization using Go's `sync.Once`:

```go
var (
    instance *Database
    once     sync.Once
)

func GetDatabaseInstance() *Database {
    once.Do(func() {
        instance = &Database{
            // Initialize fields
        }
    })
    return instance
}
```

**Advantages:**
- Thread-safe by design
- Lazy initialization (created only when first accessed)
- Clean and idiomatic Go code
- No performance overhead after first initialization

### 2. Eager Initialization

Instance created at package initialization:

```go
var loggerInstance = &Logger{
    // Initialize fields
}

func GetLoggerInstance() *Logger {
    return loggerInstance
}
```

**Advantages:**
- Simple and straightforward
- Thread-safe (initialized before main() starts)
- No synchronization overhead

**Disadvantages:**
- Eager initialization (created even if never used)
- No control over initialization timing
- Initialization errors harder to handle

## Implementation Details

The example demonstrates three singleton implementations:

1. **Database**: Lazy initialization with `sync.Once` (recommended for expensive resources)
2. **Logger**: Eager initialization (simple, always-needed resource)
3. **ConfigManager**: Lazy initialization with `sync.Once` plus `sync.RWMutex` for safe concurrent reads/writes

### Thread Safety

In concurrent environments, thread safety is critical:

- **sync.Once**: Guarantees the initialization function runs exactly once, even with concurrent goroutines
- **sync.RWMutex**: Protects shared state that can be modified after initialization
- **Package-level initialization**: Thread-safe by Go's initialization guarantees

## Advantages

- **Controlled Access**: Single point of access to the instance
- **Reduced Memory Footprint**: Only one instance in memory
- **Global State Management**: Shared state across the application
- **Lazy Initialization**: Resource created only when needed (with lazy approach)
- **Thread Safety**: Properly implemented singletons prevent race conditions
- **Easy to Test State**: Single instance makes state tracking simpler

## Disadvantages

- **Global State**: Can make code harder to test and reason about
- **Hidden Dependencies**: Classes using singletons have hidden dependencies
- **Tight Coupling**: Can create tight coupling throughout the codebase
- **Testing Challenges**: Difficult to mock or reset between tests
- **Concurrency Issues**: Improper implementation can cause race conditions
- **Violates Single Responsibility Principle**: Class manages both its responsibility and its own lifecycle
- **Can Hide Bad Design**: Often used as a band-aid for poor architecture

## Anti-Pattern Concerns

The Singleton pattern is sometimes considered an anti-pattern because:

1. **Global State**: Introduces global state, making code harder to reason about
2. **Testing Difficulty**: Hard to mock and reset for unit tests
3. **Hidden Dependencies**: Dependencies not visible in function signatures
4. **Thread Safety Complexity**: Easy to implement incorrectly in concurrent scenarios
5. **Overuse**: Often used when dependency injection would be better

### When NOT to Use Singleton

- When you can use dependency injection instead
- When you need multiple instances in testing
- When the "single instance" requirement might change
- When it creates tight coupling in your codebase
- When it makes testing significantly harder

### Better Alternatives

Consider these alternatives:

1. **Dependency Injection**: Pass instances explicitly
2. **Context Pattern**: Use Go's context to pass shared resources
3. **Package-level Functions**: Simple stateless operations
4. **Factory Pattern**: Control instance creation without global state

## Running the Example

```bash
# Navigate to the singleton directory
cd singleton

# Run the example
go run main.go
```

## Expected Output

```
=== Singleton Pattern Demo ===

--- Lazy Initialization with sync.Once (Thread-Safe) ---
[Singleton] Creating database instance (thread-safe with sync.Once)...
Database instance 1: 0xc000010230
[Database] Connection #1 established to localhost:5432/myapp
[Database] Executing query: SELECT * FROM users

Database instance 2: 0xc000010230
[Database] Connection #2 established to localhost:5432/myapp
[Database] Executing query: SELECT * FROM orders

Both references point to the same instance: true
Database instance created at 14:23:45 with 2 connections

--- Eager Initialization ---
Logger instance 1: 0xc000010240
[14:23:45] [INFO] #1: Application started
[14:23:45] [ERROR] #2: Sample error message

Logger instance 2: 0xc000010240
[14:23:45] [INFO] #3: Another log message

Both references point to the same instance: true
Logger instance created at 14:23:45 with 3 logs

--- Config Manager Singleton ---
[Singleton] Creating config manager instance...
Config instance 1: 0xc000012180
[ConfigManager] Set database_host = localhost
[ConfigManager] Set database_port = 5432

Config instance 2: 0xc000012180
Reading from config2 - database_host: localhost
Reading from config2 - app_name: MyApp

Both references point to the same instance: true

--- Testing Thread Safety: Multiple Goroutines ---
Goroutine 1 got database instance: 0xc000010230
Goroutine 3 got database instance: 0xc000010230
Goroutine 0 got database instance: 0xc000010230
Goroutine 4 got database instance: 0xc000010230
Goroutine 2 got database instance: 0xc000010230
```

Notice that all memory addresses are identical, confirming single instance per type.

## Key Takeaways

- **Core Purpose**: Ensure only one instance exists and provide global access
- **Go Implementation**: Use `sync.Once` for lazy, thread-safe initialization
- **Thread Safety**: Critical in concurrent environments - always use proper synchronization
- **Use Sparingly**: Consider if dependency injection or other patterns would be better
- **Testing Impact**: Be aware of testing challenges introduced by global state
- **Eager vs Lazy**: Choose based on whether resource is always needed and initialization cost
- **Not Always Best**: Sometimes considered an anti-pattern - use judiciously

## Comparison of Approaches

| Approach | Thread-Safe | Lazy Init | Complexity | Use Case |
|----------|-------------|-----------|------------|----------|
| **sync.Once** | ✓ | ✓ | Medium | Expensive resources, lazy loading |
| **Eager Init** | ✓ | ✗ | Low | Always-needed, cheap resources |
| **Naive (unsafe)** | ✗ | ✓ | Low | **Never use in production!** |

## Best Practices

1. **Use sync.Once**: For lazy initialization in Go, `sync.Once` is the standard approach
2. **Document Thread Safety**: Clearly document thread safety guarantees
3. **Consider Alternatives**: Think about whether dependency injection would be better
4. **Add Reset for Testing**: Consider adding a reset method for testing (but be careful!)
5. **Protect Mutable State**: Use mutexes for any state changes after initialization
6. **Keep It Simple**: Don't over-engineer - sometimes a simple package-level variable is fine

## Testing Considerations

Testing singletons can be challenging:

```go
// Problem: Singleton state persists between tests
func TestA(t *testing.T) {
    config := GetConfigManager()
    config.Set("key", "value")
    // Test relies on this state
}

func TestB(t *testing.T) {
    config := GetConfigManager()
    // Still has state from TestA!
}
```

**Solutions:**
1. Add a reset method (use with caution in production)
2. Use dependency injection instead of global singletons
3. Design tests to be independent of singleton state
4. Use test fixtures that set known state before each test

## Singleton vs Other Patterns

| Pattern | Purpose | Instances |
|---------|---------|-----------|
| **Singleton** | Ensure single instance | Exactly one |
| **Factory Method** | Create objects flexibly | Many |
| **Prototype** | Clone existing objects | Many (clones) |
| **Multiton** | Limited set of instances | Few (named) |
| **Object Pool** | Reuse expensive objects | Fixed pool |
