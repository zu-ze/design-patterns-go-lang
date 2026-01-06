# Prototype Pattern

## Overview

The Prototype Pattern is a creational design pattern that lets you copy existing objects without making your code dependent on their classes. Instead of creating new objects from scratch, you clone existing instances (prototypes). This pattern delegates the cloning process to the actual objects being cloned.

## Problem It Solves

When creating new objects is expensive (requires complex initialization, database queries, or network calls) or when you want to create objects without knowing their exact classes, the Prototype pattern provides a way to copy existing objects. It's especially useful when the object creation process is complex or when you need many variations of similar objects.

## When to Use

- When object creation is expensive or complex
- When you want to avoid subclasses of object creators (like Factory patterns)
- When you need to create objects at runtime whose types are determined dynamically
- When you want to keep the number of classes in a system to a minimum
- When creating an instance of a class is more convenient by copying an existing instance
- When you need to create objects that differ only slightly from existing objects

## Structure

### Core Components (Required)

The Prototype pattern at its essence requires only:

1. **Prototype Interface** (`Cloneable`, `Document`): Declares the cloning method
2. **Concrete Prototype** (`Resume`, `Report`): Implements the cloning method to copy itself
3. **Client**: Creates new objects by asking prototypes to clone themselves

### Optional Enhancement

4. **Prototype Registry/Manager** (`DocumentRegistry`): An optional component that maintains a registry of available prototypes for easy access. **Not part of the core pattern**, but useful when managing multiple prototype templates.

## Implementation Details

In this example, we implement a document cloning system:

- **Document Interface**: Defines the `Clone()` method that all documents must implement (core pattern)
- **Concrete Documents**: Resume and Report implement deep cloning (core pattern)
- **Nested Objects**: Author and Address demonstrate deep copying of nested structures
- **DocumentRegistry**: An optional enhancement that acts as a prototype manager, storing template documents that can be cloned
- **Deep vs Shallow Copy**: The implementation shows proper deep copying of slices, maps, and nested objects

The example demonstrates both approaches:
1. **Direct cloning** (core pattern): `clonedResume := originalResume.Clone()`
2. **Registry-based cloning** (optional enhancement): `newResume := registry.Create("resume-template")`

### Deep Copy vs Shallow Copy

**Shallow Copy**: Copies the object's fields, but references to nested objects are shared between original and clone.

**Deep Copy**: Copies the object and all nested objects recursively, creating completely independent copies.

This implementation uses **deep copying** to ensure cloned objects are completely independent:

```go
func (r *Resume) Clone() Document {
    // Deep copy slices
    experienceCopy := make([]string, len(r.Experience))
    copy(experienceCopy, r.Experience)
    
    // Deep copy nested objects
    Author: r.Author.Clone(),
    
    // Increment version for the clone
    Version: r.Version + 1,
}
```

## Use Cases

1. **Game Development**: Cloning enemy characters, items, or game states
   - Instead of creating enemies from scratch, clone a prototype with preset attributes

2. **Document Templates**: Cloning document templates (resumes, reports, invoices)
   - Users start with a template and customize it

3. **Configuration Objects**: Cloning complex configuration objects
   - Start with a base configuration and create variants

4. **UI Widgets**: Cloning UI components with preset styles and behaviors
   - Clone a button prototype and customize text/colors

5. **Database Records**: Cloning database entities with complex relationships
   - "Duplicate" functionality in applications

6. **Undo/Redo Systems**: Saving object states for undo functionality
   - Clone object state before each operation

7. **Testing**: Creating test data by cloning and modifying prototypes
   - Maintain test fixtures as prototypes

8. **Caching**: Caching expensive objects and returning clones
   - Avoid repeated expensive initialization

## Advantages

- **Performance**: Faster than creating objects from scratch when initialization is expensive
- **Reduced Subclassing**: Avoids creating factory class hierarchies
- **Dynamic Configuration**: Add and remove prototypes at runtime
- **Simpler Object Creation**: Creating complex objects is simplified to cloning
- **Encapsulation**: Clone method can access private fields
- **Flexibility**: Easy to create variations of objects

## Disadvantages

- **Complexity**: Implementing deep copy can be complex, especially with circular references
- **Clone Method Implementation**: Each class needs to implement its own cloning logic
- **Nested Objects**: Deep copying nested objects requires careful implementation
- **Circular References**: Handling circular references in object graphs is tricky
- **Immutable Fields**: Cloning objects with immutable fields or const members can be challenging

## Shallow Copy vs Deep Copy Considerations

**Shallow Copy Problems:**
```go
// BAD: Shallow copy shares references
func (r *Resume) Clone() *Resume {
    return &Resume{
        Author: r.Author, // Both point to same Author!
        Skills: r.Skills, // Both point to same slice!
    }
}
// Modifying clone.Skills affects original.Skills
```

**Deep Copy Solution:**
```go
// GOOD: Deep copy creates independent objects
func (r *Resume) Clone() *Resume {
    skillsCopy := make([]string, len(r.Skills))
    copy(skillsCopy, r.Skills)
    return &Resume{
        Author: r.Author.Clone(),
        Skills: skillsCopy,
    }
}
```

## Prototype Registry (Optional Enhancement)

**Important**: The registry is **NOT part of the core Prototype pattern**. The pattern works perfectly fine with just direct cloning: `clone := original.Clone()`.

However, a registry can be added as an optional enhancement when you have multiple prototypes to manage:

```go
registry := NewDocumentRegistry()
registry.Register("resume-template", resumePrototype)
registry.Register("report-template", reportPrototype)

// Later, create instances by cloning
newResume := registry.Create("resume-template")
```

**When to Use a Registry:**
- You have multiple prototype templates to manage
- You want centralized access to prototypes by name/key
- Prototypes need to be added/removed dynamically at runtime
- You want to simplify client code that uses many different prototypes

**When NOT to Use a Registry:**
- You have only one or two prototypes
- Prototypes are created inline or locally
- Direct cloning is simpler for your use case

**Benefits:**
- Centralized prototype management
- Easy lookup by name or key
- Can add/remove prototypes at runtime
- Simplifies client code when managing many templates

## Running the Example

```bash
# Navigate to the prototype directory
cd prototype

# Run the example
go run main.go
```

## Expected Output

```
=== Prototype Pattern Demo ===

--- Creating Original Resume ---
=== RESUME ===
Title: Software Engineer Resume
Author: John Doe (john.doe@example.com)
Location: San Francisco, USA
Education: BS Computer Science
Skills: Go, Python, Docker, Kubernetes
Experience: 2 positions
Version: 1
Created: 2024-XX-XX
Modified: 2024-XX-XX

--- Cloning Resume and Modifying ---
=== RESUME ===
Title: Senior Software Engineer Resume
Author: Jane Smith (jane.smith@example.com)
Location: San Francisco, USA
Education: BS Computer Science
Skills: Go, Python, Docker, Kubernetes, Rust, GraphQL
Experience: 2 positions
Version: 2
Created: 2024-XX-XX
Modified: 2024-XX-XX

--- Original Resume (unchanged) ---
(Shows original is unchanged, demonstrating deep copy)

--- Using Document Registry (Prototype Manager) ---
(Shows creating new documents from templates)
```

## Key Takeaways

- **Core Pattern**: Objects that can clone themselves via a `Clone()` method
- **Registry is Optional**: Not part of the core pattern, only add it if you need to manage multiple prototypes
- Use Prototype when object creation is expensive or complex
- Always implement deep copying for objects with nested structures
- Be careful with circular references and mutable fields
- Cloning is faster than reconstruction when initialization is costly
- Perfect for creating variations of similar objects
- Trade-off: Clone method complexity vs. simplified object creation

## Prototype vs Other Patterns

| Pattern | Focus | Use Case |
|---------|-------|----------|
| **Prototype** | Cloning existing objects | Expensive object creation, many similar objects |
| **Factory Method** | Creating new objects | Different implementations of same interface |
| **Abstract Factory** | Creating families | Related objects that work together |
| **Builder** | Step-by-step construction | Complex objects with many configurations |
| **Singleton** | Single instance | Exactly one instance needed |

## Go-Specific Considerations

In Go, you typically implement cloning by:

1. **Manual Copy Method**: Explicitly copy each field (shown in this example)
2. **Copy Constructor**: Function that takes an existing object and returns a copy
3. **Serialization**: Marshal/unmarshal for deep copying (less efficient)

Go doesn't have built-in clone functionality like some languages, so careful manual implementation is necessary.

## Common Pitfalls

1. **Forgetting to clone nested objects** - Results in shared references
2. **Not handling nil pointers** - Can cause panics
3. **Circular references** - Can cause infinite loops
4. **Not incrementing version numbers** - Loses track of clones
5. **Modifying slices/maps without copying** - Affects original object

## Best Practices

- Always perform deep copying for reference types (slices, maps, pointers)
- Consider implementing a `DeepCopy()` method separate from `Clone()` if both are needed
- Document whether your clone is deep or shallow
- Test clones to ensure they're truly independent
- Consider using a registry for frequently used prototypes
- Increment version numbers or timestamps to track clones
