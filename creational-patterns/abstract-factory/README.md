# Abstract Factory Pattern

## Overview

The Abstract Factory Pattern is a creational design pattern that provides an interface for creating families of related or dependent objects without specifying their concrete classes. It's a factory of factories - you use it when you need to create groups of objects that are designed to work together.

## Problem It Solves

When you need to create multiple related objects that must be used together (a family of products), and you want to ensure that the products from one family are compatible with each other. The pattern helps avoid mixing incompatible products from different families.

## Difference from Factory Method

- **Factory Method**: Creates one type of product through inheritance. Focuses on creating a single product.
- **Abstract Factory**: Creates families of related products through object composition. Focuses on creating multiple related products that work together.

## When to Use

- When your system needs to work with multiple families of related products
- When you want to enforce constraints that products from the same family should be used together
- When you want to provide a library of products and reveal only their interfaces, not implementations
- When the system should be independent of how its products are created
- When you want to switch between different product families at runtime

## Structure

The pattern consists of:

1. **Abstract Products** (`Button`, `Checkbox`, `Input`): Interfaces for different types of products
2. **Concrete Products** (`LightButton`, `DarkButton`, etc.): Specific implementations of products for each family
3. **Abstract Factory** (`UIFactory`): Interface declaring creation methods for each abstract product
4. **Concrete Factories** (`LightThemeFactory`, `DarkThemeFactory`): Implement the abstract factory interface to create concrete products from a specific family
5. **Client** (`Application`): Works only with abstract factories and products through their interfaces

## Implementation Details

In this example, we implement a UI theming system:

- **Product Interfaces**: Button, Checkbox, and Input define the operations available on UI components
- **Product Families**: Light Theme and Dark Theme, each with their own look and behavior
- **Factories**: LightThemeFactory creates light-themed components, DarkThemeFactory creates dark-themed components
- **Application**: Uses a factory to create all UI components, ensuring they're all from the same theme

The key insight is that the Application class doesn't know which concrete theme it's using - it just asks the factory for components, and they're guaranteed to be compatible.

## Use Cases

1. **Cross-Platform UI Frameworks**: Creating UI components for different operating systems (Windows, macOS, Linux) where each OS has its own look and feel
2. **Theme Systems**: Implementing light/dark themes or multiple design systems in applications
3. **Database Access**: Creating related database objects (Connection, Command, Transaction) for different database vendors (SQL Server, Oracle, PostgreSQL)
4. **Document Generation**: Creating document elements (Paragraph, Table, Image) for different formats (PDF, HTML, Word)
5. **Game Development**: Creating related game objects (Character, Weapon, Vehicle) for different game levels or environments (Medieval, Sci-Fi, Fantasy)
6. **Cloud Providers**: Creating related cloud resources (Storage, Compute, Network) for different providers (AWS, Azure, GCP)

## Advantages

- **Consistency**: Ensures that products from the same family are used together
- **Isolation**: Client code is isolated from concrete product classes
- **Open/Closed Principle**: Easy to introduce new product families without modifying existing code
- **Single Responsibility**: Product creation code is centralized in factories
- **Type Safety**: Compile-time guarantee that products from the same family are used together

## Disadvantages

- **Complexity**: Adds many new interfaces and classes
- **Rigidity**: Adding new product types requires changing all factory interfaces and implementations
- **Overkill**: May be too complex if you don't need to support multiple product families
- **Parallel Hierarchies**: Requires maintaining parallel class hierarchies for products and factories

## Running the Example

```bash
# Navigate to the abstract-factory directory
cd abstract-factory

# Run the example
go run main.go
```

## Expected Output

```
=== Abstract Factory Pattern Demo ===

--- Light Theme Application ---
[Light Button] ☐ with white background and dark text
[Light Button] Clicked with subtle shadow effect

[Light Checkbox] ☐ with light gray border
[Light Checkbox] Toggled with smooth transition

[Light Input] ___ with white background and thin border
[Light Input] Value set to: Hello World (dark text on white)

--- Dark Theme Application ---
[Dark Button] ☐ with dark background and light text
[Dark Button] Clicked with neon glow effect

[Dark Checkbox] ☐ with bright border on dark background
[Dark Checkbox] Toggled with glowing animation

[Dark Input] ___ with dark background and bright border
[Dark Input] Value set to: Hello World (light text on dark)
```

## Key Takeaways

- Use Abstract Factory when you need to create families of related objects
- The pattern ensures compatibility between products from the same family
- Client code works with interfaces, making it independent of concrete implementations
- Perfect for scenarios with multiple product variants that must be consistent
- More complex than Factory Method but provides stronger guarantees about product compatibility
- Trade-off: Flexibility in adding new families vs. difficulty in adding new product types

## Comparison with Factory Method

| Aspect | Factory Method | Abstract Factory |
|--------|----------------|------------------|
| Focus | Single product | Family of products |
| Implementation | Inheritance | Object composition |
| Complexity | Simpler | More complex |
| Products | One type | Multiple related types |
| Use case | Create one thing in different ways | Create related things consistently |
