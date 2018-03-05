// Copyright 2018 Richard Mueller.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type void struct{}

var pathSeparator = string([]byte{os.PathListSeparator})

// command line options
var (
	path                        string
	verbose, help, keepsymlinks bool
)

func init() {
	flag.StringVar(&path, "path", "", "path to use instead of PATH from environment")
	flag.BoolVar(&verbose, "verbose", false, "verbose output: print warnings to stderr")
	flag.BoolVar(&verbose, "v", false, "alias for verbose")
	flag.BoolVar(&keepsymlinks, "keepsymlinks", false,
		"do not evaluate (dereference) symbolic links")
	flag.BoolVar(&help, "help", false, "print usage and exit")
}

func log(f string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, f+"\n", a...)
}

func warn(f string, a ...interface{}) {
	if verbose {
		log("WARNING: "+f, a...)
	}
}

func splitPath(path string) []string {
	return filepath.SplitList(path)
}

func joinPath(parts []string) string {
	return strings.Join(parts, pathSeparator)
}

func cleanFilepath(p string) (string, error) {
	if keepsymlinks {
		return filepath.Clean(p), nil
	}
	return filepath.EvalSymlinks(p)
}

func canonicalFilepath(p string) (string, error) {
	p0, err := cleanFilepath(p)
	if err != nil {
		return "", err
	}
	return filepath.Abs(p0)
}

func verifyDir(d string) error {
	fi, err := os.Stat(d)
	if err != nil {
		return err
	}
	if !fi.Mode().IsDir() {
		return fmt.Errorf("%s is not a directory", d)
	}
	return nil
}

func cleanPath(path string) string {
	parts := splitPath(path)
	var newpath []string
	seen := make(map[string]void)
	for _, part := range parts {
		canon, err := canonicalFilepath(part)
		if err != nil {
			warn("invalid file path %s: %v", part, err)
			continue
		}
		if _, ok := seen[canon]; ok {
			warn("%s multiple defined", canon)
			continue
		}
		seen[canon] = void{}
		if err := verifyDir(canon); err != nil {
			warn("%v", err)
			continue
		}
		newpath = append(newpath, canon)
	}
	return joinPath(newpath)
}

func reverse(in []string) (out []string) {
	for i := len(in) - 1; i >= 0; i = i - 1 {
		out = append(out, in[i])
	}
	return
}

func prependPath(path string, args []string) {
	if len(args) > 0 {
		parts := reverse(splitPath(path))
		parts = append(parts, reverse(args)...)
		parts = reverse(parts)
		path = joinPath(parts)
	}
	newpath := cleanPath(path)
	writePath(newpath)
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	args := flag.Args()
	if len(path) == 0 {
		path = os.Getenv("PATH")
	}
	prependPath(path, args)
}
