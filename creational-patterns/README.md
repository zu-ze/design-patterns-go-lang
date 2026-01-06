# Creational Design Patterns

Creational design patterns deal with object creation mechanisms, trying to create objects in a manner suitable to the situation. These patterns provide flexibility in what gets created, who creates it, how it gets created, and when.

## Patterns in This Category

### 1. Factory Method Pattern
**Purpose**: Provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created.

**Use When**: You need flexibility in object creation without knowing the exact class beforehand.

**Example**: Creating different notification types (Email, SMS, Push).

[View Implementation](./factory-method/)

---

### 2. Abstract Factory Pattern
**Purpose**: Provides an interface for creating families of related or dependent objects without specifying their concrete classes.

**Use When**: You need to ensure that products from the same family are used together.

**Example**: Creating UI components for different themes (Light, Dark).

[View Implementation](./abstract-factory/)

---

### 3. Builder Pattern
**Purpose**: Lets you construct complex objects step by step, separating construction from representation.

**Use When**: Creating objects requires many steps or has many optional parameters.

**Example**: Building computers with various components and configurations.

[View Implementation](./builder/)

---

### 4. Prototype Pattern
**Purpose**: Lets you copy existing objects without making your code dependent on their classes.

**Use When**: Object creation is expensive, or you need variations of similar objects.

**Example**: Cloning document templates (resumes, reports).

[View Implementation](./prototype/)

---

### 5. Singleton Pattern
**Purpose**: Ensures a class has only one instance and provides a global point of access to it.

**Use When**: Exactly one instance is needed throughout the application.

**Example**: Database connection, logger, configuration manager.

[View Implementation](./singleton/)

---

## Quick Comparison

| Pattern | Focus | Key Benefit | Common Use Case |
|---------|-------|-------------|-----------------|
| **Factory Method** | Single product creation | Flexibility in object types | Notification systems |
| **Abstract Factory** | Family of products | Consistency across related objects | UI themes, cross-platform |
| **Builder** | Step-by-step construction | Complex object assembly | Configuration objects |
| **Prototype** | Clone existing objects | Avoid expensive creation | Document templates |
| **Singleton** | Single instance | Controlled access to shared resource | Database connection |

## When to Use Creational Patterns

Use creational patterns when:
- Object creation is complex or resource-intensive
- You want to decouple object creation from usage
- You need flexibility in what objects get created
- You want to control how many instances exist
- Object initialization requires many parameters

## Running Examples

Each pattern has its own folder with:
- `main.go` - Implementation example
- `README.md` - Detailed documentation
- `go.mod` - Go module file

To run any pattern:
```bash
cd <pattern-name>
go run main.go
```

For example:
```bash
cd factory-method
go run main.go
```
