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

func cleanPath(path string) string {
	var fis []os.FileInfo
	haveFile := func(undertest os.FileInfo) bool {
		for _, fi := range fis {
			if os.SameFile(fi, undertest) {
				return true
			}
		}
		return false
	}
	parts := splitPath(path)
	var newpath []string
	for _, part := range parts {
		canon, err := canonicalFilepath(part)
		if err != nil {
			warn("invalid file path \"%s\": %v", part, err)
			continue
		}
		fi, err := os.Stat(canon)
		if err != nil {
			warn("not found \"%s\": %v", canon, err)
			continue
		}
		if haveFile(fi) {
			warn("\"%s\" multiple defined", canon)
			continue
		}
		fis = append(fis, fi)
		if !fi.Mode().IsDir() {
			warn("\"%s\" is not a directory", canon)
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
