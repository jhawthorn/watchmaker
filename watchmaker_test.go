package main

import "testing"

var ignoredPaths = []string{
	// vim to test if the directory is writable
	"4913",
	"./4913",
	"src/4913",
	"./src/4913",

	// temporary and swap files
	"foo~",
	"./foo~",
	"foo.swp",
	"./foo.swp",
	"foo.swx",
	"./foo.swx",
	"foo.swpx",
	"./foo.swpx",

	// .lock
	"foo.lock",
	"./foo.lock",

	// hidden files and directories
	".git",
	"./.git",
	".git/HEAD",
	"./.git/HEAD",
	".hidden",
	"./.hidden",
}

var validPaths = []string{
	"Makefile",
	"./Makefile",
	"makefile",
	"4913.txt",
	"./4913.txt",
	"foobar",
	"./foobar",
	"foobar.txt",
	"./foobar.txt",
	"./foo/bar/baz",
}

func TestIgnored(t *testing.T) {
	for _, test := range ignoredPaths {
		if !ignored(test) {
			t.Errorf("Expected path '%s' to be ignored", test)
		}
	}

	for _, test := range validPaths {
		if ignored(test) {
			t.Errorf("Expected '%s' to be a watched path", test)
		}
	}
}
