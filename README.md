# go-pathlib

[![Go Reference](https://pkg.go.dev/badge/github.com/njchilds90/go-pathlib.svg)](https://pkg.go.dev/github.com/njchilds90/go-pathlib)  
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A simple, idiomatic Go library for filesystem path manipulation inspired by Python’s `pathlib` APIs.  
It provides a structured, object-oriented way to build, inspect, and transform file paths without resorting to raw `string` handling.

---

## ✨ Features

- 🧱 `Path` struct wrapping filesystem locations
- 🪛 Convenience methods like `Join`, `Parent`, `Exists`, `IsFile`, etc.
- 🔄 Immutable path representation (methods return new `Path` objects)
- 📦 Designed for both quick scripts and larger codebases
- 📑 Integrates naturally with Go’s `os` and `path/filepath` packages

---

## 📦 Installation

```bash
go get github.com/njchilds90/go-pathlib
```

---

## 🧠 Why Use go-pathlib?

Go’s standard library uses raw strings for paths, which is powerful but can be error-prone when composing and analyzing complex paths.

A dedicated `Path` type lets you:

- avoid mixing separators manually
- express intention more clearly
- write code closer to Python’s expressive `pathlib.Path` style

---

## 🚀 Quick Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/njchilds90/go-pathlib"
)

func main() {
	base := pathlib.NewPath("/usr/local/bin")

	// join segments
	full := base.Join("myapp", "config.json")
	fmt.Println("Full path:", full.String())

	// inspect path
	exists, err := full.Exists()
	if err != nil {
		log.Fatalf("exists check failed: %v", err)
	}

	fmt.Printf("Does config exist? %v\n", exists)

	// get parent directory
 parent := full.Parent()
 fmt.Println("Parent:", parent.String())
}
```

---

## 📘 API Overview

### Creating Paths

```go
p := pathlib.NewPath("path/to/file")
```

- Always returns a new `Path` instance  
- Behaves consistently across POSIX and Windows separators

### Common Methods

```go
p.String()            // string form of the path
p.Join("a", "b")      // returns a new Path with appended segments
p.Parent() *Path      // parent directory reference
p.IsFile() (bool, error)   // check if path is a file
p.IsDir() (bool, error)    // check if path is a directory
p.Exists() (bool, error)   // check existence
p.Base() string        // last element (filename or dir)
p.Extension() string   // file extension
```

### Usage Patterns

```go
// Join multiple segments
subpath := pathlib.NewPath("root").Join("subdir", "file.txt")

// Check if file exists
exists, err := subpath.Exists()
if err != nil {
    panic(err)
}

if exists {
    fmt.Println("Found:", subpath.String())
}
```

---

## 🧪 Example Tests

```go
func TestPathJoin(t *testing.T) {
    p := pathlib.NewPath("root")
    got := p.Join("child", "leaf.txt").String()
    want := filepath.Join("root", "child", "leaf.txt")
    if got != want {
        t.Errorf("expected %q, got %q", want, got)
    }
}
```

---

## 📌 Design Philosophy

- **Explicit** over implicit — no hidden behavior  
- **Immutable** path values — safer chaining of operations  
- **Minimal dependencies** — zero external requirements  
- **Composable** with standard library (`os`, `filepath`, `io`, etc.)

---

## 🤝 When To Use (And When Not To)

### Good For

- Scripts that build or manipulate filesystem paths
- Projects where readability of path logic matters
- Codebases that benefit from path abstraction

### Traditional Approach Still Fine For

- Simple one-off string path handling
- Low-overhead utilities that just need raw paths

---

## 🛠 Future Enhancements

- Glob and recursive matching utilities
- Built-in `Walk` helpers
- Integration with embed patterns
- More cross-platform helpers
- Better support for symlinks and canonicalization

---

## 📄 License

MIT License — see LICENSE file.

---

## 🙌 Contributions

Contributions, issues, and pull requests are welcome.  
Please keep the API clear, intuitive, and idiomatic.
