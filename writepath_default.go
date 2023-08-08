// Copyright 2018 Richard Mueller.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// for all OS but Windows

//go:build !windows
// +build !windows

package main

import "fmt"

func writePath(path string) {
	fmt.Println(path)
}
