//go:build !dev
// +build !dev

package main

import (
	"embed"
	"fmt"
	"net/http"
)

//go:embed public
var publicFS embed.FS

func public() http.Handler {
	fmt.Println("Building static files for development...")
	return http.FileServer(http.FS(publicFS))
}
