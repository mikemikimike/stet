package jsts

import (
	"reflect"
	"testing"
)

func TestSupportedExtensions(t *testing.T) {
	want := []string{".js", ".jsx", ".mjs", ".cjs", ".ts", ".tsx"}
	got := SupportedExtensions()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("SupportedExtensions() = %v, want %v", got, want)
	}

	got[0] = ".changed"
	if next := SupportedExtensions(); next[0] != ".js" {
		t.Fatalf("mutating returned extensions changed shared set: %v", next)
	}
}

func TestIsSupportedExtension(t *testing.T) {
	tests := []struct {
		ext  string
		want bool
	}{
		{ext: ".js", want: true},
		{ext: ".jsx", want: true},
		{ext: ".mjs", want: true},
		{ext: ".cjs", want: true},
		{ext: ".ts", want: true},
		{ext: ".tsx", want: true},
		{ext: ".JSX", want: true},
		{ext: ".mts", want: false},
		{ext: ".cts", want: false},
		{ext: ".go", want: false},
	}
	for _, tt := range tests {
		if got := IsSupportedExtension(tt.ext); got != tt.want {
			t.Errorf("IsSupportedExtension(%q) = %t, want %t", tt.ext, got, tt.want)
		}
	}
}

func TestFenceLanguage(t *testing.T) {
	tests := []struct {
		ext  string
		want string
	}{
		{ext: ".js", want: "javascript"},
		{ext: ".jsx", want: "javascript"},
		{ext: ".mjs", want: "javascript"},
		{ext: ".cjs", want: "javascript"},
		{ext: ".ts", want: "typescript"},
		{ext: ".tsx", want: "typescript"},
		{ext: ".go", want: ""},
	}
	for _, tt := range tests {
		if got := FenceLanguage(tt.ext); got != tt.want {
			t.Errorf("FenceLanguage(%q) = %q, want %q", tt.ext, got, tt.want)
		}
	}
}
