// Copyright (c) 2012-2022 The go-diff authors. All rights reserved.
// https://github.com/sergi/go-diff
// https://github.com/twpayne/go-diff
// See the included LICENSE file for license details.
//
// go-diff is a Go implementation of Google's Diff, Match, and Patch library
// Original library is Copyright (c) 2006 Google Inc.
// http://code.google.com/p/google-diff-match-patch/

package diffmatchpatch

import "os"

func speedtestTexts() (s1 string, s2 string) {
	d1, err := os.ReadFile("testdata/speedtest1.txt")
	if err != nil {
		panic(err)
	}
	d2, err := os.ReadFile("testdata/speedtest2.txt")
	if err != nil {
		panic(err)
	}

	return string(d1), string(d2)
}
