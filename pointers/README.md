# Pointers & Errors — Quick Notes

---

## Pointers

**The Problem**
When you pass a struct to a method Go makes a **copy**. Changes on the copy are lost when the method ends.

```
w Wallet   → copy → changes lost
w *Wallet  → real → changes persist
```

**The Fix — Pointer Receiver**
```go
func (w *Wallet) Deposit(amount Bitcoin) {
    w.balance += amount  // modifies real wallet
}
```

**Rule of thumb** — if your method needs to modify the struct use `*Type`. If it only reads use `Type`.

---

## Custom Types

```go
type Bitcoin int
```

Creates a distinct type backed by `int`. The compiler prevents accidental mixing of types. You must explicitly convert:
```go
wallet.Deposit(Bitcoin(10))  // correct
wallet.Deposit(10)           // compiler error
```

---

## Stringer Interface

```go
func (b Bitcoin) String() string {
    return fmt.Sprintf("%d BTC", b)
}
```

Any type with `String() string` automatically satisfies `fmt.Stringer`. The `fmt` package calls it automatically when using `%s` or `%v`. No registration needed — implicit satisfaction.

---

## Errors

**Creating sentinel errors:**
```go
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
```

**Returning errors:**
```go
func (w *Wallet) Withdraw(amount Bitcoin) error {
    if amount > w.balance {
        return ErrInsufficientFunds
    }
    w.balance -= amount
    return nil
}
```

**Go error handling pattern:**
```go
err := wallet.Withdraw(Bitcoin(100))
if err != nil {
    // handle error
}
```

No exceptions in Go — errors are just values returned explicitly.

---

## t.Fatal vs t.Errorf

| | Behaviour |
|---|---|
| `t.Errorf` | logs failure, test continues |
| `t.Fatal` | logs failure, test stops immediately |

Use `t.Fatal` when subsequent checks are meaningless if the first fails.

---

## Your Code Highlights

**Sentinel error pattern** — `ErrInsufficientFunds` as package level variable allows callers to compare directly against it.

**Side effect verification** — checking balance unchanged after failed withdrawal shows mature testing thinking.

**Inline helper functions** — `checkBalance`, `checkError`, `checkNoError` scoped inside the test function. Clean and purposeful.

---

## MPESA Connection

| Concept | Project Application |
|---|---|
| Pointer receivers | `Reconciler` accumulating real totals |
| Custom types | `type KES float64` for Kenyan Shillings |
| Stringer | `KES` printing as `"KES 1,250.00"` |
| Sentinel errors | `ErrAmountMismatch`, `ErrTransactionNotFound` |
| `t.Fatal` | Stopping tests when critical data missing |

---

## Key Takeaways

- Pointer receivers modify real data, value receivers work on copies
- Custom types enforce correctness at compile time
- Errors are values in Go — return and check them explicitly
- `nil` means no error, non-nil means something went wrong
- Sentinel errors allow precise error identification and comparison