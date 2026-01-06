# Proxy Pattern

## Overview

The Proxy Pattern is a structural design pattern that provides a surrogate or placeholder for another object to control access to it. A proxy acts as an intermediary, implementing the same interface as the real object, allowing it to intercept, control, or augment operations on the real object.

## Problem It Solves

Sometimes you need to add functionality around object access without modifying the object itself. This could be for lazy initialization, access control, logging, caching, or remote object access. The Proxy pattern solves this by creating a wrapper that looks like the real object but adds extra behavior.

## When to Use

- When you want to control access to an object
- When you want to add lazy initialization (create expensive objects only when needed)
- When you need access control or authentication
- When you want to add caching to avoid redundant operations
- When you need to represent a remote object locally
- When you want to add logging, monitoring, or reference counting
- When you want to add cleanup operations before/after method calls

## Types of Proxies

### 1. Virtual Proxy (Lazy Loading)
Delays expensive object creation until actually needed.

**Use Case**: Loading large images, database connections, heavy resources

### 2. Protection Proxy (Access Control)
Controls access based on permissions or credentials.

**Use Case**: User authentication, role-based access control

### 3. Caching Proxy
Caches results to avoid redundant expensive operations.

**Use Case**: API response caching, query result caching

### 4. Smart Proxy (Smart Reference)
Adds additional logic like reference counting, locking, authentication.

**Use Case**: Authentication layers, resource locking, counting references

### 5. Remote Proxy
Represents object in different address space (network, process).

**Use Case**: Network communication, RPC, load balancing, distributed systems

### 6. Logging Proxy
Adds logging before/after method calls.

**Use Case**: Audit trails, debugging, monitoring

## Structure

The pattern consists of:

1. **Subject Interface**: Common interface for both proxy and real object
2. **Real Subject**: The actual object that does the real work
3. **Proxy**: Implements the same interface, holds reference to real subject, controls access
4. **Client**: Works with objects through the Subject interface

## Implementation Details

The example demonstrates five types of proxies:

### 1. Virtual Proxy - Image Loading
- **Real Subject**: RealImage (loads from disk - expensive)
- **Proxy**: ImageProxy (delays loading until Display() called)
- **Benefit**: Images only loaded when actually displayed, not when object created

**Key Feature**: Lazy initialization - real object created on first access.

### 2. Protection Proxy - Database Access Control
- **Real Subject**: RealDatabase (executes queries)
- **Proxy**: DatabaseProxy (checks permissions before allowing queries)
- **Benefit**: Enforces access control without modifying database code

**Key Feature**: Role-based access control - admins can do anything, users only SELECT.

### 3. Caching Proxy - YouTube Downloader
- **Real Subject**: YouTubeDownloader (downloads videos - slow)
- **Proxy**: CachedYouTubeProxy (caches downloaded videos)
- **Benefit**: Repeated requests served from cache instantly

**Key Feature**: Cache management - stores results to avoid redundant downloads.

### 4. Smart Proxy - Bank Account Authentication
- **Real Subject**: RealBankAccount (performs transactions)
- **Proxy**: BankAccountProxy (adds PIN authentication and attempt limiting)
- **Benefit**: Adds security layer without modifying account logic

**Key Feature**: Additional logic - authentication, attempt counting, account locking.

### 5. Remote Proxy - Load Balancer
- **Real Subject**: RealServer instances (process requests)
- **Proxy**: LoadBalancerProxy (distributes requests across servers)
- **Benefit**: Transparent load distribution

**Key Feature**: Request routing - round-robin distribution to multiple servers.

## Use Cases

1. **Lazy Loading**: Defer expensive operations
   - Image loading in galleries
   - Document loading in editors
   - Large object initialization

2. **Access Control**: Security and permissions
   - API authentication
   - Database access control
   - File system permissions
   - Admin vs user operations

3. **Caching**: Improve performance
   - HTTP response caching
   - Database query caching
   - API call caching
   - Computed result caching

4. **Remote Objects**: Network communication
   - RPC (Remote Procedure Call)
   - Distributed objects
   - Microservice communication
   - Network load balancing

5. **Logging and Monitoring**: Track operations
   - Audit trails
   - Performance monitoring
   - Debug logging
   - Usage analytics

6. **Smart References**: Additional management
   - Reference counting
   - Copy-on-write
   - Thread safety
   - Resource cleanup

## Advantages

- **Control Access**: Full control over how object is accessed
- **Lazy Initialization**: Create expensive objects only when needed
- **Open/Closed Principle**: Add functionality without modifying real object
- **Security**: Add authentication and authorization layers
- **Performance**: Caching can dramatically improve performance
- **Transparency**: Client code doesn't know it's using a proxy
- **Separation of Concerns**: Proxy handles cross-cutting concerns

## Disadvantages

- **Complexity**: Adds extra layer and classes
- **Performance Overhead**: Extra method call in proxy (minimal but exists)
- **Response Time**: May introduce latency for first access (lazy loading)
- **Code Duplication**: Proxy must implement same interface as real object
- **Maintenance**: Proxy must be updated when real object's interface changes

## Proxy vs Other Patterns

| Pattern | Purpose | Key Difference |
|---------|---------|----------------|
| **Proxy** | Control access to object | Same interface, controls access/lifecycle |
| **Decorator** | Add responsibilities | Same interface, adds functionality |
| **Adapter** | Make interfaces compatible | Changes interface |
| **Facade** | Simplify complex system | Provides simpler interface |

**Key Distinctions**:
- **Proxy controls access** (when/how to access)
- **Decorator adds behavior** (what object does)
- **Adapter converts interface** (how you interact)
- **Facade simplifies interface** (easier way to interact)

## Proxy vs Decorator

| Aspect | Proxy | Decorator |
|--------|-------|-----------|
| **Purpose** | Control access | Enhance functionality |
| **Creation** | Usually creates/manages real object | Wraps existing object |
| **Lifecycle** | Manages real object lifecycle | Doesn't manage lifecycle |
| **Intent** | Access control, lazy loading | Add responsibilities |
| **Example** | Authentication proxy | Logging decorator |

**In Practice**: The line can blur - a logging proxy vs logging decorator is often a matter of perspective.

## Running the Example

```bash
# Navigate to the proxy directory
cd structural-patterns/proxy

# Run the example
go run main.go
```

## Expected Output

```
=== Proxy Pattern Demo ===

--- Example 1: Virtual Proxy (Lazy Loading) ---

Creating image proxies (images not loaded yet)...
[ImageProxy] Created proxy for: photo1.jpg (not loaded yet)
[ImageProxy] Created proxy for: photo2.jpg (not loaded yet)

Displaying first image (triggers loading)...
[ImageProxy] First access - loading real image...
[RealImage] Loading image from disk: photo1.jpg (expensive operation)
[RealImage] Image loaded: photo1.jpg
[RealImage] Displaying image: photo1.jpg

Displaying first image again (already loaded)...
[RealImage] Displaying image: photo1.jpg

Displaying second image (triggers loading)...
[ImageProxy] First access - loading real image...
[RealImage] Loading image from disk: photo2.jpg (expensive operation)
[RealImage] Image loaded: photo2.jpg
[RealImage] Displaying image: photo2.jpg

--- Example 2: Protection Proxy (Access Control) ---

[RealDatabase] Connecting to database: localhost:5432/mydb

[DatabaseProxy] User role: admin attempting query
[DatabaseProxy] Access granted - forwarding to real database
[RealDatabase] Executing query: SELECT * FROM users
[DatabaseProxy] Logging query: SELECT * FROM users by user role: admin

[DatabaseProxy] User role: admin attempting query
[DatabaseProxy] Access granted - forwarding to real database
[RealDatabase] Executing query: DELETE FROM logs WHERE date < '2024-01-01'
[DatabaseProxy] Logging query: DELETE FROM logs WHERE date < '2024-01-01' by user role: admin

[RealDatabase] Connecting to database: localhost:5432/mydb

[DatabaseProxy] User role: user attempting query
[DatabaseProxy] Access granted - forwarding to real database
[RealDatabase] Executing query: SELECT * FROM products
[DatabaseProxy] Logging query: SELECT * FROM products by user role: user

[DatabaseProxy] User role: user attempting query
[DatabaseProxy] ACCESS DENIED - Insufficient permissions

--- Example 3: Caching Proxy ---

First request for video:

[CachedYouTubeProxy] Request for video: dQw4w9WgXcQ
[CachedYouTubeProxy] Cache MISS - downloading from YouTube
[YouTubeDownloader] Downloading video dQw4w9WgXcQ from YouTube (slow operation)...
[CachedYouTubeProxy] Cached video: dQw4w9WgXcQ

Second request for same video (cached):

[CachedYouTubeProxy] Request for video: dQw4w9WgXcQ
[CachedYouTubeProxy] Cache HIT - returning cached data

Request for different video:

[CachedYouTubeProxy] Request for video: 9bZkp7q19f0
[CachedYouTubeProxy] Cache MISS - downloading from YouTube
[YouTubeDownloader] Downloading video 9bZkp7q19f0 from YouTube (slow operation)...
[CachedYouTubeProxy] Cached video: 9bZkp7q19f0

Request for first video again (cached):

[CachedYouTubeProxy] Request for video: dQw4w9WgXcQ
[CachedYouTubeProxy] Cache HIT - returning cached data

--- Example 4: Smart Proxy (Authentication) ---

Initial balance: $1000.00

[BankAccountProxy] Authentication attempt...
[BankAccountProxy] Invalid PIN. Attempt 1/3

[BankAccountProxy] Authentication attempt...
[BankAccountProxy] Invalid PIN. Attempt 2/3

[BankAccountProxy] Authentication attempt...
[BankAccountProxy] Invalid PIN. Attempt 3/3
[BankAccountProxy] Account LOCKED after 3 failed attempts

[BankAccountProxy] Authentication attempt...
[BankAccountProxy] Account is LOCKED due to multiple failed attempts

--- Example 5: Remote Proxy (Load Balancer) ---

Making multiple requests:

[LoadBalancerProxy] Routing request to Server-1 (Round-robin)
[RealServer Server-1] Processing request: /api/users

[LoadBalancerProxy] Routing request to Server-2 (Round-robin)
[RealServer Server-2] Processing request: /api/products

[LoadBalancerProxy] Routing request to Server-3 (Round-robin)
[RealServer Server-3] Processing request: /api/orders

[LoadBalancerProxy] Routing request to Server-1 (Round-robin)
[RealServer Server-1] Processing request: /api/payments

--- Summary ---
Virtual Proxy: Delays expensive object creation until needed
Protection Proxy: Controls access based on permissions
Caching Proxy: Caches results to avoid redundant operations
Smart Proxy: Adds additional logic like authentication
Remote Proxy: Represents object in different location (load balancing)

All proxies maintain the same interface as the real object
```

## Key Takeaways

- **Same Interface**: Proxy implements same interface as real object
- **Control Layer**: Adds control/logic without modifying real object
- **Multiple Types**: Virtual, Protection, Caching, Smart, Remote, Logging
- **Transparent**: Client doesn't know it's using proxy
- **Lazy Loading**: Virtual proxy defers expensive initialization
- **Security**: Protection proxy enforces access control
- **Performance**: Caching proxy improves performance
- **Cross-Cutting Concerns**: Handles concerns orthogonal to business logic

## Design Considerations

### When to Create Real Object

**Virtual Proxy**: Create on first access
```go
if p.realObject == nil {
    p.realObject = NewRealObject()
}
```

**Other Proxies**: Usually create in proxy constructor
```go
func NewProxy() *Proxy {
    return &Proxy{
        realObject: NewRealObject(),
    }
}
```

### Thread Safety

Proxies in concurrent environments need thread safety:

```go
type ThreadSafeProxy struct {
    mu         sync.Mutex
    realObject *RealObject
}

func (p *ThreadSafeProxy) Operation() {
    p.mu.Lock()
    defer p.mu.Unlock()
    
    if p.realObject == nil {
        p.realObject = NewRealObject()
    }
    p.realObject.Operation()
}
```

### Proxy Chains

Proxies can be chained:
```go
realServer := NewRealServer()
cachingProxy := NewCachingProxy(realServer)
loggingProxy := NewLoggingProxy(cachingProxy)
authProxy := NewAuthProxy(loggingProxy)

// Client uses authProxy
```

## Common Mistakes

1. **Not Implementing Full Interface**: Proxy must implement all methods
2. **Breaking Transparency**: Client shouldn't know about proxy
3. **Forgetting Delegation**: Proxy must eventually delegate to real object
4. **Mutable Proxies**: Proxies should typically be immutable
5. **Heavy Proxies**: Adding too much logic - keep proxies focused
6. **Not Handling Nil**: Virtual proxy must handle nil real object

## Best Practices

1. **Same Interface**: Always implement the same interface as real object
2. **Single Responsibility**: Each proxy type should handle one concern
3. **Delegate**: Proxy should delegate, not reimplement logic
4. **Transparent**: Client shouldn't need to know about proxy
5. **Thread-Safe**: Consider concurrency if needed
6. **Document Behavior**: Clearly document what proxy adds
7. **Test Both**: Test both proxy and real object

## Real-World Examples

- **Java RMI**: Remote proxies for distributed objects
- **Spring AOP**: Proxies for aspect-oriented programming
- **ORM Frameworks**: Lazy loading proxies (Hibernate, GORM)
- **HTTP Proxies**: nginx, squid for caching and load balancing
- **Smart Pointers**: C++ shared_ptr, unique_ptr
- **Python @property**: Property accessors are proxies
- **Database Connection Pools**: Proxy for actual database connections

## Combining Proxy with Other Patterns

### Proxy + Singleton
```go
var instance *RealObject
var once sync.Once

type SingletonProxy struct {}

func (p *SingletonProxy) Operation() {
    once.Do(func() {
        instance = NewRealObject()
    })
    instance.Operation()
}
```

### Proxy + Factory
```go
type ProxyFactory struct {}

func (f *ProxyFactory) CreateProxy(proxyType string) Subject {
    switch proxyType {
    case "caching":
        return NewCachingProxy()
    case "logging":
        return NewLoggingProxy()
    default:
        return NewRealObject()
    }
}
```

### Proxy + Strategy
```go
type CachingStrategy interface {
    Get(key string) (interface{}, bool)
    Set(key string, value interface{})
}

type CachingProxy struct {
    realObject *RealObject
    strategy   CachingStrategy
}
```

## When NOT to Use Proxy

- When direct access is simpler and sufficient
- When the overhead isn't justified
- When you need to change the interface (use Adapter)
- When you want to add functionality, not control access (use Decorator)
- When object creation is cheap and fast
- When access control isn't needed

## Testing Proxies

Test proxy behavior separately from real object:

```go
func TestCachingProxy(t *testing.T) {
    proxy := NewCachingProxy()
    
    // First call - should hit real object
    result1 := proxy.GetData("key1")
    
    // Second call - should hit cache
    result2 := proxy.GetData("key1")
    
    if result1 != result2 {
        t.Error("Cache should return same result")
    }
    
    // Verify cache hit
    if !proxy.IsCached("key1") {
        t.Error("Expected key1 to be cached")
    }
}
```
