# go-pathlib

[![Go](https://github.com/njchilds90/go-pathlib/actions/workflows/ci.yml/badge.svg)](https://github.com/njchilds90/go-pathlib/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/njchilds90/go-pathlib.svg)](https://pkg.go.dev/github.com/njchilds90/go-pathlib)

Lightweight, zero-dependency, fluent Path type for Go inspired by Python's pathlib.

## Installation
```bash
go get github.com/njchilds90/go-pathlib
Usage
Gopackage main

import (
	"fmt"
	"github.com/njchilds90/go-pathlib"
)

func main() {
	p := pathlib.New("/home/user")
	fmt.Println(p.Join("docs", "report.pdf"))           // /home/user/docs/report.pdf
	fmt.Println(p.WithExt(".md"))                      // /home/user.md
	fmt.Println(p.Dir().Join("backup"))                // /home/backup

	abs, _ := p.Abs()
	fmt.Println(abs.IsAbs())                           // true
}
