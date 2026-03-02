// Package pathlib provides an immutable, fluent Path type for filesystem path manipulation.
// Inspired by Python's pathlib, built on Go's path/filepath with zero external dependencies.
package pathlib

import (
	"path/filepath"
	"strings"
)

// Path represents a filesystem path. It is immutable and uses platform-native separators.
type Path string

// New returns a cleaned Path. Empty string becomes ".".
func New(path string) Path {
	if path == "" {
		path = "."
	}
	return Path(filepath.Clean(path))
}

// String returns the string representation of the path.
func (p Path) String() string { return string(p) }

// Join appends elements and returns a new cleaned Path.
func (p Path) Join(elem ...string) Path {
	elems := make([]string, len(elem)+1)
	elems[0] = string(p)
	copy(elems[1:], elem)
	return New(filepath.Join(elems...))
}

// Clean returns the cleaned version of the path.
func (p Path) Clean() Path { return New(string(p)) }

// Dir returns the directory component.
func (p Path) Dir() Path { return New(filepath.Dir(string(p))) }

// Base returns the last element of the path.
func (p Path) Base() string { return filepath.Base(string(p)) }

// Ext returns the file extension (including dot).
func (p Path) Ext() string { return filepath.Ext(string(p)) }

// WithExt returns a new Path with the extension replaced.
func (p Path) WithExt(ext string) Path {
	s := string(p)
	if ext != "" && ext[0] != '.' {
		ext = "." + ext
	}
	if s == "." || s == ".." {
		return p
	}
	base := strings.TrimSuffix(s, p.Ext())
	return New(base + ext)
}

// WithBase returns a new Path with the base name replaced.
func (p Path) WithBase(base string) Path {
	return p.Dir().Join(base)
}

// IsAbs reports whether the path is absolute.
func (p Path) IsAbs() bool { return filepath.IsAbs(string(p)) }

// Abs returns an absolute representation of the path.
// Uses filepath.Abs (may read current working directory if relative).
func (p Path) Abs() (Path, error) {
	abs, err := filepath.Abs(string(p))
	if err != nil {
		return "", err
	}
	return New(abs), nil
}

// Rel returns a relative path from base to p.
func (p Path) Rel(base Path) (Path, error) {
	rel, err := filepath.Rel(string(base), string(p))
	if err != nil {
		return "", err
	}
	return New(rel), nil
}

// VolumeName returns the leading volume name (Windows only).
func (p Path) VolumeName() string { return filepath.VolumeName(string(p)) }

// Matches reports whether the path matches the shell pattern.
func (p Path) Matches(pattern string) (bool, error) {
	return filepath.Match(pattern, string(p))
}
