[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/orzkratos/ebzkratos/release.yml?branch=main&label=BUILD)](https://github.com/orzkratos/ebzkratos/actions?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/orzkratos/ebzkratos)](https://pkg.go.dev/github.com/orzkratos/ebzkratos)
[![Coverage Status](https://img.shields.io/coveralls/github/orzkratos/ebzkratos/main.svg)](https://coveralls.io/github/orzkratos/ebzkratos?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/orzkratos/ebzkratos.svg)](https://github.com/orzkratos/ebzkratos/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/orzkratos/ebzkratos)](https://goreportcard.com/report/github.com/orzkratos/ebzkratos)

# ebzkratos

Go error wrapper for Kratos framework that solves nil interface problems.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Main Features

🎯 **Safe Error Wrapping**: Avoids Go's (*T)(nil) != nil problem by not implementing error interface  
🔄 **Type Conversion**: Safe conversion between generic errors and Ebz instances with correct nil handling  
📋 **Comprehensive API**: NewEbz, New, As, Is, FromError, From functions providing complete error handling

## Installation

```bash
go get github.com/orzkratos/ebzkratos
```

## Usage

### Basic Error Wrapping

```go
package main

import (
    "github.com/go-kratos/kratos/v2/errors"
    "github.com/orzkratos/ebzkratos"
)

func main() {
    // Create Kratos error
    erk := errors.BadRequest("INVALID_PARAM", "invalid parameter")
    
    // Wrap with Ebz - safe from nil interface issues
    ebz := ebzkratos.New(erk)
    
    // Safe nil checking - no (*T)(nil) != nil problems
    if ebz != nil {
        // Process error content in safe mode
        println("Error:", ebz.Erk.Message)
    }
}
```

### Type Conversion

```go
func processError(err error) {
    // Convert generic error to Ebz
    if ebz, ok := ebzkratos.As(err); ok {
        // Handle as Kratos error through Ebz wrapper
        println("Kratos error:", ebz.Erk.Reason)
    }
    
    // Alternative conversion
    ebz := ebzkratos.From(err)
    if ebz != nil {
        println("Converted:", ebz.Erk.Message)
    }
}
```

### Error Comparison

```go
func compareErrors(ebz1, ebz2 *ebzkratos.Ebz) {
    // Check errors equivalence with nil handling
    if ebzkratos.Is(ebz1, ebz2) {
        println("Errors are same type")
    }
}
```

### Production Assertions

```go
import "github.com/orzkratos/ebzkratos/must/ebzmust"

func criticalOperation() {
    result, ebz := someOperation()
    
    // Assert no error - panic if ebz is not nil
    ebzmust.Done(ebz)
    
    // Continue with result in safe mode
    processResult(result)
}
```

## Core Constraint

**STRUCTURAL INVARIANT**: `ebz != nil ⇒ ebz.Erk must be non-nil`

This constraint is enforced via `must.Full` validation in constructors, ensuring:

- No ambiguous intermediate states
- Simplified error handling logic  
- Fail-fast error detection
- Elimination of nil interface complications

## API Reference

### Creation Functions

- `NewEbz(erk *errors.Error) *Ebz` - Creates error wrapper instance with validation
- `New(erk *errors.Error) *Ebz` - Concise creation alias

### Conversion Functions  

- `As(err error) (*Ebz, bool)` - Type conversion from generic error to Ebz
- `FromError(err error) *Ebz` - Transforms generic error into safe Ebz instance
- `From(err error) *Ebz` - Concise alias of FromError

### Comparison Functions

- `Is(ebz1, ebz2 *Ebz) bool` - Checks if two Ebz values are equivalent

### Assert Functions

- `ebzmust.Done(ebz *Ebz)` - Checks no-error convention with panic when violated
- `ebzmust.Must(ebz *Ebz)` - Demands perfect execution with fail-fast termination

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a mistake?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/orzkratos/ebzkratos.svg?variant=adaptive)](https://starchart.cc/orzkratos/ebzkratos)