# Bridge Pattern

## Overview

The Bridge Pattern is a structural design pattern that separates an abstraction from its implementation so that the two can vary independently. It involves an interface which acts as a bridge between the abstraction and implementation classes. Rather than creating a class hierarchy, it uses composition to connect the abstraction to the implementation.

## Problem It Solves

When you have multiple dimensions of variation in a system, creating a class for every combination leads to an explosion of subclasses. For example, if you have different types of shapes (Circle, Square) and different rendering methods (Vector, Raster), you'd need CircleVector, CircleRaster, SquareVector, SquareRaster - 4 classes. Add a Triangle and a 3D renderer, and you have 9 classes. The Bridge pattern solves this by separating these concerns.

## When to Use

- When you want to avoid a permanent binding between an abstraction and its implementation
- When both the abstraction and implementation should be extensible through subclasses
- When changes in the implementation should not affect clients
- When you have a proliferation of classes resulting from multiple dimensions of variation
- When you want to share an implementation among multiple objects
- When you need to switch implementations at runtime

## Structure

The pattern consists of:

1. **Abstraction**: Defines the abstraction's interface and maintains a reference to the Implementor
2. **Refined Abstraction**: Extends the abstraction interface
3. **Implementor**: Defines the interface for implementation classes
4. **Concrete Implementor**: Implements the Implementor interface

## Key Difference: Bridge vs Adapter

| Aspect | Bridge | Adapter |
|--------|--------|---------|
| **Intent** | Design upfront for flexibility | Fix incompatibility after the fact |
| **When Applied** | During initial design | When integrating existing code |
| **Purpose** | Decouple abstraction from implementation | Make incompatible interfaces work together |
| **Structure** | Both sides designed to vary | Adapts existing interface to expected one |
| **Example** | Remote controls for various devices | Converting legacy API to new interface |

**Bridge is proactive** (designed from the start), **Adapter is reactive** (fixes existing problems).

## Implementation Details

The example demonstrates two bridge implementations:

### 1. Remote Control Bridge
- **Abstraction Hierarchy**: Remote, AdvancedRemote
- **Implementation Hierarchy**: Device (TV, Radio)
- **Bridge**: Remote holds reference to Device interface
- **Benefit**: Can create any remote type that works with any device

**Without Bridge**: Would need TVRemote, RadioRemote, TVAdvancedRemote, RadioAdvancedRemote (4 classes)
**With Bridge**: Remote + AdvancedRemote + TV + Radio (4 classes, but extensible)

### 2. Notification System Bridge
- **Abstraction Hierarchy**: Notification, UrgentNotification, ScheduledNotification
- **Implementation Hierarchy**: MessageSender (Email, SMS, Push)
- **Bridge**: Notification holds reference to MessageSender interface
- **Benefit**: Any notification type can use any message sender

## Use Cases

1. **GUI Frameworks**: Separating window abstractions from platform-specific implementations
   - Window abstraction works with Windows, macOS, Linux implementations
   - Button, Dialog abstractions work with any platform renderer

2. **Database Drivers**: Separating database abstraction from specific database implementations
   - Repository abstraction works with MySQL, PostgreSQL, MongoDB implementations
   - Query builder works with any database driver

3. **Graphics Rendering**: Separating shape abstractions from rendering methods
   - Shape abstractions (Circle, Square) work with Vector, Raster, 3D renderers
   - Drawing operations work with any rendering engine

4. **Device Drivers**: Separating device abstraction from platform-specific code
   - Printer abstraction works with Windows, Linux, macOS drivers
   - Network device abstraction works with any network stack

5. **Media Players**: Separating player controls from media format handlers
   - Player controls work with MP3, MP4, AVI decoders
   - Playlist manager works with any media format

6. **Payment Processing**: Separating payment abstraction from payment gateways
   - Payment abstraction works with Stripe, PayPal, Square implementations
   - Refund logic works with any payment provider

## Advantages

- **Decoupling**: Abstraction and implementation are independent
- **Extensibility**: Easy to extend both abstraction and implementation hierarchies
- **Hiding Implementation Details**: Client code only sees the abstraction
- **Single Responsibility Principle**: Abstraction focuses on high-level logic, implementation on platform details
- **Open/Closed Principle**: Can introduce new abstractions and implementations independently
- **Runtime Binding**: Can switch implementations at runtime
- **Reduced Class Explosion**: Avoids exponential growth of subclasses

## Disadvantages

- **Complexity**: Increases code complexity with additional layers
- **Indirection**: More difficult to understand the flow
- **Initial Overhead**: Requires more upfront design
- **May Be Overkill**: For simple systems with single dimension of variation
- **Learning Curve**: Can be harder to understand than direct inheritance

## Bridge in Action: Class Explosion Prevention

**Problem**: Without Bridge pattern
```
RemoteControl
├── TVRemote
├── RadioRemote
├── SpeakerRemote
└── Advanced variants
    ├── TVAdvancedRemote
    ├── RadioAdvancedRemote
    └── SpeakerAdvancedRemote
```
7+ classes and growing exponentially!

**Solution**: With Bridge pattern
```
Remote (Abstraction)          Device (Implementation)
├── Remote                    ├── TV
└── AdvancedRemote            ├── Radio
                              └── Speaker
```
5 classes, linear growth!

## Running the Example

```bash
# Navigate to the bridge directory
cd structural-patterns/bridge

# Run the example
go run main.go
```

## Expected Output

```
=== Bridge Pattern Demo ===

--- Example 1: Remote Control and Devices ---

** Using Basic Remote with TV **

[Remote] Checking device status...
[TV] Status: OFF | Volume: 50% | Channel: 1

[Remote] Toggling power...
[TV] TV is now ON

[Remote] Volume up...
[TV] Volume set to 60%

[Remote] Volume up...
[TV] Volume set to 70%

[Remote] Channel up...
[TV] Channel set to 2

[Remote] Checking device status...
[TV] Status: ON | Volume: 70% | Channel: 2

** Using Advanced Remote with Radio **

[Remote] Checking device status...
[Radio] Status: OFF | Volume: 30% | Frequency: 88.5 FM

[Remote] Toggling power...
[Radio] Radio is now ON

[Remote] Volume up...
[Radio] Volume set to 40%

[AdvancedRemote] Setting channel to 1015...
[Radio] Frequency set to 101.5 FM

[AdvancedRemote] Muting device...
[Radio] Volume set to 0%

[Remote] Checking device status...
[Radio] Status: ON | Volume: 0% | Frequency: 101.5 FM

** Switching Advanced Remote to TV **

[Remote] Toggling power...
[TV] TV is now ON

[AdvancedRemote] Setting channel to 42...
[TV] Channel set to 42

[Remote] Checking device status...
[TV] Status: ON | Volume: 70% | Channel: 42

--- Example 2: Notification System ---

** Regular Notifications **
[Email] Sending via smtp.example.com: Welcome to our service!
[SMS] Sending via twilio.com: Your verification code is 123456
[Push] Sending via firebase.com: You have a new message

** Urgent Notifications **
[Email] Sending via smtp.example.com: [URGENT] Server is down!
[SMS] Sending via twilio.com: [URGENT] Security alert detected

** Scheduled Notifications **
[Email] Sending via smtp.example.com: [Scheduled for 2024-12-25 09:00] Merry Christmas!
[Push] Sending via firebase.com: [Scheduled for 2024-01-01 00:00] Happy New Year!

--- Summary ---
Bridge pattern separates abstraction from implementation
Remotes (abstraction) work with any device (implementation)
Notifications (abstraction) work with any sender (implementation)
Both hierarchies can evolve independently
```

## Key Takeaways

- **Composition Over Inheritance**: Uses composition to connect abstraction and implementation
- **Two Hierarchies**: Maintains separate hierarchies for abstraction and implementation
- **Independent Evolution**: Both hierarchies can evolve without affecting each other
- **Design Upfront**: Unlike Adapter, Bridge is designed into the system from the start
- **Prevents Class Explosion**: Avoids exponential growth when you have multiple dimensions of variation
- **Runtime Flexibility**: Can switch implementations at runtime by changing the reference
- **Use When**: You have or anticipate having multiple dimensions of variation

## Comparison with Similar Patterns

### Bridge vs Strategy

| Aspect | Bridge | Strategy |
|--------|--------|----------|
| **Focus** | Structural (object structure) | Behavioral (algorithm selection) |
| **Purpose** | Separate abstraction from implementation | Encapsulate interchangeable algorithms |
| **Hierarchies** | Two separate hierarchies | One algorithm hierarchy |
| **Variation** | Platform/implementation details | Algorithm/behavior |

### Bridge vs Adapter vs Decorator

| Pattern | Purpose | Timing | Interface |
|---------|---------|--------|-----------|
| **Bridge** | Separate abstraction from implementation | Design time | Creates new interface |
| **Adapter** | Make incompatible interfaces compatible | After implementation | Adapts to existing |
| **Decorator** | Add responsibilities dynamically | Runtime | Keeps same interface |

## Best Practices

1. **Identify Dimensions Early**: Identify multiple dimensions of variation during design
2. **Keep Abstractions Simple**: Abstraction should focus on high-level operations
3. **Define Clear Interfaces**: Both abstraction and implementor need clear, stable interfaces
4. **Document the Bridge**: Clearly document which side is abstraction vs implementation
5. **Consider Complexity**: Don't use Bridge if you only have one dimension of variation
6. **Use Dependency Injection**: Pass implementor through constructor for flexibility
7. **Test Independently**: Test abstractions and implementations separately

## Common Mistakes

1. **Confusing with Adapter**: Remember Bridge is designed upfront, Adapter is retrofitted
2. **Over-Engineering**: Using Bridge when simple inheritance would suffice
3. **Tight Coupling**: Letting abstraction depend on concrete implementation details
4. **Ignoring Single Dimension**: Using Bridge when you only have one dimension to vary
5. **Not Using Interfaces**: Forgetting to define proper implementor interfaces

## Real-World Examples

- **Java AWT/Swing**: Window toolkit abstraction bridged to platform-specific implementations
- **Database Drivers**: JDBC provides abstraction bridged to database-specific drivers
- **Go database/sql**: Generic database interface bridged to specific driver implementations
- **Graphics APIs**: OpenGL/DirectX abstractions bridged to GPU implementations
- **Logging Frameworks**: Abstract logging interface bridged to specific output mechanisms

## When NOT to Use Bridge

- When you only have one dimension of variation (use simple inheritance)
- When the abstraction and implementation won't vary independently
- When the added complexity outweighs the benefits
- When the system is simple and unlikely to change
- When you're dealing with existing incompatible interfaces (use Adapter instead)
