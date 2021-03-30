//go:build !windows
// +build !windows

package main

import _ "embed"

//go:embed image.png
var iconData []byte
