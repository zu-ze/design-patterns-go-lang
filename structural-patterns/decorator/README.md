# Decorator Pattern

## Overview

The Decorator Pattern is a structural design pattern that lets you attach new behaviors to objects by placing these objects inside special wrapper objects called decorators. Decorators provide a flexible alternative to subclassing for extending functionality. The pattern allows behavior to be added to individual objects, dynamically, without affecting the behavior of other objects from the same class.

## Problem It Solves

When you need to add responsibilities to objects dynamically and transparently without affecting other objects, or when extending functionality through inheritance is impractical (too many subclasses needed or subclassing is not possible). For example, if you have 4 features to add optionally, subclassing would require 2^4 = 16 classes for all combinations.

## When to Use

- When you want to add responsibilities to individual objects dynamically and transparently
- When extension by subclassing is impractical or impossible
- When you need to add/remove responsibilities at runtime
- When you want to avoid an explosion of subclasses to support every combination
- When you want to wrap objects with additional behavior without changing their interface
- When multiple independent extensions need to be combined flexibly

## Structure

The pattern consists of:

1. **Component Interface**: Defines the interface for objects that can have responsibilities added
2. **Concrete Component**: The object to which additional responsibilities can be attached
3. **Decorator Base Class**: Maintains a reference to a Component and implements the Component interface
4. **Concrete Decorators**: Add responsibilities to the component

## Key Characteristics

### Same Interface
Decorators implement the same interface as the component they wrap:
```go
type Component interface {
    Operation()
}

type Decorator struct {
    component Component  // Wraps a component
}
```

### Recursive Wrapping
Decorators can wrap other decorators:
```
ConcreteDecorator1 → ConcreteDecorator2 → ConcreteComponent
```

### Transparent to Client
Client code treats decorated and undecorated objects identically.

## Implementation Details

The example demonstrates three decorator scenarios:

### 1. Coffee Shop System
- **Component**: Coffee interface
- **Concrete Component**: SimpleCoffee
- **Decorators**: MilkDecorator, SugarDecorator, WhippedCreamDecorator, CaramelDecorator
- **Behavior**: Each decorator adds to the description and cost

**Key Feature**: Mix and match any combination of add-ons dynamically.

### 2. Data Source with Security Layers
- **Component**: DataSource interface
- **Concrete Component**: FileDataSource
- **Decorators**: EncryptionDecorator, CompressionDecorator
- **Behavior**: Add encryption and compression layers

**Key Feature**: Layer security features (compress then encrypt, or vice versa).

### 3. Multi-Channel Notification System
- **Component**: Notifier interface
- **Concrete Component**: EmailNotifier
- **Decorators**: SMSDecorator, SlackDecorator, FacebookDecorator, LogDecorator
- **Behavior**: Add additional notification channels

**Key Feature**: Send notifications through multiple channels by stacking decorators.

## Use Cases

1. **Stream Processing**: Adding buffering, encryption, compression to I/O streams
   - Java's InputStream/OutputStream decorators
   - Go's io.Reader/io.Writer wrappers

2. **GUI Components**: Adding borders, scrollbars, shadows to windows
   - Adding visual effects dynamically
   - Combining multiple visual enhancements

3. **Web Middleware**: Adding authentication, logging, caching to HTTP handlers
   - Express.js middleware
   - Go HTTP middleware chains

4. **Text Formatting**: Adding bold, italic, underline formatting
   - Can combine multiple formats
   - Each decorator adds HTML/markup

5. **Coffee Shop/Pizza Orders**: Adding toppings and extras
   - Add ingredients dynamically
   - Calculate total cost recursively

6. **Access Control**: Adding permissions checks, rate limiting
   - Layer security features
   - Each decorator adds a security check

7. **Caching**: Adding cache layers at different levels
   - Memory cache → Disk cache → Network
   - Multiple caching strategies

## Advantages

- **Flexibility**: More flexible than static inheritance
- **Single Responsibility Principle**: Each decorator handles one concern
- **Runtime Modification**: Add/remove responsibilities at runtime
- **Combination**: Can combine decorators in various ways
- **Open/Closed Principle**: Extend functionality without modifying existing code
- **Avoid Class Explosion**: Don't need subclasses for every combination
- **Incremental**: Add functionality incrementally

## Disadvantages

- **Complexity**: Many small objects can be hard to understand
- **Order Dependency**: Order of decorators can matter (encrypt then compress ≠ compress then encrypt)
- **Identical Objects**: Hard to distinguish between identically decorated objects
- **Instantiation Complexity**: Complicated instantiation code with many decorators
- **Debugging Difficulty**: Stack of decorators can make debugging harder
- **Type Checking**: Harder to check for specific decorator types

## Decorator vs Other Patterns

| Pattern | Purpose | Structure | Key Difference |
|---------|---------|-----------|----------------|
| **Decorator** | Add responsibilities dynamically | Wraps one object | Keeps same interface, adds behavior |
| **Adapter** | Make incompatible interfaces work | Wraps one object | Changes interface |
| **Proxy** | Control access to object | Wraps one object | Same interface, controls access |
| **Composite** | Treat objects uniformly | Tree structure | Groups multiple objects |
| **Strategy** | Change algorithm | Holds strategy object | Changes behavior, not structure |

**Key Distinction**: 
- **Decorator adds to what object does** (new responsibilities)
- **Adapter changes how you interact with it** (interface conversion)
- **Proxy controls when/how you access it** (access control)

## Decorator vs Inheritance

| Aspect | Decorator | Inheritance |
|--------|-----------|-------------|
| **Timing** | Runtime | Compile-time |
| **Flexibility** | Mix and match | Fixed hierarchy |
| **Classes Needed** | Few decorators | Many subclasses |
| **Combinations** | Any combination | Must create subclass for each |
| **Example** | 4 features = 4 decorators | 4 features = 16 subclasses (2^4) |

## Running the Example

```bash
# Navigate to the decorator directory
cd structural-patterns/decorator

# Run the example
go run main.go
```

## Expected Output

```
=== Decorator Pattern Demo ===

--- Example 1: Coffee Shop ---

** Simple Coffee **
Order: Simple Coffee
Total Cost: $2.00

** Coffee with Milk **
Order: Simple Coffee, Milk
Total Cost: $2.50

** Coffee with Milk and Sugar **
Order: Simple Coffee, Milk, Sugar
Total Cost: $2.75

** Deluxe Coffee (Milk, Sugar, Whipped Cream, Caramel) **
Order: Simple Coffee, Milk, Sugar, Whipped Cream, Caramel
Total Cost: $4.10

--- Example 2: Data Source with Encryption and Compression ---

** Plain File Write/Read **
[FileDataSource] Writing to data.txt: Hello World
[FileDataSource] Reading from data.txt
Retrieved: Hello World

** Encrypted File Write/Read **
[EncryptionDecorator] Encrypting data
[FileDataSource] Writing to encrypted.txt: ENCRYPTED[Secret Message]
[FileDataSource] Reading from encrypted.txt
[EncryptionDecorator] Decrypting data
Retrieved: Secret Message

** Compressed and Encrypted File Write/Read **
[CompressionDecorator] Compressing data
[EncryptionDecorator] Encrypting data
[FileDataSource] Writing to secure.txt: ENCRYPTED[COMPRESSED[Confidential Data]]
[FileDataSource] Reading from secure.txt
[EncryptionDecorator] Decrypting data
[CompressionDecorator] Decompressing data
Retrieved: Confidential Data

--- Example 3: Multi-Channel Notifications ---

** Email Only **
[Email] Sending: Server is up!

** Email + SMS **
[Email] Sending: Server is down!
[SMS] Also sending: Server is down!

** Email + SMS + Slack **
[Email] Sending: Deployment complete!
[SMS] Also sending: Deployment complete!
[Slack] Also sending: Deployment complete!

** All Channels with Logging **
[Log] 14:23:45 - Sending notification
[Email] Sending: Critical alert!
[SMS] Also sending: Critical alert!
[Slack] Also sending: Critical alert!
[Facebook] Also sending: Critical alert!
[Log] 14:23:45 - Notification sent

--- Summary ---
Decorators add responsibilities to objects dynamically
Multiple decorators can be stacked/chained
Decorators maintain the same interface as the component
Provides flexible alternative to subclassing
```

## Key Takeaways

- **Dynamic Enhancement**: Add functionality at runtime, not compile-time
- **Same Interface**: Decorators maintain the same interface as wrapped objects
- **Wrapping Chain**: Decorators can wrap other decorators recursively
- **Transparent**: Client code doesn't need to know about decorators
- **Composition Over Inheritance**: Uses composition instead of inheritance
- **Order Matters**: Order of decorator wrapping can affect behavior
- **Single Responsibility**: Each decorator adds one specific feature

## Implementation Patterns

### Classic Decorator
Uses a base decorator class:
```go
type Decorator struct {
    component Component
}

type ConcreteDecorator struct {
    Decorator
}
```

### Direct Wrapping
Each decorator directly wraps the interface:
```go
type ConcreteDecorator struct {
    wrapped Component
}
```

This implementation uses **direct wrapping** with a base decorator struct for code reuse.

## Order Dependency

The order of decorators can be important:

```go
// Compress then encrypt (recommended for security)
encrypted := NewEncryptionDecorator(
    NewCompressionDecorator(source)
)

// Encrypt then compress (less efficient)
compressed := NewCompressionDecorator(
    NewEncryptionDecorator(source)
)
```

**Best Practice**: Document if order matters!

## Design Considerations

### Interface Design
- Keep interfaces focused and cohesive
- Decorators should enhance, not change behavior fundamentally

### Decorator Initialization
- Consider using builder pattern for complex decorator chains
- Provide factory methods for common decorator combinations

### Performance
- Be aware of overhead with many layers
- Consider caching if decorators are expensive

### Null Object
- Handle nil wrapped objects gracefully
- Consider providing a null object implementation

## Common Mistakes

1. **Changing Interface**: Decorators should maintain the same interface
2. **Too Many Decorators**: Over-decorating makes code hard to understand
3. **Ignoring Order**: Not considering that decorator order can matter
4. **Stateful Decorators**: Be careful with decorators that maintain state
5. **Not Delegating**: Forgetting to call the wrapped object's methods
6. **Type Casting**: Trying to cast to concrete types defeats the purpose

## Best Practices

1. **Keep Decorators Simple**: Each should add one clear responsibility
2. **Document Order Dependencies**: Clearly note if order matters
3. **Provide Factories**: Create factory methods for common combinations
4. **Maintain Interface**: Never break the component's interface contract
5. **Consider Performance**: Be aware of overhead with deep decorator chains
6. **Test Combinations**: Test various decorator combinations
7. **Use Descriptive Names**: Name decorators clearly (e.g., EncryptionDecorator)

## Real-World Examples

- **Java I/O**: BufferedReader, InputStreamReader wrapping streams
- **Go http.Handler**: Middleware wrapping handlers
- **Python Decorators**: @decorator syntax (function decorators)
- **Logging Libraries**: Adding context, formatting, filtering to loggers
- **Web Frameworks**: Middleware in Express.js, Django, etc.
- **GUI Toolkits**: Adding borders, scrollbars to components

## Functional Programming Alternative

In languages with first-class functions, decorators can be implemented as higher-order functions:

```go
func LoggingDecorator(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Request: %s", r.URL.Path)
        handler.ServeHTTP(w, r)
    })
}
```

This is the approach used by Go's http.Handler middleware.

## When NOT to Use Decorator

- When you only need one or two variations (simple inheritance may suffice)
- When the interface is complex and decorators would be cumbersome
- When you need to add methods, not just enhance existing ones (use subclassing)
- When performance overhead of wrapping is unacceptable
- When the component's interface is unstable (decorators become brittle)

## Decorator vs Proxy

Both wrap objects, but:
- **Decorator**: Adds new functionality
- **Proxy**: Controls access or adds infrastructure concerns (lazy loading, access control, logging)

The line can blur, but intent differs:
- Decorator: "Give me X with extra features"
- Proxy: "Give me controlled access to X"
