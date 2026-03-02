package pathlib

import (
	"testing"
)

func TestPath(t *testing.T) {
	tests := []struct {
		name     string
		p        Path
		wantStr  string
		wantBase string
		wantExt  string
	}{
		{"simple", New("/home/user/file.txt"), "/home/user/file.txt", "file.txt", ".txt"},
		{"relative", New("docs/../file.go"), "file.go", "file.go", ".go"},
		{"empty", New(""), ".", ".", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.wantStr {
				t.Errorf("String() = %q, want %q", got, tt.wantStr)
			}
			if got := tt.p.Base(); got != tt.wantBase {
				t.Errorf("Base() = %q, want %q", got, tt.wantBase)
			}
			if got := tt.p.Ext(); got != tt.wantExt {
				t.Errorf("Ext() = %q, want %q", got, tt.wantExt)
			}
		})
	}
}

func TestJoin(t *testing.T) {
	p := New("/home")
	if got := p.Join("user", "docs", "file.txt").String(); got != "/home/user/docs/file.txt" {
		t.Errorf("Join() = %q", got)
	}
}

func TestWithExt(t *testing.T) {
	p := New("/home/file.txt")
	if got := p.WithExt(".md").String(); got != "/home/file.md" {
		t.Errorf("WithExt = %q", got)
	}
	if got := p.WithExt("json").String(); got != "/home/file.json" {
		t.Errorf("WithExt(no dot) = %q", got)
	}
}

func TestAbsRel(t *testing.T) {
	p := New("file.txt")
	abs, err := p.Abs()
	if err != nil {
		t.Error(err)
	}
	if !abs.IsAbs() {
		t.Error("Abs() should be absolute")
	}

	rel, err := abs.Rel(New("/"))
	if err != nil {
		t.Error(err)
	}
	if rel.String() == "" {
		t.Error("Rel() should return non-empty")
	}
}
