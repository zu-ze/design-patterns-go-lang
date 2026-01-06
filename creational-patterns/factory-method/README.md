# Factory Method Pattern

## Overview

The Factory Method Pattern is a creational design pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created. Instead of calling a constructor directly, you call a factory method that returns an object.

## Problem It Solves

When you have a class that needs to create objects, but you don't know in advance which exact class of object to create, or you want to delegate the responsibility of instantiation to subclasses. This pattern promotes loose coupling by eliminating the need to bind application-specific classes into your code.

## When to Use

- When a class can't anticipate the type of objects it needs to create
- When a class wants its subclasses to specify the objects it creates
- When you want to provide a library of products and only reveal their interfaces, not implementations
- When you want to delegate the instantiation logic to child classes
- When you need to manage or encapsulate object creation

## Structure

The pattern consists of:

1. **Product Interface** (`Notification`): Defines the interface of objects the factory method creates
2. **Concrete Products** (`EmailNotification`, `SMSNotification`, `PushNotification`): Specific implementations of the Product interface
3. **Factory Interface** (`NotificationFactory`): Declares the factory method that returns Product objects
4. **Concrete Factories** (`EmailNotificationFactory`, `SMSNotificationFactory`, `PushNotificationFactory`): Implement the factory method to create specific Product instances

## Implementation Details

In this example, we implement a notification system:

- **Notification Interface**: Defines `Send()` and `GetType()` methods
- **Concrete Notifications**: Email, SMS, and Push notifications, each implementing the Notification interface
- **NotificationFactory Interface**: Declares the `CreateNotification()` factory method
- **Concrete Factories**: Each factory creates a specific type of notification with its required configuration

The client code (`sendNotification` function) works with factories through the common interface, without knowing the concrete classes of notifications being created.

## Use Cases

1. **UI Component Libraries**: Creating different types of buttons, dialogs, or windows for different operating systems (Windows, macOS, Linux)
2. **Database Connections**: Creating connections to different database types (MySQL, PostgreSQL, MongoDB) based on configuration
3. **Document Processors**: Creating parsers for different document formats (PDF, Word, Excel)
4. **Payment Processing**: Creating payment handlers for different payment methods (Credit Card, PayPal, Cryptocurrency)
5. **Logging Systems**: Creating loggers that output to different destinations (file, console, cloud service)
6. **Notification Systems**: Creating different notification channels as shown in this example

## Advantages

- **Open/Closed Principle**: You can introduce new types of products without breaking existing client code
- **Single Responsibility Principle**: Product creation code is in one place, making it easier to maintain
- **Loose Coupling**: Client code works with interfaces rather than concrete classes
- **Flexibility**: Easy to extend with new product types
- **Encapsulation**: Complex object creation logic is hidden from client code

## Disadvantages

- **Complexity**: Can make code more complicated with additional classes and interfaces
- **Overhead**: May introduce unnecessary abstraction if you only have a few product types
- **Learning Curve**: New developers need to understand the factory hierarchy

## Running the Example

```bash
# Navigate to the factory-method directory
cd factory-method

# Run the example
go run main.go
```

## Expected Output

```
=== Factory Method Pattern Demo ===

Created Email notification
[EMAIL] Sending to user@example.com: Your order has been shipped!

Created SMS notification
[SMS] Sending to +1234567890: Your verification code is 123456

Created Push notification
[PUSH] Sending to device device-abc-123: You have a new message
```

## Key Takeaways

- Use Factory Method when you need flexibility in object creation
- The pattern delegates instantiation to subclasses or factory implementations
- It promotes loose coupling between client code and concrete classes
- Perfect for scenarios where the exact types of objects aren't known until runtime
- Balances the trade-off between flexibility and complexity
