# Composite Pattern

## Overview

The Composite Pattern is a structural design pattern that lets you compose objects into tree structures to represent part-whole hierarchies. It allows clients to treat individual objects and compositions of objects uniformly. The pattern creates a tree structure where individual objects (leaves) and groups of objects (composites) implement the same interface.

## Problem It Solves

When you have hierarchical data structures (trees) where individual elements and groups of elements need to be treated the same way, writing separate code for each type becomes cumbersome. The Composite pattern solves this by letting you treat both simple and complex elements through a common interface, eliminating the need for type checking.

## When to Use

- When you want to represent part-whole hierarchies of objects
- When you want clients to treat individual objects and compositions uniformly
- When you have a tree structure of objects
- When you want to ignore the difference between compositions of objects and individual objects
- When you need to apply operations to all elements in a hierarchical structure
- When the structure can be represented as a tree (one root, multiple branches and leaves)

## Structure

The pattern consists of:

1. **Component Interface**: Common interface for both leaf and composite objects
2. **Leaf**: Represents individual objects with no children (end nodes)
3. **Composite**: Contains child components (can be leaves or other composites)
4. **Client**: Works with components through the common interface

## Key Characteristics

### Uniformity
Both leaves and composites implement the same interface:
```go
type Component interface {
    Operation()
}
```

### Tree Structure
```
Composite (root)
├── Leaf
├── Composite
│   ├── Leaf
│   └── Leaf
└── Leaf
```

### Recursive Composition
Composites can contain other composites, forming a recursive tree structure.

## Implementation Details

The example demonstrates three composite scenarios:

### 1. File System Structure
- **Component**: FileSystemComponent interface
- **Leaf**: File (individual files)
- **Composite**: Directory (contains files and subdirectories)
- **Operation**: Print hierarchy, calculate total size

**Key Feature**: Directories can contain files and other directories recursively.

### 2. Company Organization Structure
- **Component**: Employee interface
- **Leaf**: Developer, Designer (individual employees)
- **Composite**: Manager (has subordinates)
- **Operation**: Print hierarchy, calculate total payroll

**Key Feature**: Managers can have other managers as subordinates, forming a management hierarchy.

### 3. Restaurant Menu System
- **Component**: MenuComponent interface
- **Leaf**: MenuItem (individual dishes)
- **Composite**: Menu (contains menu items and submenus)
- **Operation**: Print menu, calculate total price

**Key Feature**: Menus can contain other menus (e.g., Main Menu → Breakfast Menu → Items).

## Use Cases

1. **GUI Component Systems**: Building complex UI from simple widgets
   - Window contains Panels
   - Panels contain Buttons, TextFields, etc.
   - All components can be rendered, hidden, resized uniformly

2. **File Systems**: Representing files and directories
   - Files and directories both have names and sizes
   - Directories contain files and subdirectories
   - Operations like copy, delete work on both

3. **Graphics Systems**: Building complex shapes from primitives
   - Group contains Circles, Rectangles, other Groups
   - All can be drawn, moved, scaled uniformly

4. **Document Structures**: Representing document hierarchies
   - Document contains Sections
   - Sections contain Paragraphs, Images, Tables
   - All can be rendered, exported, printed

5. **Organization Charts**: Representing company hierarchies
   - Employees and managers both have names and salaries
   - Managers have subordinates
   - Calculate total payroll recursively

6. **Menu Systems**: Building nested menu structures
   - Menus contain MenuItems and SubMenus
   - All can be displayed, enabled/disabled
   - Calculate total price for combo meals

7. **Expression Trees**: Mathematical or logical expressions
   - Numbers and Variables are leaves
   - Operators are composites containing sub-expressions
   - Evaluate the entire tree recursively

## Advantages

- **Uniformity**: Treat individual objects and compositions identically
- **Simplicity**: Client code doesn't need to distinguish between types
- **Flexibility**: Easy to add new component types
- **Recursive Operations**: Operations naturally propagate through the tree
- **Open/Closed Principle**: Can add new components without changing existing code
- **Natural Tree Representation**: Intuitive for hierarchical data

## Disadvantages

- **Overly General**: Can make design too general and harder to restrict component types
- **Type Safety**: Harder to restrict what can be added to composites
- **Complexity**: Additional abstraction layer increases complexity
- **Leaf-Specific Operations**: Leaves can't implement composite operations (or must throw errors)
- **Performance**: Recursive operations can be expensive for deep trees

## Composite vs Other Patterns

| Pattern | Purpose | Structure |
|---------|---------|-----------|
| **Composite** | Part-whole hierarchies | Tree structure |
| **Decorator** | Add responsibilities dynamically | Linear chain/wrapping |
| **Chain of Responsibility** | Pass requests along chain | Linear chain |
| **Flyweight** | Share common state efficiently | Flat structure with shared objects |

**Key Distinction**: Composite creates tree structures for part-whole relationships, while Decorator wraps a single object (though decorators can be nested).

## Running the Example

```bash
# Navigate to the composite directory
cd structural-patterns/composite

# Run the example
go run main.go
```

## Expected Output

```
=== Composite Pattern Demo ===

============================================================
Example 1: File System Structure
============================================================

📁 root/ (32512 bytes total)
  📁 home/ (30720 bytes total)
    📄 resume.pdf (2048 bytes)
    📄 photo.jpg (4096 bytes)
    📁 documents/ (24576 bytes total)
      📄 report.docx (8192 bytes)
      📄 presentation.pptx (16384 bytes)
  📁 projects/ (1536 bytes total)
    📄 main.go (1024 bytes)
    📄 README.md (512 bytes)
  📄 config.txt (256 bytes)

Total size: 32512 bytes

============================================================
Example 2: Company Organization Structure
============================================================

👔 Alice Johnson - CEO ($200000.00, Team Budget: $895000.00)
  👔 Bob Smith - CTO ($150000.00, Team Budget: $390000.00)
    👨‍💻 Charlie Brown - Senior Developer ($100000.00)
    👨‍💻 Diana Prince - Developer ($80000.00)
    👨‍💻 Eve Adams - Junior Developer ($60000.00)
  👔 Frank Miller - CDO ($140000.00, Team Budget: $305000.00)
    🎨 Grace Lee - Senior Designer ($90000.00)
    🎨 Henry Wong - UI/UX Designer ($75000.00)

Total company payroll: $895000.00

============================================================
Example 3: Restaurant Menu System
============================================================

📋 Main Menu (Total: $67.40)
  📋 Breakfast Menu (Total: $22.97)
    - Pancakes: Fluffy pancakes with syrup ($8.99)
    - Omelette: Three-egg omelette ($10.99)
    - Coffee: Fresh brewed coffee ($2.99)
  📋 Lunch Menu (Total: $36.97)
    - Burger: Beef burger with fries ($12.99)
    - Salad: Caesar salad ($9.99)
    - Pasta: Spaghetti carbonara ($13.99)
  📋 Drinks Menu (Total: $7.46)
    - Soda: Soft drink ($2.49)
    - Juice: Fresh orange juice ($3.99)
    - Water: Bottled water ($1.99)

--- Summary ---
Composite pattern treats individual objects and compositions uniformly
Tree structures can be built and traversed easily
Operations apply to both leaves and composites
```

## Key Takeaways

- **Uniform Treatment**: Leaves and composites share the same interface
- **Tree Structures**: Natural fit for hierarchical data
- **Recursive Operations**: Operations automatically propagate through the tree
- **Transparency vs Safety**: Trade-off between treating all components uniformly vs type safety
- **Client Simplification**: Client doesn't need to know if dealing with leaf or composite
- **Add/Remove Operations**: Composites need methods to manage children
- **Recursive Composition**: Composites can contain other composites indefinitely

## Implementation Variations

### Transparent Composite
All components have add/remove methods (may not be applicable to leaves):
```go
type Component interface {
    Operation()
    Add(Component)      // Leaves throw error
    Remove(Component)   // Leaves throw error
}
```
**Pros**: Uniform interface
**Cons**: Leaves must implement meaningless methods

### Safe Composite
Only composites have add/remove methods:
```go
type Component interface {
    Operation()
}

type Composite struct {
    children []Component
}

func (c *Composite) Add(component Component) { }
```
**Pros**: Type safe
**Cons**: Need type checking to use composite-specific methods

This implementation uses the **Safe Composite** approach.

## Design Considerations

### Child Management
- **Where to store parent reference?** Usually in the child (for traversing up)
- **Who can delete components?** Usually only the composite that contains them
- **Order of children?** May need to maintain insertion order

### Operation Execution
- **Pre-order**: Process parent before children
- **Post-order**: Process children before parent
- **Level-order**: Process level by level

### Caching
Composites can cache results of expensive operations:
```go
func (c *Composite) GetSize() int {
    if c.cachedSize != 0 {
        return c.cachedSize
    }
    // Calculate and cache
}
```

## Common Mistakes

1. **Not Using Common Interface**: Forgetting that leaves and composites must share an interface
2. **Circular References**: Accidentally creating cycles in the tree
3. **Inefficient Recursion**: Not caching results of expensive recursive operations
4. **Ignoring Memory**: Deep trees with many nodes can consume significant memory
5. **No Null Checks**: Not checking for nil children
6. **Wrong Pattern Choice**: Using Composite when structure isn't truly hierarchical

## Best Practices

1. **Define Clear Interface**: Common operations should make sense for both leaves and composites
2. **Document Hierarchy**: Clearly document the tree structure
3. **Consider Caching**: Cache expensive recursive calculations
4. **Provide Iterators**: Consider providing iterators for tree traversal
5. **Handle Edge Cases**: Check for cycles, null references, empty composites
6. **Use Safe or Transparent Consistently**: Choose one approach and stick with it
7. **Test Thoroughly**: Test deep hierarchies and edge cases

## Real-World Examples

- **Java AWT/Swing**: Container and Component hierarchy
- **HTML DOM**: Elements contain other elements
- **Company Organizational Charts**: Employees and departments
- **File Systems**: Unix/Linux file system structure
- **Graphics Editors**: Grouped shapes that can be grouped again
- **AST (Abstract Syntax Trees)**: In compilers and interpreters

## When NOT to Use Composite

- When your structure is not hierarchical
- When leaves and composites have fundamentally different operations
- When you need strict type safety and can't afford the generality
- When the tree structure is very shallow (just one level)
- When operations on composites are significantly different from leaves
