//go:build tools
// +build tools

package main

import (
	_ "github.com/cespare/reflex"
	_ "github.com/g4s8/envdoc"
	_ "github.com/go-delve/delve/cmd/dlv"
	_ "github.com/jmattheis/goverter/cmd/goverter"
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
	_ "mvdan.cc/gofumpt"
)
