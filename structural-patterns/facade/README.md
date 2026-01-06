# Facade Pattern

## Overview

The Facade Pattern is a structural design pattern that provides a simplified, unified interface to a complex subsystem. It defines a higher-level interface that makes the subsystem easier to use by wrapping a complicated subsystem with a simpler interface. The facade doesn't hide the subsystem - it just provides a more convenient way to access it.

## Problem It Solves

When a system is complex with many interdependent classes, or when you need to interact with a sophisticated library/framework, the learning curve is steep and code becomes cluttered with subsystem details. The Facade pattern solves this by providing a simple interface that handles the complexity internally, making the subsystem much easier to use for common tasks.

## When to Use

- When you want to provide a simple interface to a complex subsystem
- When there are many dependencies between clients and implementation classes
- When you want to layer your subsystems with facades at different levels
- When you want to reduce coupling between subsystems and clients
- When you need to wrap a poorly designed collection of APIs with a single well-designed API
- When you want to provide a default view of a subsystem for most users while still allowing power users direct access

## Structure

The pattern consists of:

1. **Facade**: Provides simple methods that delegate to subsystem classes
2. **Subsystem Classes**: Implement subsystem functionality and handle work assigned by Facade
3. **Client**: Uses the Facade instead of calling subsystem classes directly

## Key Characteristics

### Simplified Interface
The facade provides a simpler interface than the subsystem:
```go
// Instead of:
cpu.Freeze()
memory.Load(0x00, data)
cpu.Jump(0x00)
hardDrive.Read(sector, size)
cpu.Execute()

// Just call:
computer.Start()
```

### Doesn't Hide Subsystem
Advanced users can still access subsystem classes directly if needed.

### One-to-Many Relationship
One facade can coordinate many subsystem objects.

## Implementation Details

The example demonstrates three facade scenarios:

### 1. Computer Boot System
- **Subsystem**: CPU, Memory, HardDrive (complex boot process)
- **Facade**: ComputerFacade with simple `Start()` method
- **Complexity Hidden**: Multiple subsystem calls coordinated internally

**Without Facade**: Client must understand boot sequence, memory addresses, sector positions
**With Facade**: Client just calls `Start()`

### 2. Video Conversion System
- **Subsystem**: VideoFile, CodecFactory, BitrateReader, AudioMixer
- **Facade**: VideoConverter with simple `Convert(filename, format)` method
- **Complexity Hidden**: Codec detection, bitrate conversion, audio sync

**Without Facade**: Client must understand video codecs, bitrate manipulation
**With Facade**: Client just calls `Convert("video.ogg", "mp4")`

### 3. Banking System
- **Subsystem**: Account, SecurityValidator, LedgerService, NotificationService
- **Facade**: BankingFacade with `Withdraw()`, `Deposit()`, `GetBalance()` methods
- **Complexity Hidden**: Security validation, ledger recording, notifications

**Without Facade**: Client must coordinate security, transactions, logging, notifications
**With Facade**: Client just calls `Withdraw(account, pin, amount)`

## Use Cases

1. **Library/Framework Wrappers**: Simplifying complex third-party libraries
   - Wrapping database libraries (ORM facades)
   - Simplifying HTTP client libraries
   - Wrapping cloud service SDKs (AWS, Azure, GCP)

2. **Legacy System Integration**: Providing modern interface to old systems
   - SOAP service facade for REST clients
   - Mainframe system facades
   - Legacy database facades

3. **Microservices**: Aggregating multiple service calls
   - API Gateway as a facade to microservices
   - Backend-for-Frontend (BFF) pattern
   - Composite services

4. **Hardware Abstraction**: Simplifying hardware interactions
   - Device driver interfaces
   - Printer subsystem facades
   - Graphics card facades

5. **Complex Workflows**: Simplifying multi-step processes
   - Order processing (payment, inventory, shipping)
   - User registration (validation, database, email)
   - Report generation (data collection, processing, formatting)

6. **Testing**: Creating test facades for complex systems
   - Test database facades
   - Mock service facades
   - Simplified test utilities

## Advantages

- **Simplified Interface**: Makes complex subsystems easier to use
- **Reduced Coupling**: Clients depend on facade, not subsystem details
- **Layered Architecture**: Can create layers of facades for different abstraction levels
- **Easier Maintenance**: Changes to subsystem don't affect clients using facade
- **Better Organization**: Groups related operations in one place
- **Improved Readability**: Client code is cleaner and more understandable
- **Flexibility**: Subsystem still accessible for power users

## Disadvantages

- **Limited Functionality**: Facade may not expose all subsystem features
- **God Object Risk**: Facade can become a god object if it does too much
- **Additional Layer**: Adds another layer of indirection
- **Tight Coupling**: Facade becomes tightly coupled to subsystem
- **Over-Simplification**: May hide important details that some clients need
- **Maintenance**: Facade needs updates when subsystem changes

## Facade vs Other Patterns

| Pattern | Purpose | Key Difference |
|---------|---------|----------------|
| **Facade** | Simplify complex subsystem | Provides simplified interface, doesn't hide subsystem |
| **Adapter** | Make incompatible interfaces compatible | Changes interface to match expected one |
| **Proxy** | Control access to object | Same interface, controls access/lifecycle |
| **Mediator** | Reduce coupling between objects | Coordinates communication between colleagues |
| **Decorator** | Add responsibilities dynamically | Same interface, adds functionality |

**Key Distinctions**:
- **Facade simplifies** (many interfaces → one simple interface)
- **Adapter translates** (one interface → another interface)
- **Proxy controls** (same interface, adds control layer)

## Facade vs Adapter

| Aspect | Facade | Adapter |
|--------|--------|---------|
| **Intent** | Simplify | Make compatible |
| **Interfaces** | New simplified interface | Converts existing interface |
| **Complexity** | Simplifies complex subsystem | Wraps single class typically |
| **Subsystem Access** | Still accessible | Usually hidden |
| **Example** | `computer.Start()` for boot process | Convert Celsius API to Fahrenheit |

## Running the Example

```bash
# Navigate to the facade directory
cd structural-patterns/facade

# Run the example
go run main.go
```

## Expected Output

```
=== Facade Pattern Demo ===

--- Example 1: Computer Boot System ---

Without Facade, client would need to:
- Create CPU, Memory, HardDrive
- Call CPU.Freeze()
- Call Memory.Load() with correct parameters
- Call CPU.Jump() to correct position
- Call HardDrive.Read() with correct sector
- Call Memory.Load() again with boot data
- Call CPU.Execute()

With Facade:

[ComputerFacade] Starting computer...

[CPU] Freezing processor
[Memory] Loading 'BOOT_SECTOR' at position 0
[CPU] Jumping to position 0
[HardDrive] Reading 1024 bytes from sector 0
[Memory] Loading 'BOOT_DATA' at position 256
[CPU] Executing instructions

[ComputerFacade] Computer started successfully!

--- Example 2: Video Conversion ---

Without Facade, client would need to:
- Understand video codecs (MPEG4, Ogg, etc.)
- Create CodecFactory, BitrateReader, AudioMixer
- Extract source codec
- Create destination codec
- Read and convert bitrate
- Fix audio synchronization

With Facade:

[VideoConverter] Starting conversion of video.ogg to mp4 format

[VideoFile] Loading file: video.ogg
[CodecFactory] Extracting codec from file
[BitrateReader] Reading bitrate for Ogg codec
[BitrateReader] Converting buffer using MPEG4 codec
[AudioMixer] Fixing audio synchronization

[VideoConverter] Conversion complete!

--- Example 3: Banking System ---

Without Facade, client would need to:
- Validate security with SecurityValidator
- Get Account object
- Perform debit/credit operations
- Record transaction in LedgerService
- Send notification via NotificationService

With Facade:

[BankingFacade] Checking balance for ACC001

[SecurityValidator] Validating account ACC001

[BankingFacade] Current balance: $1000.00

[BankingFacade] Processing withdrawal of $200.00 from ACC001

[SecurityValidator] Validating account ACC001
[Account] Debited $200.00. New balance: $800.00
[LedgerService] Recording WITHDRAWAL transaction of $200.00 for account ACC001
[NotificationService] Sending notification to ACC001: Withdrawal of $200.00 completed

[BankingFacade] Withdrawal successful!

[BankingFacade] Processing deposit of $150.00 to ACC001

[SecurityValidator] Validating account ACC001
[Account] Credited $150.00. New balance: $950.00
[LedgerService] Recording DEPOSIT transaction of $150.00 for account ACC001
[NotificationService] Sending notification to ACC001: Deposit of $150.00 completed

[BankingFacade] Deposit successful!

--- Summary ---
Facade provides a simplified interface to complex subsystems
Clients interact with one simple interface instead of many complex ones
Reduces coupling between client and subsystem
Subsystems remain accessible for advanced users who need direct access
```

## Key Takeaways

- **Simplification**: Provides a simple interface to complex subsystems
- **Not Hiding**: Doesn't prevent direct access to subsystem (unlike encapsulation)
- **Coordination**: Coordinates multiple subsystem objects
- **Decoupling**: Reduces dependencies between clients and subsystem
- **Common Use Cases**: Most common for simplifying third-party libraries and complex APIs
- **Multiple Facades**: Can have multiple facades for different use cases
- **Layering**: Can create layers of facades at different abstraction levels

## Design Considerations

### Facade Scope
- **Narrow Facade**: Focuses on specific use cases (recommended)
- **Wide Facade**: Tries to cover everything (becomes god object)

### Subsystem Access
- **Transparent**: Subsystem classes remain accessible
- **Opaque**: Facade is the only access point (rare)

### Multiple Facades
Create different facades for different client needs:
```go
type SimpleBankingFacade struct { } // Basic operations
type AdvancedBankingFacade struct { } // Advanced operations
type AdminBankingFacade struct { } // Administrative operations
```

### Facade Composition
Facades can use other facades:
```go
type HighLevelFacade struct {
    subsystemFacade1 *SubsystemFacade1
    subsystemFacade2 *SubsystemFacade2
}
```

## Common Mistakes

1. **God Object**: Making facade do everything
2. **Business Logic**: Putting business logic in facade (should delegate)
3. **Hiding Too Much**: Making subsystem completely inaccessible
4. **Not Enough Abstraction**: Facade that's just as complex as subsystem
5. **Tight Coupling**: Making facade depend on too many concrete classes
6. **Stateful Facade**: Adding unnecessary state to facade

## Best Practices

1. **Keep It Simple**: Facade should be simpler than using subsystem directly
2. **Delegate, Don't Implement**: Facade coordinates but doesn't implement logic
3. **One Purpose**: Each facade should serve one clear purpose
4. **Allow Direct Access**: Don't make facade the only way to access subsystem
5. **Document Well**: Clearly document what complexity is hidden
6. **Consider Versioning**: When API changes, consider versioned facades
7. **Stateless When Possible**: Prefer stateless facades (just coordinate)

## Layered Facades

Create abstraction layers with facades:

```
Client
  ↓
HighLevelFacade (Application operations)
  ↓
MiddleLevelFacade (Domain operations)
  ↓
LowLevelFacade (Technical operations)
  ↓
Subsystem (Complex implementation)
```

Each layer provides appropriate abstraction for its level.

## Real-World Examples

- **Database ORMs**: Sequelize, Hibernate, GORM hide SQL complexity
- **HTTP Clients**: Axios, Requests library simplify HTTP operations
- **Cloud SDKs**: AWS SDK facades simplify AWS service interactions
- **jQuery**: Facade over browser DOM manipulation
- **Spring Framework**: Many facade classes (JdbcTemplate, RestTemplate)
- **Standard Libraries**: Many standard library packages are facades (e.g., Go's `http` package)

## Testing with Facades

Facades make testing easier:

```go
// Test with mock facade
type MockBankingFacade struct {
    WithdrawCalled bool
}

func (m *MockBankingFacade) Withdraw(account, pin string, amount float64) bool {
    m.WithdrawCalled = true
    return true
}
```

## When NOT to Use Facade

- When subsystem is already simple
- When you need full control and flexibility of subsystem
- When facade would just duplicate subsystem interface
- When creating facade would be more work than using subsystem directly
- When subsystem is likely to change frequently (facade becomes burden)

## Facade and Dependency Injection

Combine with dependency injection for testability:

```go
type PaymentProcessor struct {
    paymentGateway PaymentGatewayFacade // Injected facade
}

func NewPaymentProcessor(facade PaymentGatewayFacade) *PaymentProcessor {
    return &PaymentProcessor{paymentGateway: facade}
}
```

## Evolutionary Design

Start without facade, add when complexity grows:

1. **Phase 1**: Use subsystem directly
2. **Phase 2**: Notice repetitive patterns
3. **Phase 3**: Create facade for common operations
4. **Phase 4**: Migrate clients to facade gradually

Don't over-engineer with facades upfront - add them when complexity warrants it.
