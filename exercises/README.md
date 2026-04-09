# Exercise 1 — FetchWithTimeout

## Objective

Build a `Fetch` function that calls external APIs with context timeout and cancellation support.

---

## Package

```
package exercise1
```

---

## Requirements

### Sentinel Errors

Define two package-level sentinel errors:

- `ErrTimeout` — returned when the context deadline expires before the server responds
- `ErrCancelled` — returned when the context is already cancelled before the request is made

### Function Signature

```go
func Fetch(ctx context.Context, url string) ([]byte, error)
```

### Behaviour

1. If the context is already cancelled before making the request, return `nil, ErrCancelled`
2. Build the HTTP GET request with the context attached using `http.NewRequestWithContext`
3. If the request fails because the context deadline exceeded, return `nil, ErrTimeout`
4. On success, read and return the full response body

---

## Deliverables

### Subtest 1 — Successful GET request

- Spin up a test server using `httptest.NewServer`
- Handler writes `"hello"` to the response
- Call `Fetch` with `context.Background()`
- Assert no error and body equals `"hello"`

### Subtest 2 — Context already cancelled

- Create a context with `context.WithCancel`
- Cancel it immediately before calling `Fetch`
- Assert `got` is `nil`
- Assert `errors.Is(err, ErrCancelled)` is true

### Subtest 3 — Context times out before server responds

- Spin up a slow test server whose handler blocks using `select` on `ctx.Done()` and `time.After(5 * time.Second)`
- Use `r.Context()` inside the handler so it reacts to client cancellation
- Create a context with `context.WithTimeout` set to `100ms`
- Assert `got` is `nil`
- Assert `errors.Is(err, ErrTimeout)` is true
- Full test suite must complete in approximately 100ms, not 5 seconds

---

## Stretch Goal

Implement `FetchWithRetry` with exponential backoff:

```go
func FetchWithRetry(ctx context.Context, url string, maxRetries int) ([]byte, error)
```

- Retry only on `ErrTimeout`, not on `ErrCancelled`
- Double the wait time between each retry
- If the context expires during a retry wait, return immediately