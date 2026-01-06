# Flyweight Pattern

## Overview

The Flyweight Pattern is a structural design pattern that minimizes memory usage by sharing as much data as possible with similar objects. It's used to support large numbers of fine-grained objects efficiently by sharing common parts of object state between multiple objects instead of keeping all data in each object.

## Problem It Solves

When you need to create a large number of similar objects, memory consumption can become prohibitive. For example, rendering 10,000 trees in a game where each tree stores its texture, color, and other visual data would consume massive amounts of memory. The Flyweight pattern solves this by extracting the shared state (intrinsic state) and storing it in flyweight objects that can be shared, while keeping unique state (extrinsic state) separate.

## When to Use

- When your application uses a large number of similar objects
- When storage costs are high because of the quantity of objects
- When most object state can be made extrinsic (extracted out)
- When many objects can be replaced by few shared objects
- When the application doesn't depend on object identity (using == to compare objects)
- When you need to optimize memory usage in resource-constrained environments

## Core Concepts

### Intrinsic State (Shared)
State that is shared among many objects and doesn't change:
- Tree type, color, texture
- Character font, size, style
- Particle sprite, animation

### Extrinsic State (Unique)
State that varies between objects and is stored externally:
- Tree position (x, y)
- Character position in document
- Particle position, velocity

### Factory
Manages flyweight instances and ensures sharing:
```go
type Factory struct {
    flyweights map[string]*Flyweight
}

func (f *Factory) GetFlyweight(key string) *Flyweight {
    if flyweight, exists := f.flyweights[key]; exists {
        return flyweight // Reuse existing
    }
    flyweight := NewFlyweight()
    f.flyweights[key] = flyweight
    return flyweight
}
```

## Structure

The pattern consists of:

1. **Flyweight Interface**: Declares methods that accept extrinsic state
2. **Concrete Flyweight**: Implements the interface and stores intrinsic state (shared)
3. **Flyweight Factory**: Creates and manages flyweight objects, ensures sharing
4. **Context**: Contains extrinsic state and references to flyweight objects
5. **Client**: Maintains references to flyweights and computes/stores extrinsic state

## Implementation Details

The example demonstrates three flyweight scenarios:

### 1. Forest Rendering System
- **Intrinsic State**: Tree type, color, texture (TreeType)
- **Extrinsic State**: Tree position x, y (Tree)
- **Factory**: TreeTypeFactory manages shared tree types
- **Memory Savings**: 8 trees with 3 unique types instead of 8 full tree objects

**Key Insight**: Instead of storing full tree data 8 times, store 3 shared types and 8 position references.

### 2. Text Editor
- **Intrinsic State**: Font, size, color, bold, italic (CharacterStyle)
- **Extrinsic State**: Character value, position (Character)
- **Factory**: StyleFactory manages shared styles
- **Memory Savings**: Thousands of characters share a few dozen styles

**Key Insight**: A document with 10,000 characters might only use 20 unique styles.

### 3. Particle System
- **Intrinsic State**: Sprite graphics, dimensions (ParticleSprite)
- **Extrinsic State**: Particle position, velocity (Particle)
- **Factory**: SpriteFactory manages shared sprites
- **Memory Savings**: 1,000 particles share 3 sprite types

**Key Insight**: 1,000 particles reusing 3 sprites = 333x average reuse per sprite.

## Use Cases

1. **Game Development**: Managing thousands of game objects
   - Trees, rocks, grass in terrain
   - Bullets, particles in effects systems
   - NPCs with shared models and textures
   - Tile-based maps

2. **Text Rendering**: Displaying documents with styled text
   - Character glyphs in fonts
   - Text formatting (bold, italic, colors)
   - Document editors with thousands of characters

3. **Graphics Systems**: Rendering complex scenes
   - Reusing textures and materials
   - Sprite sheets in 2D games
   - Instanced rendering in 3D graphics

4. **UI Frameworks**: Managing many UI elements
   - Icons and images
   - Common UI components (buttons, checkboxes)
   - Theme elements (colors, fonts)

5. **String Interning**: Optimizing string storage
   - Java String pool
   - Symbol tables in compilers
   - Database connection strings

6. **Cache Systems**: Sharing cached data
   - HTTP response caching
   - Database query result caching
   - Computed value caching

## Advantages

- **Memory Reduction**: Dramatically reduces memory usage for large numbers of objects
- **Performance**: Can improve performance by reducing memory allocations
- **Scalability**: Enables applications to handle more objects
- **Efficient Sharing**: Automatically shares common state
- **Centralized Management**: Factory provides single point for managing shared objects

## Disadvantages

- **Complexity**: Adds complexity by separating intrinsic and extrinsic state
- **Code Complexity**: More difficult to understand and maintain
- **Runtime Cost**: Computation overhead to calculate/pass extrinsic state
- **Thread Safety**: Factory must be thread-safe in concurrent environments
- **Design Constraints**: Not applicable when objects don't share state
- **Debugging**: Harder to debug shared state issues

## Flyweight vs Other Patterns

| Pattern | Purpose | Key Difference |
|---------|---------|----------------|
| **Flyweight** | Share state to reduce memory | Shares intrinsic state, extrinsic stored separately |
| **Singleton** | Single instance | One instance total, not about sharing state |
| **Prototype** | Clone objects | Creates copies, doesn't share state |
| **Object Pool** | Reuse expensive objects | Reuses instances, not about shared state |

**Key Distinction**: Flyweight is about **sharing state**, not sharing instances (though it does both).

## Memory Savings Calculation

### Without Flyweight
```
Total Memory = Number of Objects × Size of Each Object
Example: 10,000 trees × 100 bytes = 1,000,000 bytes
```

### With Flyweight
```
Total Memory = (Number of Objects × Size of Extrinsic State) + 
               (Number of Flyweights × Size of Intrinsic State)
Example: (10,000 trees × 16 bytes) + (10 types × 100 bytes) = 161,000 bytes
Savings: 84% reduction
```

## Running the Example

```bash
# Navigate to the flyweight directory
cd structural-patterns/flyweight

# Run the example
go run main.go
```

## Expected Output

```
=== Flyweight Pattern Demo ===

--- Example 1: Forest Rendering ---

Planting trees in a forest...
[TreeType] Creating new tree type: Oak (Green, Rough)
[TreeTypeFactory] Reusing existing tree type: Oak-Green-Rough
[TreeType] Creating new tree type: Pine (Dark Green, Smooth)
[TreeTypeFactory] Reusing existing tree type: Pine-Dark Green-Smooth
[TreeType] Creating new tree type: Birch (White, Smooth)
[TreeTypeFactory] Reusing existing tree type: Oak-Green-Rough
...

[Forest] Drawing all trees:
[TreeType] Drawing Oak tree at (10, 20) - Color: Green, Texture: Rough
[TreeType] Drawing Pine tree at (30, 40) - Color: Dark Green, Texture: Smooth
...

[Forest] Memory Usage:
Trees: 8, Types: 3
Memory without Flyweight: ~800 bytes
Memory with Flyweight: ~428 bytes
Memory saved: ~372 bytes (46.5% reduction)

--- Example 2: Text Editor ---

Creating a document with styled text...
[CharacterStyle] Creating style: Arial, 12pt, Black
[StyleFactory] Reusing style: Arial-12-Black-true-false
...

[TextEditor] Rendering document:
[Arial, 12pt, Black]: Hello 
[Arial, 12pt, Black, Bold]: World
...

[TextEditor] Statistics:
Characters: 73, Unique Styles: 5
Memory without Flyweight: ~3650 bytes
Memory with Flyweight: ~930 bytes
Memory saved: ~2720 bytes (74.5% reduction)

--- Example 3: Particle System ---

Creating particle effects...
[ParticleSprite] Loading sprite: spark (orange, 16x16)
[ParticleSprite] Loading sprite: smoke (gray, 16x16)
[ParticleSprite] Loading sprite: spark (yellow, 16x16)

[ParticleSystem] Statistics:
Particles: 1000, Unique Sprites: 3
Average sprite reuse: 333.3x

--- Summary ---
Flyweight pattern shares common state (intrinsic) among many objects
Each object stores only unique state (extrinsic)
Dramatically reduces memory usage when many similar objects exist
Factory manages flyweight instances and ensures sharing
```

## Key Takeaways

- **State Separation**: Separate intrinsic (shared) from extrinsic (unique) state
- **Factory Pattern**: Use factory to manage and share flyweight instances
- **Memory Optimization**: Primary benefit is dramatic memory reduction
- **Many Similar Objects**: Most effective with large numbers of similar objects
- **Trade-offs**: Adds complexity but saves memory
- **Immutability**: Flyweights should be immutable to be safely shared
- **Context Objects**: Store extrinsic state in lightweight context objects

## Design Considerations

### Identifying Intrinsic vs Extrinsic State

**Intrinsic State (Should be in Flyweight)**:
- Shared by many objects
- Doesn't change (immutable)
- Independent of flyweight's context
- Examples: type, color, texture, font, sprite

**Extrinsic State (Should be outside Flyweight)**:
- Unique to each object
- May change over time
- Depends on flyweight's context
- Examples: position, velocity, size, selection state

### Thread Safety

Flyweight factories must be thread-safe:

```go
type Factory struct {
    mu         sync.RWMutex
    flyweights map[string]*Flyweight
}

func (f *Factory) GetFlyweight(key string) *Flyweight {
    f.mu.RLock()
    if fw, exists := f.flyweights[key]; exists {
        f.mu.RUnlock()
        return fw
    }
    f.mu.RUnlock()
    
    f.mu.Lock()
    defer f.mu.Unlock()
    
    // Double-check after acquiring write lock
    if fw, exists := f.flyweights[key]; exists {
        return fw
    }
    
    fw := NewFlyweight()
    f.flyweights[key] = fw
    return fw
}
```

### Immutability

Flyweights should be immutable to be safely shared:

```go
type TreeType struct {
    name    string // Immutable
    color   string // Immutable
    texture string // Immutable
}

// No setters - state cannot be modified after creation
```

## Common Mistakes

1. **Mutable Flyweights**: Making flyweights mutable breaks sharing safety
2. **Too Much Extrinsic State**: If too much state is extrinsic, memory savings diminish
3. **Premature Optimization**: Using flyweight when you don't have many objects
4. **Not Using Factory**: Creating flyweights directly instead of through factory
5. **Wrong State Classification**: Misidentifying intrinsic vs extrinsic state
6. **Ignoring Thread Safety**: Not making factory thread-safe in concurrent code

## Best Practices

1. **Use Factory**: Always create flyweights through a factory
2. **Make Immutable**: Flyweights should be immutable
3. **Profile First**: Measure memory usage before applying pattern
4. **Clear Keys**: Use clear, consistent keys for flyweight lookup
5. **Document State**: Clearly document what's intrinsic vs extrinsic
6. **Consider Weak References**: In some languages, use weak references to allow garbage collection
7. **Cache Strategically**: Consider cache eviction for rarely-used flyweights

## Flyweight vs Object Pool

| Aspect | Flyweight | Object Pool |
|--------|-----------|-------------|
| **Purpose** | Share state | Reuse expensive instances |
| **State** | Intrinsic state shared | Full objects reused |
| **Lifetime** | Lives as long as needed | Returned to pool |
| **Identity** | Many contexts use same flyweight | One client uses object at a time |
| **Memory Goal** | Reduce memory by sharing | Reduce allocation cost |

## Real-World Examples

- **Java String Interning**: String pool shares identical string instances
- **Swing Font Rendering**: Shares font glyph data
- **Game Engines**: Unity/Unreal use instancing (flyweight concept)
- **Web Browsers**: Share CSS styles, fonts, images across DOM elements
- **Database Connection Pools**: Share connection configuration
- **Compiler Symbol Tables**: Share symbol information

## Performance Considerations

### When Flyweight Helps
- Large numbers of objects (thousands to millions)
- Objects have significant shared state
- Object creation is expensive
- Memory is constrained

### When Flyweight May Not Help
- Few objects (less than hundreds)
- Little shared state
- Extrinsic state is expensive to compute/pass
- Complexity cost outweighs memory savings

## Testing Flyweight Pattern

Test both functionality and memory savings:

```go
func TestTreeTypeSharing(t *testing.T) {
    factory := NewTreeTypeFactory()
    
    type1 := factory.GetTreeType("Oak", "Green", "Rough")
    type2 := factory.GetTreeType("Oak", "Green", "Rough")
    
    // Same flyweight instance should be returned
    if type1 != type2 {
        t.Error("Expected same flyweight instance")
    }
    
    if factory.GetTotalTypes() != 1 {
        t.Error("Expected only 1 tree type")
    }
}
```

## When NOT to Use Flyweight

- When you have few objects
- When objects don't share much state
- When the complexity cost outweighs memory savings
- When object identity is important (can't share instances)
- When state is mostly unique to each object
- When premature optimization would hurt maintainability
