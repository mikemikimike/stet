// Package jsts defines the JavaScript and TypeScript extensions supported by the CLI.
package jsts

import "strings"

type extension struct {
	suffix        string
	fenceLanguage string
}

var supportedExtensions = [...]extension{
	{suffix: ".js", fenceLanguage: "javascript"},
	{suffix: ".jsx", fenceLanguage: "javascript"},
	{suffix: ".mjs", fenceLanguage: "javascript"},
	{suffix: ".cjs", fenceLanguage: "javascript"},
	{suffix: ".ts", fenceLanguage: "typescript"},
	{suffix: ".tsx", fenceLanguage: "typescript"},
}

// SupportedExtensions returns the supported JavaScript and TypeScript suffixes.
func SupportedExtensions() []string {
	extensions := make([]string, len(supportedExtensions))
	for i, ext := range supportedExtensions {
		extensions[i] = ext.suffix
	}
	return extensions
}

// IsSupportedExtension reports whether ext is a supported JavaScript or TypeScript suffix.
func IsSupportedExtension(ext string) bool {
	for _, supported := range supportedExtensions {
		if strings.EqualFold(ext, supported.suffix) {
			return true
		}
	}
	return false
}

// FenceLanguage returns the Markdown code-fence language for a supported suffix.
func FenceLanguage(ext string) string {
	for _, supported := range supportedExtensions {
		if strings.EqualFold(ext, supported.suffix) {
			return supported.fenceLanguage
		}
	}
	return ""
}
