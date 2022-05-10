package internal

import "unsafe"

// Getg return the current go routine address
func Getg() unsafe.Pointer
