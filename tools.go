//go:build tools
// +build tools

package tools

import (
	// Generate embedded files.
	_ "github.com/go-bindata/go-bindata/v3/go-bindata"
	_ "k8s.io/code-generator"
)
