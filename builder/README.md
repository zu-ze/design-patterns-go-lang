# Builder Pattern

## Overview

The Builder Pattern is a creational design pattern that lets you construct complex objects step by step. It allows you to produce different types and representations of an object using the same construction code. The pattern separates the construction of a complex object from its representation.

## Problem It Solves

When you need to create complex objects with many optional parameters or configurations, constructors can become unwieldy with too many parameters (telescoping constructor anti-pattern). The Builder pattern provides a clean, flexible way to construct objects step by step without having massive constructors or multiple constructor overloads.

## When to Use

- When creating an object requires many steps or configurations
- When an object can have different representations but the same construction process
- When you want to avoid telescoping constructors (constructors with many parameters)
- When the construction algorithm should be independent of the parts that make up the object
- When you need to create immutable objects with many optional fields
- When you want to provide a clear, readable way to construct complex objects

## Structure

The pattern consists of:

1. **Product** (`Computer`): The complex object being built
2. **Builder** (`ComputerBuilder`): Interface or class that specifies methods for creating parts of the Product
3. **Concrete Builder**: Implements the Builder interface and constructs the product step by step
4. **Director** (optional): Defines the order in which to call construction steps for common configurations
5. **Client**: Uses the Builder (directly or through Director) to construct objects

## Implementation Details

In this example, we implement a computer building system:

- **Computer**: The product with many components (CPU, RAM, GPU, etc.)
- **ComputerBuilder**: Provides methods to set each component, using method chaining (fluent interface)
- **Director**: Encapsulates common PC configurations (Gaming PC, Office PC, Workstation)
- **Method Chaining**: Each builder method returns the builder itself, enabling fluent syntax

The builder allows constructing computers in two ways:
1. **With Director**: Use predefined configurations for common use cases
2. **Without Director**: Manually chain methods for custom configurations

## Use Cases

1. **Database Query Builders**: Building complex SQL queries step by step
   ```go
   query := QueryBuilder().
       Select("name", "email").
       From("users").
       Where("age > ?", 18).
       OrderBy("name").
       Build()
   ```

2. **HTTP Request Builders**: Constructing HTTP requests with headers, body, authentication
3. **Document/Report Generators**: Building complex documents with various sections and formatting
4. **UI Component Builders**: Creating complex UI widgets with many optional properties
5. **Configuration Objects**: Building application configs with many optional settings
6. **Test Data Builders**: Creating test objects with various states for testing
7. **Email Builders**: Constructing emails with recipients, attachments, formatting
8. **Pizza/Meal Ordering**: Building custom food orders with various toppings and options

## Advantages

- **Readability**: Method chaining creates self-documenting, readable code
- **Flexibility**: Easy to add new optional parameters without breaking existing code
- **Immutability**: Can build immutable objects by only exposing the builder for construction
- **Step-by-Step Construction**: Complex objects can be built gradually
- **Reusability**: Same builder can create different representations
- **Single Responsibility**: Construction logic is separate from business logic
- **Director Encapsulation**: Common configurations can be encapsulated and reused

## Disadvantages

- **Complexity**: Increases the number of classes in the codebase
- **Overhead**: May be overkill for simple objects with few fields
- **Incomplete Objects**: Without validation, objects might be built in invalid states
- **Memory**: Creates an additional builder object for each product instance

## Method Chaining (Fluent Interface)

The implementation uses method chaining where each setter returns the builder itself:

```go
func (b *ComputerBuilder) SetCPU(cpu string) *ComputerBuilder {
    b.computer.CPU = cpu
    return b  // Returns the builder for chaining
}
```

This enables the fluent syntax:
```go
computer := NewComputerBuilder().
    SetCPU("Intel i9").
    SetRAM(32).
    SetGPU("RTX 4090").
    Build()
```

## Director Pattern (Optional)

The Director class encapsulates common build sequences:

- **Pros**: Reusable configurations, consistent object creation
- **Cons**: Additional class, less flexibility for unique configurations
- **When to Use**: When you have common configurations that are used frequently

You can use the builder with or without the Director depending on your needs.

## Running the Example

```bash
# Navigate to the builder directory
cd builder

# Run the example
go run main.go
```

## Expected Output

```
=== Builder Pattern Demo ===

--- Building Gaming PC (using Director) ---
=== Computer Specifications ===
CPU: Intel Core i9-13900K
RAM: 32 GB
Storage: 2TB NVMe SSD
GPU: NVIDIA RTX 4090
Motherboard: ASUS ROG Maximus Z790
Power Supply: 1000W 80+ Gold
Cooling: Liquid Cooling 360mm
Case: Full Tower RGB
Features: WiFi, Bluetooth, RGB Lighting

--- Building Office PC (using Director) ---
=== Computer Specifications ===
CPU: Intel Core i5-13400
RAM: 16 GB
Storage: 512GB SSD
GPU: Integrated Graphics
Motherboard: ASUS Prime B660
Power Supply: 450W 80+ Bronze
Cooling: Air Cooling
Case: Mid Tower
Features: WiFi

--- Building Workstation PC (using Director) ---
=== Computer Specifications ===
CPU: AMD Ryzen 9 7950X
RAM: 64 GB
Storage: 4TB NVMe SSD + 8TB HDD
GPU: NVIDIA RTX 4080
Motherboard: ASUS Pro WS X670E
Power Supply: 850W 80+ Platinum
Cooling: Liquid Cooling 280mm
Case: Mid Tower
Features: WiFi, Bluetooth

--- Building Custom PC (without Director) ---
=== Computer Specifications ===
CPU: AMD Ryzen 7 7800X3D
RAM: 32 GB
Storage: 1TB NVMe SSD
GPU: AMD Radeon RX 7900 XTX
Motherboard: MSI MAG X670E
Power Supply: 750W 80+ Gold
Cooling: Air Cooling - Noctua NH-D15
Case: Mid Tower
Features: WiFi, RGB Lighting
```

## Key Takeaways

- Use Builder when constructing complex objects with many optional parameters
- Method chaining (fluent interface) makes the code readable and maintainable
- Director is optional - use it for common configurations, skip it for flexibility
- The pattern is especially useful for creating immutable objects
- Avoids the telescoping constructor anti-pattern
- Separates construction logic from the object's representation
- Trade-off: Added complexity vs. improved code readability and maintainability

## Builder vs Other Patterns

| Pattern | Focus | Use Case |
|---------|-------|----------|
| **Builder** | Step-by-step construction | Complex objects with many optional parts |
| **Factory Method** | Single object creation | Simple objects with different implementations |
| **Abstract Factory** | Families of objects | Related objects that work together |
| **Prototype** | Cloning existing objects | Creating copies with modifications |

## Validation

In production code, you might want to add validation in the `Build()` method:

```go
func (b *ComputerBuilder) Build() (*Computer, error) {
    if b.computer.CPU == "" {
        return nil, errors.New("CPU is required")
    }
    if b.computer.RAM < 4 {
        return nil, errors.New("minimum 4GB RAM required")
    }
    return b.computer, nil
}
```

This ensures that only valid objects are constructed.
