# Centralize Error Wrapper

Simplify error handling in your Go projects with the Centralize Error Wrapper. Achieve uniformity in error management, customize error messages, seamlessly integrate with logging, and simplify reporting to external systems. Enhance the robustness of your codebase effortlessly.

**Centralize Error Wrapper for Go-Projects**

Simplify error handling in your Go projects with the Centralize Error Wrapper. Achieve uniformity in error management, customize error messages, seamlessly integrate with logging, and simplify reporting to external systems. Enhance the robustness of your codebase effortlessly.

**Usage:**

1. Import the package:

```go
import "github.com/pixel8labs/centralize-error-wrapper"
```

2. Wrap errors, unwrap, and make a new err effortlessly:

```go
// Wrap an error with additional context
wrappedErr := centralizeerr.Wrap(originalErr, "Failed to perform operation")

// Unwrap the error
centralizeerr.Unwrap(wrappedErr)

// Make a new error
centralizeerr.New(Key)
```

Elevate your Go projects by standardizing and customizing error handling

### by Abdul Salam (abdulsalam01)
