//+build dev
//go:build dev
// +build dev

package main

import (
	"fmt"
	"net/http"
)

func public() http.Handler {
	fmt.Println("Building static files for development...")
	return http.StripPrefix("/public/", http.FileServer(http.Dir("public")))
}
