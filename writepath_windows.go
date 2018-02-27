// Copyright 2018 Richard Mueller.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
)

var namepath string

func init() {
	flag.StringVar(&namepath, "name", "PATH", "name of path to use instead of PATH (for output only!)")
}

func writePath(path string) {
	fmt.Printf("%s=%s\n", namepath, path)
}
