# Adapter Pattern

## Overview

The Adapter Pattern is a structural design pattern that allows objects with incompatible interfaces to collaborate. It acts as a bridge between two incompatible interfaces by wrapping an existing class with a new interface. The adapter converts the interface of a class into another interface that clients expect.

## Problem It Solves

When you need to use an existing class but its interface doesn't match the one you need, or when you want to create a reusable class that cooperates with classes that don't have compatible interfaces. This is especially common when integrating legacy code or third-party libraries into new systems.

## When to Use

- When you want to use an existing class but its interface doesn't match what you need
- When you need to integrate legacy code with new systems
- When you want to create reusable classes that work with unrelated or unforeseen classes
- When you need to use several existing subclasses but it's impractical to adapt their interface by subclassing
- When you want to wrap third-party libraries with incompatible interfaces
- When different parts of your system use different interfaces for similar functionality

## Structure

The pattern consists of:

1. **Target Interface**: The interface that the client expects
2. **Adaptee**: The existing class with an incompatible interface
3. **Adapter**: Wraps the adaptee and implements the target interface
4. **Client**: Works with objects through the target interface

## Implementation Types

### Object Adapter (Composition)
Uses composition to wrap the adaptee:
```go
type Adapter struct {
    adaptee *Adaptee
}
```

**Advantages:**
- More flexible (can adapt multiple adaptees)
- Follows composition over inheritance principle
- Can adapt an entire class hierarchy

### Class Adapter (Inheritance)
Uses inheritance/embedding to adapt:
```go
type Adapter struct {
    Adaptee
}
```

**Note:** In Go, this is achieved through embedding rather than traditional inheritance.

**Advantages:**
- Simpler implementation
- Can override adaptee behavior

This implementation primarily uses the **Object Adapter** approach as it's more flexible and idiomatic in Go.

## Implementation Details

The example demonstrates three adapter scenarios:

### 1. Media Player Adapter
- **Problem**: Different media players (Audio, Video, VLC) have incompatible interfaces
- **Solution**: MediaAdapter provides a unified `Play(mediaType, fileName)` interface
- **Benefit**: Client code works with any media type through one interface

### 2. Payment Processor Adapters
- **Problem**: Payment services (Stripe, PayPal, Crypto) have different APIs
- **Solution**: Individual adapters (StripeAdapter, PayPalAdapter, CryptoAdapter) implement `PaymentProcessor` interface
- **Benefit**: Easy to add new payment methods without changing client code

### 3. Legacy Shape System Adapter
- **Problem**: Legacy rectangle drawing uses (x1, y1, x2, y2) coordinates, but modern system uses (x, y, width, height)
- **Solution**: RectangleAdapter converts between coordinate systems
- **Benefit**: Legacy code integrated without modification

## Use Cases

1. **Third-Party Library Integration**: Wrapping external libraries with incompatible APIs
   - Database drivers with different connection methods
   - Multiple SMS/email service providers
   - Different cloud storage APIs (AWS S3, Azure Blob, Google Cloud Storage)

2. **Legacy System Integration**: Making old code work with new systems
   - Adapting SOAP services to REST APIs
   - Converting legacy data formats to modern ones
   - Bridging old authentication systems with new ones

3. **API Versioning**: Supporting multiple API versions
   - Adapting v1 API to v2 interface
   - Maintaining backward compatibility

4. **Cross-Platform Development**: Handling platform-specific implementations
   - File system operations across different OS
   - Network APIs with different implementations

5. **Multiple Data Sources**: Unifying different data sources
   - Different database types (SQL, NoSQL, Graph)
   - Various file formats (JSON, XML, CSV)

6. **Testing**: Creating test adapters for external dependencies
   - Mock adapters for third-party services
   - Test doubles for legacy systems

## Advantages

- **Compatibility**: Allows incompatible interfaces to work together
- **Reusability**: Existing code can be reused without modification
- **Single Responsibility Principle**: Separates interface conversion from business logic
- **Open/Closed Principle**: Can introduce new adapters without changing existing code
- **Flexibility**: Easy to switch between different implementations
- **Legacy Integration**: Integrates legacy code without modifying it

## Disadvantages

- **Complexity**: Increases overall complexity with additional classes
- **Performance**: Additional layer may introduce slight overhead
- **Multiple Adapters**: May need many adapters for complex systems
- **Maintenance**: More classes to maintain and understand
- **Over-engineering**: Can be overkill for simple interface mismatches

## Adapter vs Other Patterns

| Pattern | Purpose | Key Difference |
|---------|---------|----------------|
| **Adapter** | Makes incompatible interfaces work together | Converts existing interface to expected one |
| **Decorator** | Adds behavior to objects | Keeps same interface, adds functionality |
| **Facade** | Simplifies complex subsystems | Provides simplified interface to subsystem |
| **Proxy** | Controls access to objects | Same interface, adds access control |
| **Bridge** | Separates abstraction from implementation | Designed upfront for flexibility |

**Key Distinction**: Adapter works with existing code to fix incompatibility, while other patterns are typically designed into systems from the start.

## Running the Example

```bash
# Navigate to the adapter directory
cd structural-patterns/adapter

# Run the example
go run main.go
```

## Expected Output

```
=== Adapter Pattern Demo ===

--- Example 1: Media Player Adapter ---

[AdvancedMediaPlayer] Received request to play mp3: song.mp3
[AudioPlayer] Playing audio file: song.mp3

[AdvancedMediaPlayer] Received request to play mp4: movie.mp4
[VideoPlayer] Playing video file: movie.mp4

[AdvancedMediaPlayer] Received request to play vlc: documentary.mkv
[VLCPlayer] Playing VLC media: documentary.mkv

[AdvancedMediaPlayer] Received request to play wav: podcast.wav
[AudioPlayer] Playing audio file: podcast.wav

[AdvancedMediaPlayer] Error: unsupported media type: unknown

--- Example 2: Payment Processor Adapters ---

Processing payment of $99.99...
[Stripe] Charging card: $99.99
Payment processed successfully!

Processing payment of $149.50...
[PayPal] Sending payment: $149.50
Payment processed successfully!

Processing payment of $500.00...
[Crypto] Transferring 0.010000 BTC (equivalent to $500.00)
Payment processed successfully!

--- Example 3: Legacy Shape System Adapter ---

Drawing shape at position (10.0,10.0) with size 50.0x50.0
[Circle] Drawing circle at center (35.0,35.0) with radius 25.0

Drawing shape at position (100.0,100.0) with size 80.0x60.0
[LegacyRectangle] Drawing rectangle from (100.0,100.0) to (180.0,160.0)

--- Summary ---
Adapters allow incompatible interfaces to work together
Legacy systems can be integrated without modification
Different payment systems work through a common interface
Media players with different APIs unified under one interface
```

## Key Takeaways

- **Core Purpose**: Convert one interface to another that clients expect
- **Wrapper Pattern**: Adapters wrap existing objects without modifying them
- **Integration Tool**: Perfect for integrating legacy or third-party code
- **Flexibility**: Easy to add new adapters for new incompatible classes
- **Composition Over Inheritance**: Object adapter (composition) is more flexible than class adapter (inheritance)
- **Single Responsibility**: Each adapter handles one specific interface conversion
- **Use Sparingly**: Don't over-engineer; sometimes direct interface changes are better

## Two-Way Adapter (Advanced)

In some cases, you might need a two-way adapter that can convert in both directions:

```go
type TwoWayAdapter struct {
    legacySystem *LegacySystem
}

// Modern interface implementation
func (a *TwoWayAdapter) ModernMethod() {
    a.legacySystem.LegacyMethod()
}

// Legacy interface implementation
func (a *TwoWayAdapter) LegacyMethod() {
    // Convert modern calls to legacy format
}
```

This allows the adapter to be used by both modern and legacy code.

## Best Practices

1. **Keep Adapters Simple**: Focus on interface conversion, not business logic
2. **One Adapter Per Adaptee**: Don't create god-object adapters
3. **Document Conversions**: Clearly document what conversions are happening
4. **Consider Performance**: Be aware of any performance impact from adaptation
5. **Test Thoroughly**: Test both adapted and direct usage
6. **Use Composition**: Prefer object adapter over class adapter for flexibility
7. **Avoid Over-Use**: Don't use adapters when you can change the interface directly

## Common Mistakes

1. **Adding Business Logic**: Adapters should only convert interfaces, not add logic
2. **Too Many Responsibilities**: Adapting multiple unrelated classes in one adapter
3. **Ignoring Performance**: Not considering overhead in performance-critical code
4. **Not Handling Errors**: Forgetting to properly convert error handling between interfaces
5. **Over-Engineering**: Using adapters when a simple interface change would suffice

## Real-World Examples

- **Database Drivers**: `database/sql` package in Go adapts various database drivers
- **HTTP Clients**: Adapting different HTTP client libraries to a common interface
- **Logging Libraries**: Creating a common logging interface for different loggers (zap, logrus, standard log)
- **Cloud Providers**: Unifying AWS, Azure, and GCP APIs under common interfaces
- **Message Queues**: Adapting RabbitMQ, Kafka, and SQS to common message interface
